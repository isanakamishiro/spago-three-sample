package views

import (
	"app/frontend/lib/threejs"
	"app/frontend/lib/threejs/cameras"
	"app/frontend/lib/threejs/controls"
	"app/frontend/lib/threejs/datgui"
	"app/frontend/lib/threejs/geometries"
	"app/frontend/lib/threejs/loaders"
	"app/frontend/lib/threejs/materials"
	"app/frontend/lib/threejs/stats"
	"log"

	"syscall/js"

	"github.com/nobonobo/spago"
)

//go:generate spago generate -c Fundamental4 -p views threejs_fundamental4.html

// Fundamental4  ...
type Fundamental4 struct {
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

// NewFundamental4 is ...
func NewFundamental4() *Fundamental4 {

	return &Fundamental4{
		callbacks: threejs.NewCallbackRepository(),
	}
}

// Mount is ...
func (c *Fundamental4) Mount() {

	// shortcut of js.Funcs
	c.renderFunc = c.callbacks.Register(c.render)

	// initialize objects
	c.initSceneAndRenderer()

	// first render
	c.renderID = threejs.RequestAnimationFrame(c.renderFunc)

}

// Unmount ...
func (c *Fundamental4) Unmount() {

	// Cancel rendering
	threejs.CancelAnimationFrame(c.renderID)

	// Release all js.Funcs
	c.callbacks.ReleaseAll()

	// Dispose current rendering context
	c.renderer.Dispose()

}

func (c *Fundamental4) initSceneAndRenderer() {

	// Renderer
	renderer := threejs.NewWebGLRenderer(map[string]interface{}{
		"canvas":    js.Global().Get("document").Call("querySelector", "#c"),
		"antialias": true,
		// "alpha":     true,
	})
	c.renderer = renderer

	// GUI
	gui := datgui.NewGUIWithParameter(map[string]interface{}{
		"autoPlace": false,
	})
	// js.Global().Get("document").Call("querySelector", "#gui").Call("appendChild", gui.DomElement())
	c.gui = gui

	// Stats
	// stats := stats.NewStats()
	// js.Global().Get("document").Get("body").Call("appendChild", stats.DomElement())
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
	controls := controls.NewOrbitControls(c.camera, c.renderer.DomElement())
	controls.SetEnableDamping(true)
	controls.Target().Set2(0, 0, 0)
	controls.Update()
	c.controls = controls

	// Scene
	scene := threejs.NewScene()
	c.scene = scene

	// Lights
	// const (
	// 	lightColor     = threejs.ColorValue(0xffffff)
	// 	lightIntensity = threejs.LightIntensity(3)
	// )
	// light := lights.NewPointLight(lightColor, lightIntensity, 0, 1)
	// scene.AddLight(light)

	// Objects
	const (
		boxWidth  = 1
		boxHeight = 1
		boxDepth  = 1
		segments  = 1
	)
	geometry := geometries.NewBoxBufferGeometry(
		boxWidth,
		boxHeight,
		boxDepth,
		segments, segments, segments,
	)

	loadManager := threejs.NewLoadingManager()
	loader := loaders.NewTextureLoaderWithManager(loadManager)

	c.objects = make([]threejs.Object3D, 0, 1)

	// material := materials.NewMeshBasicMaterial(map[string]interface{}{
	// 	"map": loader.LoadSimply("https://threejsfundamentals.org/threejs/resources/images/wall.jpg"),
	// })

	var materialArray []threejs.Material = make([]threejs.Material, 6)
	materialArray[0] = materials.NewMeshBasicMaterial(map[string]interface{}{
		"map": loader.LoadSimply("https://threejsfundamentals.org/threejs/resources/images/flower-1.jpg"),
	})
	materialArray[1] = materials.NewMeshBasicMaterial(map[string]interface{}{
		"map": loader.LoadSimply("https://threejsfundamentals.org/threejs/resources/images/flower-2.jpg"),
	})
	materialArray[2] = materials.NewMeshBasicMaterial(map[string]interface{}{
		"map": loader.LoadSimply("https://threejsfundamentals.org/threejs/resources/images/flower-3.jpg"),
	})
	materialArray[3] = materials.NewMeshBasicMaterial(map[string]interface{}{
		"map": loader.LoadSimply("https://threejsfundamentals.org/threejs/resources/images/flower-4.jpg"),
	})
	materialArray[4] = materials.NewMeshBasicMaterial(map[string]interface{}{
		"map": loader.LoadSimply("https://threejsfundamentals.org/threejs/resources/images/flower-5.jpg"),
	})
	materialArray[5] = materials.NewMeshBasicMaterial(map[string]interface{}{
		"map": loader.LoadSimply("https://threejsfundamentals.org/threejs/resources/images/flower-6.jpg"),
	})

	loadManager.SetOnLoad(func() {
		log.Println("Texture loaded.")

		cube := threejs.NewMeshWithMultiMaterial(geometry, materialArray)
		scene.AddMesh(cube)
		c.objects = append(c.objects, cube)
	})

	loadManager.SetOnProgress(func(url string, itemsLoaded, itemsTotal int) {
		progress := float64(itemsLoaded) / float64(itemsTotal) * 100
		log.Printf("Progress : %.2f%%", progress)
	})

}

// resizeRendererToDisplaySize resizes render display size.
func (c *Fundamental4) resizeRendererToDisplaySize(renderer threejs.Renderer) (sizeChanged bool) {
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
func (c *Fundamental4) render(this js.Value, args []js.Value) interface{} {

	if sizeChanged := c.resizeRendererToDisplaySize(c.renderer); sizeChanged {
		canvas := c.renderer.DomElement()
		clientWidth := canvas.Get("clientWidth").Float()
		clientHeight := canvas.Get("clientHeight").Float()
		c.camera.(cameras.PerspectiveCamera).SetAspect(clientWidth / clientHeight)
		c.camera.(cameras.PerspectiveCamera).UpdateProjectionMatrix()
	}

	time := args[0].Float() * 0.001
	for i, obj := range c.objects {
		speed := 0.2 + float64(i)*0.1
		rot := time * speed

		obj.Rotation().SetX(rot)
		obj.Rotation().SetY(rot)
	}

	c.controls.Update()
	c.renderer.Render(c.scene, c.camera)

	// c.stats.Update()

	c.renderID = threejs.RequestAnimationFrame(c.renderFunc)

	return nil
}
