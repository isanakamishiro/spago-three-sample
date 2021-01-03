package views

import (
	"app/frontend/lib/threejs"
	"app/frontend/lib/threejs/cameras"
	"app/frontend/lib/threejs/controls"
	"app/frontend/lib/threejs/effects"
	"app/frontend/lib/threejs/geometries"
	"app/frontend/lib/threejs/lights"
	"app/frontend/lib/threejs/loaders"
	"app/frontend/lib/threejs/loaders/mmdloaders"
	"app/frontend/lib/threejs/materials"
	"math"
	"syscall/js"

	"github.com/nobonobo/spago"
)

//go:generate spago generate -c MikuView -p views mikuview.html

// MikuView  ...
type MikuView struct {
	spago.Core

	camera        threejs.Camera
	scene         threejs.Scene
	mesh          threejs.Mesh
	renderer      threejs.Renderer
	outlineEffect *effects.OutlineEffect
	mmdHelper     *mmdloaders.MMDAnimationHelper

	clock   threejs.Clock
	control controls.OrbitControls

	callbacks threejs.CallbackRepository

	renderFunc js.Func
}

// NewMikuView is ...
func NewMikuView() *MikuView {

	// loadScript("./assets/threejs/ex/js/libs/ammo.wasm.js")

	view := &MikuView{
		callbacks: threejs.NewCallbackRepository(),
	}

	return view
}

func (c *MikuView) initSceneAndRenderer() {

	// Renderer
	renderer := threejs.NewWebGLRenderer(map[string]interface{}{
		"canvas":    js.Global().Get("document").Call("querySelector", "#c"),
		"antialias": true,
		// "alpha":     true,
	})
	renderer.ShadowMap().SetEnabled(true)
	renderer.SetPhysicallyCorrectLights(true)
	c.renderer = renderer

	// Clock
	c.clock = threejs.NewClock(true)

	// Camera
	camera := cameras.NewPerspectiveCamera(45, 4/3, 1, 500)
	camera.Position().SetZ(40)
	camera.Position().SetY(16)
	c.camera = camera

	// Scene
	scene := threejs.NewScene()
	scene.SetBackgroundColor(threejs.NewColorFromColorValue(0xffffff))
	c.scene = scene

	// Control
	control := controls.NewOrbitControls(camera, renderer.DomElement())
	control.Target().Set2(0, 12, 0)
	control.Update()
	c.control = control

	// Grid
	// gridHelper := helpers.NewPolarGridHelper(30, 10)
	// gridHelper.Position().SetY(0)
	// scene.Add(gridHelper)

	// axes := helpers.NewAxesHelper(5)
	// axes.Material().SetDepthTest(false)
	// scene.Add(axes)

	// Ground
	loader := loaders.NewTextureLoader()

	const planeSize = 40

	texture := loader.LoadSimply("https://threejsfundamentals.org/threejs/resources/images/checker.png")
	texture.SetWrapS(threejs.RepeatWrapping)
	texture.SetWrapT(threejs.RepeatWrapping)
	texture.SetMagFilter(threejs.NearestFilter)
	repeats := planeSize / 2.0
	texture.Repeat().Set2(repeats, repeats)

	planeGeo := geometries.NewPlaneBufferGeometry(planeSize, planeSize, 1, 1)
	planeMat := materials.NewMeshPhongMaterial(map[string]interface{}{
		"map":  texture,
		"side": threejs.DoubleSide.JSValue(),
	})
	planeMat.Color().SetRGB(1.5, 1.5, 1.5)
	planeMesh := threejs.NewMesh(planeGeo, planeMat)
	planeMesh.Rotation().SetX(math.Pi * -0.5)
	planeMesh.Scale().Set2(2, 2, 2)
	planeMesh.SetReceiveShadow(true)
	scene.Add(planeMesh)

	// Light

	// HemisphereLight
	{
		const (
			skyColor       = threejs.ColorValue(0xffffff)
			groundColor    = threejs.ColorValue(0xb97a20)
			lightIntensity = threejs.LightIntensity(2)
		)
		light := lights.NewHemisphereLight(skyColor, groundColor, lightIntensity)
		scene.AddLight(light)
	}

	// DirectionalLight
	{
		const (
			lightColor     = threejs.ColorValue(0xffffff)
			lightIntensity = threejs.LightIntensity(1)
			width          = 12.0
			height         = 4.0
		)
		light := lights.NewDirectionalLight(lightColor, lightIntensity)
		light.SetCastShadow(true)
		light.Position().Set2(-15, 40, 15)
		light.Target().Position().Set2(-4, 0, -4)

		light.Shadow().Camera().SetLeft(-20)
		light.Shadow().Camera().SetRight(20)
		light.Shadow().Camera().SetTop(20)
		light.Shadow().Camera().SetBottom(-20)

		scene.AddLight(light)
		scene.Add(light.Target())

		// cameraHelper := cameras.NewCameraHelper(light.Shadow().Camera())
		// scene.Add(cameraHelper)

		// helper := lights.NewDirectionalLightHelper(light)
		// scene.Add(helper)
	}

	// Effect
	outlineEffect := effects.NewOutlineEffect(renderer)
	c.outlineEffect = outlineEffect

	// model
	// modelFile := "./assets/models/mmd/lat/Lat式ミクVer2.31_Normal.pmd"
	modelFile := "./assets/models/mmd/miku/miku_v2.pmd"
	vmdFiles := []string{"./assets/models/mmd/vmds/wavefile_v2.vmd"}
	// vmdFiles := "./assets/models/mmd/vmds/Unknown.vmd"
	// cameraFiles := "./assets/models/mmd/vmds/wavefile_camera.vmd"

	mmdHelper := mmdloaders.NewMMDAnimationHelper(map[string]interface{}{
		"afterglow": 2.0,
	})
	c.mmdHelper = mmdHelper

	mmdLoader := mmdloaders.NewMMDLoader()
	mmdLoader.LoadWithAnimation(modelFile, vmdFiles, func(mesh mmdloaders.MMDMesh, animation mmdloaders.MMDAnimation) {

		mesh.SetCastShadow(true)
		// mesh.SetReceiveShadow(true)

		scene.AddMesh(mesh)
		mmdHelper.Add(mesh, map[string]interface{}{
			"animation": animation.JSValue(),
			"physics":   true,
		})
		c.mesh = mesh

	}, nil, nil)

}

