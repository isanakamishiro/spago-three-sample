package views

import (
	"app/frontend/lib/threejs"
	"app/frontend/lib/threejs/cameras"
	"app/frontend/lib/threejs/controls"
	"app/frontend/lib/threejs/datgui"
	"app/frontend/lib/threejs/fogs"
	"app/frontend/lib/threejs/geometries"
	"app/frontend/lib/threejs/lights"
	"app/frontend/lib/threejs/materials"
	"app/frontend/lib/threejs/stats"

	"syscall/js"

	"github.com/nobonobo/spago"
)

//go:generate spago generate -c Fundamental9 -p views threejs_fundamental9.html

// Fundamental9  ...
type Fundamental9 struct {
	spago.Core

	canvas js.Value
	camera threejs.Camera
	// cameraHelper *cameras.CameraHelper

	scene    threejs.Scene
	renderer threejs.Renderer
	objects  []threejs.Object3D

	control controls.OrbitControls

	gui   datgui.GUI
	stats stats.Stats

	callbacks threejs.CallbackRepository

	renderID   threejs.RenderID
	renderFunc js.Func
}

// NewFundamental9 is ...
func NewFundamental9() *Fundamental9 {

	return &Fundamental9{
		callbacks: threejs.NewCallbackRepository(),
	}
}

// Mount is ...
func (c *Fundamental9) Mount() {

	// shortcut of js.Funcs
	c.renderFunc = c.callbacks.Register(c.render)

	// initialize objects
	c.initSceneAndRenderer()

	// first render
	c.renderID = threejs.RequestAnimationFrame(c.renderFunc)

}

// Unmount ...
func (c *Fundamental9) Unmount() {

	// Cancel rendering
	threejs.CancelAnimationFrame(c.renderID)

	// Release all js.Funcs
	c.callbacks.ReleaseAll()

	// Dispose current rendering context
	c.renderer.Dispose()

}

func (c *Fundamental9) initSceneAndRenderer() {

	// Renderer
	renderer := threejs.NewWebGLRenderer(map[string]interface{}{
		"canvas":    js.Global().Get("document").Call("querySelector", "#c"),
		"antialias": true,
		// "alpha":     true,
	})
	renderer.ShadowMap().SetEnabled(true)
	c.renderer = renderer

	// GUI
	gui := datgui.NewGUIWithParameter(map[string]interface{}{
		"autoPlace": false,
	})
	// js.Global().Get("document").Call("querySelector", "#gui").Call("appendChild", gui.DomElement())
	c.gui = gui

	// Stats
	// stats := stats.NewStats()
	// js.Global().Get("document").Call("querySelector", "#main").Call("appendChild", stats.DomElement())
	// c.stats = stats

	// Camera
	const (
		fov    = 75
		aspect = 2
		near   = 0.1
		far    = 5
	)
	camera := cameras.NewPerspectiveCamera(fov, aspect, near, far)
	camera.Position().SetZ(2)
	c.camera = camera

	// Controls
	control := controls.NewOrbitControls(c.camera, c.renderer.DomElement())
	// control.Target().Set2(0, 5, 0)
	control.Update()
	c.control = control

	// Scene
	scene := threejs.NewScene()
	c.scene = scene

	// Fog
	const (
		fogNear  = 1
		fogFar   = 2
		fogColor = threejs.ColorValue(0xadd8e6)
	)
	fog := fogs.NewFog(threejs.NewColorFromColorValue(fogColor), fogNear, fogFar)
	scene.SetFog(fog)
	scene.SetBackgroundColor(threejs.NewColorFromColorValue(fogColor))

	// Objects : Sphere
	const (
		boxWidth  = 1
		boxHeight = 1
		boxDepth  = 1
		segment   = 1
	)
	geometry := geometries.NewBoxBufferGeometry(boxWidth, boxHeight, boxDepth, segment, segment, segment)

	c.objects = make([]threejs.Object3D, 0, 1)
	c.objects = append(c.objects, c.makeInstance(geometry, 0x44aa88, 0))
	c.objects = append(c.objects, c.makeInstance(geometry, 0x8844aa, -2))
	c.objects = append(c.objects, c.makeInstance(geometry, 0xaa8844, 2))

	// Light
	const (
		lightColor = threejs.ColorValue(0xffffff)
		intensity  = threejs.LightIntensity(1)
	)
	light := lights.NewDirectionalLight(lightColor, intensity)
	light.Position().Set2(-1, 2, 4)
	scene.AddLight(light)

	// build GUI

	params := map[string]interface{}{
		"color": fogColor.JSValue(),
	}
	gui.Add(fog.JSValue(), "near", float64(fogNear), float64(fogFar))
	gui.Add(fog.JSValue(), "far", float64(fogNear), float64(fogFar))
	gui.AddColor(params, "color").Name("color").OnChange(c.callbacks.Register(
		func(this js.Value, args []js.Value) interface{} {
			col := args[0].Int()

			fog.SetColor(threejs.NewColorFromColorValue(threejs.ColorValue(col)))
			scene.SetBackgroundColor(threejs.NewColorFromColorValue(threejs.ColorValue(col)))
			return nil
		},
	))
	// gui.Add(light.JSValue(), "intensity", 0, 2)
	// gui.Add(light.JSValue(), "distance", 0, 40)

}

func (c *Fundamental9) makeInstance(geometry threejs.Geometry, color int, x float64) threejs.Mesh {

	material := materials.NewMeshPhongMaterial(map[string]interface{}{
		"color": color,
	})

	cube := threejs.NewMesh(geometry, material)
	c.scene.AddMesh(cube)

	cube.Position().SetX(x)

	return cube
}

// resizeRendererToDisplaySize resizes render display size.
func (c *Fundamental9) resizeRendererToDisplaySize(renderer threejs.Renderer) (sizeChanged bool) {
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
func (c *Fundamental9) render(this js.Value, args []js.Value) interface{} {

	if sizeChanged := c.resizeRendererToDisplaySize(c.renderer); sizeChanged {
		canvas := c.renderer.DomElement()
		clientWidth := canvas.Get("clientWidth").Float()
		clientHeight := canvas.Get("clientHeight").Float()
		c.camera.(cameras.PerspectiveCamera).SetAspect(clientWidth / clientHeight)
		c.camera.(cameras.PerspectiveCamera).UpdateProjectionMatrix()
	}

	time := args[0].Float() * 0.001
	for i, obj := range c.objects {
		speed := 1 + float64(i)*0.1
		rot := time * speed

		obj.Rotation().SetX(rot)
		obj.Rotation().SetY(rot)
	}

	c.control.Update()
	c.renderer.Render(c.scene, c.camera)

	// c.stats.Update()

	c.renderID = threejs.RequestAnimationFrame(c.renderFunc)

	return nil
}
