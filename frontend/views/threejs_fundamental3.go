package views

import (
	"app/frontend/lib/threejs"
	"app/frontend/lib/threejs/cameras"
	"app/frontend/lib/threejs/controls"
	"app/frontend/lib/threejs/datgui"
	"app/frontend/lib/threejs/geometries"
	"app/frontend/lib/threejs/helpers"
	"app/frontend/lib/threejs/lights"
	"app/frontend/lib/threejs/materials"
	"app/frontend/lib/threejs/stats"
	"syscall/js"

	"github.com/nobonobo/spago"
)

//go:generate spago generate -c Fundamental3 -p views threejs_fundamental3.html

// Fundamental3  ...
type Fundamental3 struct {
	spago.Core

	camera   threejs.Camera
	scene    threejs.Scene
	renderer threejs.Renderer
	objects  []threejs.Object3D

	controls controls.OrbitControls
	gui      datgui.GUI
	stats    stats.Stats

	callbacks threejs.CallbackRepository

	renderID   threejs.RenderID
	renderFunc js.Func

	renderRequested bool
}

// NewFundamental3 is ...
func NewFundamental3() *Fundamental3 {

	return &Fundamental3{
		callbacks: threejs.NewCallbackRepository(),
	}
}

// Mount is ...
func (c *Fundamental3) Mount() {

	// shortcut of js.Funcs
	c.renderFunc = c.callbacks.Register(c.render)

	// initialize objects
	c.initSceneAndRenderer()

	// Add event listeners
	// js.Global().Call("addEventListener", "resize", c.renderRequestFunc)

	// first render
	threejs.RequestAnimationFrame(c.renderFunc)

}

// Unmount ...
func (c *Fundamental3) Unmount() {

	// Remove event listeners
	// js.Global().Call("removeEventListener", "resize", c.renderRequestFunc)

	// Release all js.Funcs
	c.callbacks.ReleaseAll()

}

