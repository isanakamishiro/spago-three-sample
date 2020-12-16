package views

import (
	"app/frontend/lib/threejs"
	"log"
	"math"
	"syscall/js"

	"github.com/nobonobo/spago"
)

//go:generate spago generate -c MikuView -p views mikuview.html

// loadScript synchronous javascript loader
func loadScript(url string) {

	document := js.Global().Get("document")

	ch := make(chan bool)
	script := document.Call("createElement", "script")
	script.Set("src", url)
	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		defer fn.Release()
		close(ch)
		return nil
	})
	script.Call("addEventListener", "load", fn)
	document.Get("head").Call("appendChild", script)
	<-ch
}

// MikuView  ...
type MikuView struct {
	spago.Core

	three    *threejs.ThreeJs
	camera   threejs.Camera
	scene    threejs.Scene
	mesh     threejs.Mesh
	renderer threejs.Renderer
	clock    threejs.Clock

	renderID     js.Value
	OnResizeFunc js.Func

	outlineEffectMod js.Value
	// outlineEffect    *effects.OutlineEffect
	outlineEffect js.Value

	orbitControlsMod js.Value
	controls         js.Value

	mmdLoaderMod js.Value
	loader       js.Value

	mmdAnimationHelperMod js.Value
	helper                js.Value

	statsMod js.Value
	stats    js.Value
}

// NewMikuView is ...
func NewMikuView(three *threejs.ThreeJs) *MikuView {

	loadScript("assets/js/libs/ammo.wasm.js")
	outlineEffectMod := spago.LoadModule([]string{"OutlineEffect"}, "./assets/jsm/effects/OutlineEffect2.js")
	orbitControlsMod := spago.LoadModule([]string{"OrbitControls"}, "./assets/jsm/controls/OrbitControls.js")
	mmdLoaderMod := spago.LoadModule([]string{"MMDLoader"}, "./assets/jsm/loaders/MMDLoader.js")
	mmdAnimationHelperMod := spago.LoadModule([]string{"MMDAnimationHelper"}, "./assets/jsm/animation/MMDAnimationHelper.js")
	statsMod := spago.LoadModule([]string{"Stats"}, "./assets/jsm/libs/stats.module.js")

	view := &MikuView{
		three:                 three,
		outlineEffectMod:      outlineEffectMod[0],
		orbitControlsMod:      orbitControlsMod[0],
		mmdLoaderMod:          mmdLoaderMod[0],
		mmdAnimationHelperMod: mmdAnimationHelperMod[0],
		statsMod:              statsMod[0],
		// camera:   camera,
		// scene:    scene,
		// mesh:     mesh,
		// renderer: renderer,
	}

	return view
}

// Mount is ...
func (c *MikuView) Mount() {

	fn := js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		js.Global().Set("Ammo", args[0])

		c.initScene()

		js.Global().Get("document").Get("body").Call("appendChild", c.renderer.DomElement())
		js.Global().Get("document").Get("body").Call("appendChild", c.stats.Get("dom"))

		c.OnResizeFunc = js.FuncOf(c.OnResize)
		js.Global().Call("addEventListener", "resize", c.OnResizeFunc)

		c.OnResize(js.Null(), nil)
		c.animate(js.Null(), nil)

		return nil
	})
	js.Global().Call("Ammo").Call("then", fn)
	// fn.Release()

	// println("Mount!!")

}

// Unmount ...
func (c *MikuView) Unmount() {

	js.Global().Call("cancelAnimationFrame", c.renderID)
	js.Global().Call("removeEventListener", "resize", c.OnResizeFunc)
	// c.OnResizeFunc.Release()

	// println("Unmount!!")

}

