package views

import (
	"app/frontend/lib/threejs"
	"app/frontend/lib/threejs/cameras"
	"app/frontend/lib/threejs/geometries"
	"app/frontend/lib/threejs/lights"
	"app/frontend/lib/threejs/materials"
	"syscall/js"

	"github.com/nobonobo/spago"
)

//go:generate spago generate -c Fundamental1 -p views threejs_fundamental1.html

// Fundamental1  ...
type Fundamental1 struct {
	spago.Core

	camera   threejs.Camera
	scene    threejs.Scene
	cube     threejs.Mesh
	cubes    []threejs.Mesh
	renderer threejs.Renderer

	callbacks threejs.CallbackRepository

	renderID   threejs.RenderID
	renderFunc js.Func
}

// NewFundamental1 is ...
func NewFundamental1() *Fundamental1 {

	return &Fundamental1{
		callbacks: threejs.NewCallbackRepository(),
	}
}

func (c *Fundamental1) initSceneAndRenderer() {
	// Renderer
	renderer := threejs.NewWebGLRenderer(map[string]interface{}{
		"canvas":    js.Global().Get("document").Call("querySelector", "#c"),
		"antialias": true,
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

	cubes := []threejs.Mesh{
		c.makeInstance(geometry, 0x44aa88, 0),
		c.makeInstance(geometry, 0x8844aa, -2),
		c.makeInstance(geometry, 0xaa8844, 2),
	}
	c.cubes = cubes

}

// Mount is ...
func (c *Fundamental1) Mount() {

	// shortcut of js.Funcs
	c.renderFunc = c.callbacks.Register(c.animate)

	// initialize objects
	c.initSceneAndRenderer()

	// first render
	c.renderID = threejs.RequestAnimationFrame(c.renderFunc)

}

// Unmount ...
func (c *Fundamental1) Unmount() {

	// Cancel rendering
	threejs.CancelAnimationFrame(c.renderID)

	// Release all js.Funcs
	c.callbacks.ReleaseAll()

	// Dispose current rendering context
	c.renderer.Dispose()

}

func (c *Fundamental1) animate(this js.Value, args []js.Value) interface{} {

	time := args[0].Float()
	time = time * 0.001 // convert time to seconds

	if sizeChanged := resizeRendererToDisplaySize(c.renderer); sizeChanged {
		canvas := c.renderer.DomElement()
		clientWidth := canvas.Get("clientWidth").Float()
		clientHeight := canvas.Get("clientHeight").Float()
		c.camera.(cameras.PerspectiveCamera).SetAspect(clientWidth / clientHeight)
		c.camera.(cameras.PerspectiveCamera).UpdateProjectionMatrix()
	}

	for ndx, cube := range c.cubes {
		speed := 1.0 + float64(ndx)*0.1
		rot := time * speed
		cube.Rotation().SetX(rot)
		cube.Rotation().SetY(rot)
	}

	c.renderer.Render(c.scene, c.camera)

	c.renderID = threejs.RequestAnimationFrame(c.renderFunc)

	return nil
}

func (c *Fundamental1) makeInstance(geometry threejs.Geometry, color int, x float64) threejs.Mesh {

	material := materials.NewMeshPhongMaterial(map[string]interface{}{
		"color": color,
	})

	cube := threejs.NewMesh(geometry, material)
	c.scene.AddMesh(cube)

	cube.Position().SetX(x)

	return cube
}

func resizeRendererToDisplaySize(renderer threejs.Renderer) (sizeChanged bool) {
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