func (c *Fundamental3) initSceneAndRenderer() {

	// Renderer
	renderer := threejs.NewWebGLRenderer(map[string]interface{}{
		"canvas":    js.Global().Get("document").Call("querySelector", "#c"),
		"antialias": true,
		// "alpha":     true,
	})
	c.renderer = renderer

	// GUI
	gui := datgui.NewGUI()
	c.gui = gui

	// Stats
	stats := stats.NewStats()
	js.Global().Get("document").Get("body").Call("appendChild", stats.DomElement())
	c.stats = stats

	// Camera
	const (
		fov    = 40
		aspect = 2
		near   = 0.1
		far    = 1000
	)
	camera := cameras.NewPerspectiveCamera(fov, aspect, near, far)
	camera.Position().Set2(0, 50, 0)
	camera.Up().Set2(0, 0, 1)
	camera.LookAtXYZ(0, 0, 0)
	c.camera = camera

	// Controls
	controls := controls.NewOrbitControls(c.camera, c.renderer.DomElement())
	controls.SetEnableDamping(true)
	controls.Target().Set2(0, 0, 0)
	controls.Update()
	c.controls = controls

	// Scene
	scene := threejs.NewScene()
	c.scene = scene

	// Lights
	const (
		lightColor     = threejs.ColorValue(0xffffff)
		lightIntensity = threejs.LightIntensity(3)
	)
	light := lights.NewPointLight(lightColor, lightIntensity, 0, 1)
	scene.AddLight(light)

	// Objects
	const (
		radius         = 1
		widthSegments  = 6
		heightSegments = 6
	)
	sphereGeometry := geometries.NewSphereBufferGeometry(radius, widthSegments, heightSegments)

	// objects : SolarSystem
	solarSystem := threejs.NewObject3D()
	scene.Add(solarSystem)

	// objects : Sun
	sunMaterial := materials.NewMeshPhongMaterial(map[string]interface{}{
		"emissive": 0xffff00,
	})
	sunMesh := threejs.NewMesh(sphereGeometry, sunMaterial)
	sunMesh.Scale().Set2(5, 5, 5)
	solarSystem.Add(sunMesh)

	// objects : EarthOrbit
	earthOrbit := threejs.NewObject3D()
	earthOrbit.Position().SetX(10)
	solarSystem.Add(earthOrbit)

	// objects : Earth
	earthMaterial := materials.NewMeshPhongMaterial(map[string]interface{}{
		"color":    0x2233ff,
		"emissive": 0x112244,
	})
	earthMesh := threejs.NewMesh(sphereGeometry, earthMaterial)
	earthOrbit.Add(earthMesh)

	// objects : MoonOrbit
	moonOrbit := threejs.NewObject3D()
	moonOrbit.Position().SetX(2)
	earthOrbit.Add(moonOrbit)

	// objects : Moon
	moonMaterial := materials.NewMeshPhongMaterial(map[string]interface{}{
		"color":    0x888888,
		"emissive": 0x222222,
	})
	moonMesh := threejs.NewMesh(sphereGeometry, moonMaterial)
	moonMesh.Scale().Set2(0.5, 0.5, 0.5)
	moonOrbit.Add(moonMesh)

	// objects : push to repository
	c.objects = make([]threejs.Object3D, 0, 6)
	c.objects = append(c.objects, solarSystem)
	c.objects = append(c.objects, sunMesh)
	c.objects = append(c.objects, earthOrbit)
	c.objects = append(c.objects, earthMesh)
	c.objects = append(c.objects, moonOrbit)
	c.objects = append(c.objects, moonMesh)

	// objects : add an AxesHelper to each node
	// for _, node := range c.objects {
	// 	axes := helpers.NewAxesHelper(1)
	// 	axes.Material().SetDepthTest(false)
	// 	axes.SetRenderOrder(1)

	// 	node.Add(axes)
	// }

	c.makeAxisGrid(solarSystem, "solarSystem", 26)
	c.makeAxisGrid(sunMesh, "sunMesh", 10)
	c.makeAxisGrid(earthOrbit, "earthOrbit", 10)
	c.makeAxisGrid(earthMesh, "earthMesh", 10)
	c.makeAxisGrid(moonOrbit, "moonOrbit", 10)
	c.makeAxisGrid(moonMesh, "moonMesh", 10)

}

func (c *Fundamental3) makeAxisGrid(node threejs.Object3D, label string, units float64) {

	axes := helpers.NewAxesHelper(1)
	axes.Material().SetDepthTest(false)
	axes.SetRenderOrder(2)
	axes.SetVisible(false)
	node.Add(axes)

	grid := helpers.NewGridHelper(units, units)
	grid.Material().SetDepthTest(false)
	grid.SetRenderOrder(1)
	grid.SetVisible(false)
	node.Add(grid)

	obj := js.ValueOf(map[string]interface{}{
		"visible": false,
	})

	c.gui.AddSimply(obj, "visible").Name(label).OnChange(c.callbacks.Register(
		func(this js.Value, args []js.Value) interface{} {
			visible := args[0].Bool()

			axes.SetVisible(visible)
			grid.SetVisible(visible)

			return nil
		},
	))
}

// resizeRendererToDisplaySize resizes render display size.
func (c *Fundamental3) resizeRendererToDisplaySize(renderer threejs.Renderer) (sizeChanged bool) {
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
func (c *Fundamental3) render(this js.Value, args []js.Value) interface{} {

	if sizeChanged := c.resizeRendererToDisplaySize(c.renderer); sizeChanged {
		canvas := c.renderer.DomElement()
		clientWidth := canvas.Get("clientWidth").Float()
		clientHeight := canvas.Get("clientHeight").Float()
		c.camera.(cameras.PerspectiveCamera).SetAspect(clientWidth / clientHeight)
		c.camera.(cameras.PerspectiveCamera).UpdateProjectionMatrix()
	}

	time := args[0].Float() * 0.001
	for _, obj := range c.objects {
		obj.Rotation().SetY(time)
	}

	c.controls.Update()
	c.renderer.Render(c.scene, c.camera)

	c.stats.Update()

	threejs.RequestAnimationFrame(c.renderFunc)

	return nil
}