// resizeRendererToDisplaySize resizes render display size.
func (c *MikuView) resizeRendererToDisplaySize(renderer threejs.Renderer) (sizeChanged bool) {
	canvas := renderer.DomElement()
	pixelRatio := js.Global().Get("devicePixelRatio").Float()
	width := canvas.Get("width").Int()
	height := canvas.Get("height").Int()
	clientWidth := int(canvas.Get("clientWidth").Float() * pixelRatio)
	clientHeight := int(canvas.Get("clientHeight").Float() * pixelRatio)

	needResize := (width != clientWidth || height != clientHeight)
	if needResize {
		renderer.SetSize(float64(clientWidth), float64(clientHeight), false)
	}
	return needResize
}

// render renders all objects
func (c *MikuView) render(this js.Value, args []js.Value) interface{} {

	if sizeChanged := c.resizeRendererToDisplaySize(c.renderer); sizeChanged {
		canvas := c.renderer.DomElement()
		clientWidth := canvas.Get("clientWidth").Float()
		clientHeight := canvas.Get("clientHeight").Float()
		c.camera.(cameras.PerspectiveCamera).SetAspect(clientWidth / clientHeight)
		c.camera.(cameras.PerspectiveCamera).UpdateProjectionMatrix()
	}

	c.control.Update()
	c.mmdHelper.Update(c.clock.Delta())

	c.renderer.Render(c.scene, c.camera)
	c.outlineEffect.Render(c.scene, c.camera)

	// c.renderID = threejs.RequestAnimationFrame(c.renderFunc)

	return nil
}

// Mount is ...
func (c *MikuView) Mount() {

	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		defer fn.Release()

		// initialize objects
		c.initSceneAndRenderer()

		// shortcut of js.Funcs
		c.renderFunc = c.callbacks.Register(c.render)
		c.renderer.SetAnimationLoop(c.renderFunc)

		// first render
		// c.renderID = threejs.RequestAnimationFrame(c.renderFunc)

		return nil
	})

	// Ammo functionが呼ばれていない状態 = Function, コール後、Object
	if js.Global().Get("Ammo").Type() == js.TypeFunction {
		js.Global().Call("Ammo").Call("then", fn)
	} else {
		fn.Invoke()
	}

}

// Unmount ...
func (c *MikuView) Unmount() {

	// Cancel rendering
	// threejs.CancelAnimationFrame(c.renderID)
	c.renderer.CancelAnimationLoop()

	// Release all js.Funcs
	c.callbacks.ReleaseAll()

	// Dispose current rendering context
	c.renderer.Dispose()

}
