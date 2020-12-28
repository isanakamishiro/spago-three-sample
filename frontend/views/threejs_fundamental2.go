package views

import (
	"app/frontend/lib/threejs"
	"app/frontend/lib/threejs/cameras"
	"app/frontend/lib/threejs/controls"
	"app/frontend/lib/threejs/datgui"
	"app/frontend/lib/threejs/geometries"
	"app/frontend/lib/threejs/lights"
	"app/frontend/lib/threejs/materials"
	"app/frontend/lib/threejs/stats"
	"fmt"
	"syscall/js"

	"github.com/nobonobo/spago"
)

//go:generate spago generate -c Fundamental2 -p views threejs_fundamental2.html

// Fundamental2  ...
type Fundamental2 struct {
	spago.Core

	camera   threejs.Camera
	scene    threejs.Scene
	renderer threejs.Renderer

	controls controls.OrbitControls
	gui      datgui.GUI
	stats    stats.Stats

	callbacks threejs.CallbackRepository

	renderID          threejs.RenderID
	renderRequestFunc js.Func
	renderFunc        js.Func

	renderRequested bool
}

// NewFundamental2 is ...
func NewFundamental2() *Fundamental2 {

	return &Fundamental2{
		callbacks: threejs.NewCallbackRepository(),
	}
}

// Mount is ...
func (c *Fundamental2) Mount() {

	// shortcut of js.Funcs
	c.renderFunc = c.callbacks.Register(c.animate)
	c.renderRequestFunc = c.callbacks.Register(c.requestRenderIfNotRequested)

	// initialize objects
	c.initSceneAndRenderer()

	// Add event listeners
	c.controls.AddEventListener("change", c.renderRequestFunc)
	js.Global().Call("addEventListener", "resize", c.renderRequestFunc)

	// first render
	c.requestRenderIfNotRequested(js.Null(), nil)
}

// Unmount ...
func (c *Fundamental2) Unmount() {

	// Remove event listeners
	c.controls.RemoveEventListener("change", c.renderRequestFunc)
	js.Global().Call("removeEventListener", "resize", c.renderRequestFunc)

	// Release all js.Funcs
	c.callbacks.ReleaseAll()

}

func (c *Fundamental2) initSceneAndRenderer() {

	// Renderer
	renderer := threejs.NewWebGLRenderer(map[string]interface{}{
		"canvas":    js.Global().Get("document").Call("querySelector", "#c"),
		"antialias": true,
		// "alpha":     true,
	})
	c.renderer = renderer

	// Scene
	scene := threejs.NewScene()
	scene.SetBackground(threejs.NewColorFromColorValue(0xAAAAAA).JSValue())
	c.scene = scene

	// Camera
	const (
		fov    = 75
		aspect = 2
		near   = 0.1
		far    = 1000
	)
	camera := cameras.NewPerspectiveCamera(fov, aspect, near, far)
	camera.Position().SetZ(2)
	c.camera = camera

	// Controls
	controls := controls.NewOrbitControls(camera, renderer.DomElement())
	controls.SetEnableDamping(true)
	controls.Target().Set2(0, 0, 0)
	controls.Update()
	c.controls = controls

	// GUI
	gui := datgui.NewGUI()
	c.gui = gui

	// Stats
	stats := stats.NewStats()
	js.Global().Get("document").Get("body").Call("appendChild", stats.DomElement())
	c.stats = stats

	// Light
	const (
		lightColor = 0xFFFFFF
		intensity  = 1
	)
	light := lights.NewDirectionalLight(lightColor, intensity)
	light.Position().Set2(01, 2, 4)
	scene.AddLight(light)

	// Box
	const (
		boxWidth  = 1
		boxHeight = 1
		boxDepth  = 1
	)
	geometry := geometries.NewBoxBufferGeometry(boxWidth, boxHeight, boxDepth, 1, 1, 1)

	c.makeInstance(geometry, 0x44aa88, 0)
	c.makeInstance(geometry, 0x8844aa, -2)
	c.makeInstance(geometry, 0xaa8844, 2)

}

// animate renders all objects
func (c *Fundamental2) animate(this js.Value, args []js.Value) interface{} {
	c.renderRequested = false

	if sizeChanged := c.resizeRendererToDisplaySize(c.renderer); sizeChanged {
		canvas := c.renderer.DomElement()
		clientWidth := canvas.Get("clientWidth").Float()
		clientHeight := canvas.Get("clientHeight").Float()
		c.camera.(cameras.PerspectiveCamera).SetAspect(clientWidth / clientHeight)
		c.camera.(cameras.PerspectiveCamera).UpdateProjectionMatrix()
	}

	c.controls.Update()
	c.renderer.Render(c.scene, c.camera)

	c.stats.Update()

	return nil
}

// requestRenderIfNotRequested publishes render request
func (c *Fundamental2) requestRenderIfNotRequested(this js.Value, args []js.Value) interface{} {
	if !c.renderRequested {
		c.renderRequested = true
		threejs.RequestAnimationFrame(c.renderFunc)
	}

	return nil
}

// makeInstance creates mesh with parameters
func (c *Fundamental2) makeInstance(geometry threejs.Geometry, color int, x float64) threejs.Mesh {

	material := materials.NewMeshPhongMaterial(map[string]interface{}{
		"color": color,
	})

	cube := threejs.NewMesh(geometry, material)
	c.scene.AddMesh(cube)

	cube.Position().SetX(x)

	folder := c.gui.AddFolder(fmt.Sprintf("Cube%v", x))
	folder.Add(cube.Scale().JSValue(), "x", 0.1, 1.5).Name("scale x").OnChange(c.renderRequestFunc)

	obj := map[string]interface{}{
		"color": material.Color().Hex(),
	}
	folder.AddColor(obj, "color").Name("Color").OnChange(
		c.callbacks.Register(func(this js.Value, args []js.Value) interface{} {

			color := args[0].Int()
			material.Color().SetHex(color)

			c.requestRenderIfNotRequested(js.Null(), nil)

			return nil
		}))

	folder.Open()

	return cube
}

// resizeRendererToDisplaySize resizes render display size.
func (c *Fundamental2) resizeRendererToDisplaySize(renderer threejs.Renderer) (sizeChanged bool) {
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
