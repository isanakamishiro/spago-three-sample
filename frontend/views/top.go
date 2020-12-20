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

//go:generate spago generate -c Top -p views top.html

// Top  ...
type Top struct {
	spago.Core

	camera   threejs.Camera
	scene    threejs.Scene
	mesh     threejs.Mesh
	renderer threejs.Renderer

	renderID     js.Value
	OnResizeFunc js.Func
}

// NewTop is ...
func NewTop() *Top {

	// Renderer
	renderer := threejs.NewWebGLRenderer(map[string]interface{}{
		// "canvas":    js.Global().Get("document").Call("querySelector", "#myCanvas"),
		"antialias": true,
		"alpha":     true,
	})

	// Scene
	scene := threejs.NewScene()

	// Camera
	const (
		fov    = 40
		aspect = 4 / 3
		near   = 0.1
		far    = 1000
	)
	camera := cameras.NewPerspectiveCamera(fov, aspect, near, far)
	camera.Position().SetY(20)
	camera.Up().SetZ(1)
	camera.LookAt(threejs.NewVector3(0, 0, 0))

	// Light
	const (
		color     = 0xFFFFFF
		intensity = 3
	)
	light := lights.NewPointLight(color, intensity, 0, 1)
	scene.AddLight(light)

	// geometry
	const (
		radius         = 1
		widthSegments  = 6
		heightSegments = 6
	)
	geometry := geometries.NewBoxBufferGeometry(radius, radius, radius, widthSegments, heightSegments, heightSegments)
	material := materials.NewMeshPhongMaterial(map[string]interface{}{
		"emissive": 0xFFFF00,
		// "color":    0xdddddd,
	})

	mesh := threejs.NewMesh(geometry, material)

	scene.AddMesh(mesh)

	top := &Top{
		camera:   camera,
		scene:    scene,
		mesh:     mesh,
		renderer: renderer,
	}

	return top
}

// Mount is ...
func (c *Top) Mount() {

	// c.renderer = c.three.NewWebGLRenderer(map[string]interface{}{
	// 	// "canvas":    js.Global().Get("document").Call("querySelector", "#myCanvas"),
	// 	"antialias": true,
	// })

	js.Global().Get("document").Get("body").Call("appendChild", c.renderer.DomElement())

	c.OnResizeFunc = js.FuncOf(c.OnResize)
	js.Global().Call("addEventListener", "resize", c.OnResizeFunc)

	c.OnResize(js.Null(), nil)
	c.animate(js.Null(), nil)

	// println("Mount!!")

}

// Unmount ...
func (c *Top) Unmount() {

	js.Global().Call("cancelAnimationFrame", c.renderID)
	js.Global().Call("removeEventListener", "resize", c.OnResizeFunc)

	// println("Unmount!!")

}

func (c *Top) animate(this js.Value, args []js.Value) interface{} {

	c.renderID = js.Global().Call("requestAnimationFrame", js.FuncOf(c.animate))
	c.mesh.Rotation().SetX(c.mesh.Rotation().X() + 0.005)
	c.mesh.Rotation().SetY(c.mesh.Rotation().Y() + 0.01)
	c.renderer.Render(c.scene, c.camera)

	return nil
}

// OnResize ...
func (c *Top) OnResize(this js.Value, args []js.Value) interface{} {
	width := js.Global().Get("innerWidth").Float()
	height := js.Global().Get("innerHeight").Float()

	c.renderer.SetPixelRatio(js.Global().Get("devicePixelRatio").Float())
	c.renderer.SetSize(width, height, true)

	c.camera.(cameras.PerspectiveCamera).SetAspect(width / height)

	// println("Fire OnResize")

	return nil

}

// OnClick ...
func (c *Top) OnClick(ev js.Value) {

	spago.Rerender(c)
}
