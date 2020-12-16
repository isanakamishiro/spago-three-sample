package views

import (
	"app/frontend/lib/threejs"
	"syscall/js"

	"github.com/nobonobo/spago"
)

//go:generate spago generate -c Top -p views top.html

// Top  ...
type Top struct {
	spago.Core

	three    *threejs.ThreeJs
	camera   threejs.Camera
	scene    threejs.Scene
	mesh     threejs.Mesh
	renderer threejs.Renderer

	renderID     js.Value
	OnResizeFunc js.Func
}

// NewTop is ...
func NewTop(three *threejs.ThreeJs) *Top {

	camera := three.NewPerspectiveCamera(70, 4/3, 0.01, 10)
	camera.Position().SetZ(1)

	scene := three.NewScene()

	geometry := three.NewBoxGeometry(0.2, 0.2, 0.2, 1, 1, 1)
	material := three.NewMeshNormalMaterial(js.Null())
	mesh := three.NewMesh(geometry, material)

	scene.Add(mesh.JSValue())

	renderer := three.NewWebGLRenderer(map[string]interface{}{
		// "canvas":    js.Global().Get("document").Call("querySelector", "#myCanvas"),
		"antialias": true,
		"alpha":     true,
	})

	top := &Top{
		three:    three,
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

	c.camera.(*threejs.PerspectiveCamera).SetAspect(width / height)

	// println("Fire OnResize")

	return nil

}

// OnClick ...
func (c *Top) OnClick(ev js.Value) {

	spago.Rerender(c)
}