func (c *MikuView) initScene() {

	// Clock
	c.clock = c.three.NewClock(true)

	// Camera
	c.camera = c.three.NewPerspectiveCamera(45, 4/3, 1, 2000)
	c.camera.Position().SetZ(40)
	c.camera.Position().SetY(40)

	// Scene
	c.scene = c.three.NewScene()
	c.scene.SetBackground(c.three.NewColor2(1, 1, 1).JSValue())

	// Grid
	// gridHelper := c.three.NewPolarGridHelper(30, 10)
	// gridHelper.Position().SetY(-10)
	// c.scene.Add(gridHelper.JSValue())

	// Ground
	ground := c.three.NewMesh(
		c.three.NewPlaneBufferGeometry(100, 100, 1, 1),
		c.three.NewMeshPhongMaterial(map[string]interface{}{
			"color": 0xdddddd,
		}),
	)
	ground.Rotation().SetX(-90 * math.Pi / 180)
	// ground.SetReceiveShadow(true)
	ground.JSValue().Set("receiveShadow", true)
	c.scene.Add(ground.JSValue())

	// ambient := c.three.NewAmbientLight(0x666666)
	// c.scene.Add(ambient.JSValue())

	// directional := c.three.NewDirectionalLightFromColor(c.three.NewColor2(0.2, 0.1, 0.05), 1)
	directional := c.three.NewDirectionalLight(0xFFFFFF)
	directional.SetPosition(c.three.NewVector3(-15, 15, 15))
	// Shadow parameters
	// directional.JSValue().Set("castShadow", true)
	directional.SetCastShadow(true)
	directional.Shadow().MapSize().SetWidth(1024)
	directional.Shadow().MapSize().SetHeight(1024)
	directional.Shadow().Camera().SetRight(20)
	directional.Shadow().Camera().SetTop(20)
	directional.Shadow().Camera().SetLeft(-20)
	directional.Shadow().Camera().SetBottom(-20)

	// directional.Position().Normalize()
	c.scene.Add(directional.JSValue())

	c.renderer = c.three.NewWebGLRenderer(map[string]interface{}{
		"antialias": true,
		"alpha":     false,
	})
	c.renderer.JSValue().Get("shadowMap").Set("enabled", true)
	// c.renderer.ShadowMap().SetEnabled(true)

	c.outlineEffect = c.outlineEffectMod.New(c.renderer.JSValue())
	// c.outlineEffect = effects.NewOutlineEffect(c.outlineEffectMod, c.renderer)

	// Model specific Shadow parameters
	// c.renderer.ShadowMap().SetRenderSinglesSided(false)
	// c.renderer.ShadowMap().SetRenderReverseSided(false)
	directional.Shadow().SetBias(-0.001)

	// STATS
	c.stats = c.statsMod.New()

	// model
	var onProgress js.Func
	onProgress = js.FuncOf(func(this js.Value, args []js.Value) interface{} {

		xhr := args[0]
		if xhr.Get("lengthComputable").Bool() {
			loaded := xhr.Get("loaded").Float()
			total := xhr.Get("total").Float()

			percentComplete := loaded / total * 100
			log.Printf("%.1f %% downloaded.\n", percentComplete)

		}

		return nil
	})

	modelFile := "./assets/models/mmd/lat/Lat式ミクVer2.31_Normal.pmd"
	// modelFile := "./assets/models/mmd/miku/miku_v2.pmd"

	vmdFiles := "./assets/models/mmd/vmds/wavefile_v2.vmd"
	// vmdFiles := "./assets/models/mmd/vmds/Unknown.vmd"
	// cameraFiles := "./assets/models/mmd/vmds/wavefile_camera.vmd"

	c.helper = c.mmdAnimationHelperMod.New(map[string]interface{}{
		"afterglow": 2.0,
	})

	c.loader = c.mmdLoaderMod.New()
	c.loader.Call("loadWithAnimation", modelFile, vmdFiles,
		js.FuncOf(func(this js.Value, args []js.Value) interface{} {

			mmd := args[0]
			mesh := mmd.Get("mesh")
			// mesh.Get("position").Set("y", -10)
			mesh.Set("castShadow", true)
			mesh.Set("receiveShadow", true)

			c.scene.Add(mesh)

			c.helper.Call("add", mesh, map[string]interface{}{
				"animation": mmd.Get("animation"),
				"physics":   true,
			})

			return nil
		}),
		onProgress,
		js.Null(),
	)

	// Controls
	c.controls = c.orbitControlsMod.New(c.camera.JSValue(), c.renderer.DomElement())
	c.controls.Set("minDistance", 10)
	c.controls.Set("maxDistance", 100)

}

func (c *MikuView) animate(this js.Value, args []js.Value) interface{} {

	c.stats.Call("begin")

	c.helper.Call("update", c.clock.GetDelta())

	c.renderID = js.Global().Call("requestAnimationFrame", js.FuncOf(c.animate))
	c.outlineEffect.Call("render", c.scene.JSValue(), c.camera.JSValue())

	c.stats.Call("end")

	return nil
}

// OnResize ...
func (c *MikuView) OnResize(this js.Value, args []js.Value) interface{} {
	width := js.Global().Get("innerWidth").Float()
	height := js.Global().Get("innerHeight").Float()

	c.renderer.SetPixelRatio(js.Global().Get("devicePixelRatio").Float())
	c.renderer.SetSize(width, height, true)
	// c.outlineEffect.SetSize(width, height, true)

	c.camera.(*threejs.PerspectiveCamera).SetAspect(width / height)
	c.camera.(*threejs.PerspectiveCamera).UpdateProjectionMatrix()

	// println("Fire OnResize")

	return nil

}
