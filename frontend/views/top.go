package views

import (
	"app/frontend/lib/threejs"
	"app/frontend/lib/threejs/cameras"
	"app/frontend/lib/threejs/geometries"
	"app/frontend/lib/threejs/materials"
	"syscall/js"

	"github.com/nobonobo/spago"
)

//go:generate spago generate -c Top -p views top.html

// Top  ...
type Top struct {
	spago.Core

	camera   threejs.Camera
	scene    threejs.Scene
	mesh     threejs.Mesh
	renderer threejs.Renderer

	callbacks threejs.CallbackRepository

	renderID   threejs.RenderID
	renderFunc js.Func
}

// NewTop is ...
func NewTop() *Top {

	top := &Top{
		callbacks: threejs.NewCallbackRepository(),
	}
	return top
}

func (c *Top) initRenderScene() {

	// Renderer
	renderer := threejs.NewWebGLRenderer(map[string]interface{}{
		"canvas":    js.Global().Get("document").Call("querySelector", "#c"),
		"antialias": true,
		"alpha":     true,
	})
	c.renderer = renderer

	// Scene
	scene := threejs.NewScene()
	c.scene = scene

	// Camera
	const (
		fov    = 40
		aspect = 4 / 3
		near   = 0.1
		far    = 50
	)
	camera := cameras.NewPerspectiveCamera(fov, aspect, near, far)
	camera.Position().SetY(8)
	camera.Up().SetZ(1)
	camera.LookAtXYZ(0, 0, 0)
	c.camera = camera

	// geometry
	const (
		radius         = 1
		widthSegments  = 2
		heightSegments = 2
	)
	geometry := geometries.NewBoxBufferGeometry(radius, radius, radius, widthSegments, heightSegments, heightSegments)
	material := materials.NewMeshBasicMaterial(map[string]interface{}{
		// "emissive": 0xFFFF00,
		"color": 0xdddddd,
	})

	mesh := threejs.NewMesh(geometry, material)
	scene.AddMesh(mesh)
	c.mesh = mesh
}

// Mount is ...
func (c *Top) Mount() {

	// shortcut of js.Funcs
	c.renderFunc = c.callbacks.Register(c.render)

	// initialize objects
	c.initRenderScene()

	// first render
	c.renderID = threejs.RequestAnimationFrame(c.renderFunc)

}

// Unmount ...
func (c *Top) Unmount() {

	// Cancel rendering
	threejs.CancelAnimationFrame(c.renderID)

	// Release all js.Funcs
	c.callbacks.ReleaseAll()

}

// resizeRendererToDisplaySize resizes render display size.
func (c *Top) resizeRendererToDisplaySize(renderer threejs.Renderer) (sizeChanged bool) {
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
func (c *Top) render(this js.Value, args []js.Value) interface{} {

	if sizeChanged := c.resizeRendererToDisplaySize(c.renderer); sizeChanged {
		canvas := c.renderer.DomElement()
		clientWidth := canvas.Get("clientWidth").Float()
		clientHeight := canvas.Get("clientHeight").Float()
		c.camera.(cameras.PerspectiveCamera).SetAspect(clientWidth / clientHeight)
		c.camera.(cameras.PerspectiveCamera).UpdateProjectionMatrix()
	}

	time := args[0].Float() * 0.001
	speed := 1.0
	rot := time * speed

	c.mesh.Rotation().SetX(rot)
	c.mesh.Rotation().SetY(rot)

	c.renderer.Render(c.scene, c.camera)

	c.renderID = threejs.RequestAnimationFrame(c.renderFunc)

	return nil
}
