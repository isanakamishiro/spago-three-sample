package views

import (
	"app/frontend/lib/threejs"
	"app/frontend/lib/threejs/cameras"
	"app/frontend/lib/threejs/controls"
	"app/frontend/lib/threejs/datgui"
	"app/frontend/lib/threejs/geometries"
	"app/frontend/lib/threejs/lights"
	"app/frontend/lib/threejs/loaders"
	"app/frontend/lib/threejs/materials"
	"app/frontend/lib/threejs/stats"
	"math"

	"syscall/js"

	"github.com/nobonobo/spago"
)

//go:generate spago generate -c Fundamental5 -p views threejs_fundamental5.html

// Fundamental5  ...
type Fundamental5 struct {
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

// NewFundamental5 is ...
func NewFundamental5() *Fundamental5 {

	return &Fundamental5{
		callbacks: threejs.NewCallbackRepository(),
	}
}

// Mount is ...
func (c *Fundamental5) Mount() {

	// shortcut of js.Funcs
	c.renderFunc = c.callbacks.Register(c.render)

	// initialize objects
	c.initSceneAndRenderer()

	// first render
	c.renderID = threejs.RequestAnimationFrame(c.renderFunc)

}

// Unmount ...
func (c *Fundamental5) Unmount() {

	// Cancel rendering
	threejs.CancelAnimationFrame(c.renderID)

	// Release all js.Funcs
	c.callbacks.ReleaseAll()

	// Dispose current rendering context
	c.renderer.Dispose()

}

func (c *Fundamental5) initSceneAndRenderer() {

	// Renderer
	renderer := threejs.NewWebGLRenderer(map[string]interface{}{
		"canvas":    js.Global().Get("document").Call("querySelector", "#c"),
		"antialias": true,
		// "alpha":     true,
	})
	renderer.SetPhysicallyCorrectLights(true)
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
		fov    = 45
		aspect = 2
		near   = 0.1
		far    = 100
	)
	camera := cameras.NewPerspectiveCamera(fov, aspect, near, far)
	camera.Position().Set2(0, 10, 20)
	c.camera = camera

	// Controls
	controls := controls.NewOrbitControls(c.camera, c.renderer.DomElement())
	controls.Target().Set2(0, 5, 0)
	controls.Update()
	c.controls = controls

	// Scene
	scene := threejs.NewScene()
	c.scene = scene

	// Texture
	const planeSize = 40
	loader := loaders.NewTextureLoader()
	texture := loader.LoadSimply("https://threejsfundamentals.org/threejs/resources/images/checker.png")
	texture.SetWrapS(threejs.RepeatWrapping)
	texture.SetWrapT(threejs.RepeatWrapping)
	texture.SetMagFilter(threejs.NearestFilter)
	repeats := planeSize / 2.0
	texture.Repeat().Set2(repeats, repeats)

	// Objects : Plane
	planeGeo := geometries.NewPlaneBufferGeometry(planeSize, planeSize, 1, 1)
	// planeMat := materials.NewMeshPhongMaterial(map[string]interface{}{
	planeMat := materials.NewMeshStandardMaterial(map[string]interface{}{
		"map":  texture,
		"side": threejs.DoubleSide.JSValue(),
	})
	planeMesh := threejs.NewMesh(planeGeo, planeMat)
	planeMesh.Rotation().SetX(math.Pi * -0.5)
	scene.Add(planeMesh)

	// Objects : Cube
	const cubeSize = 4.0
	cubeGeo := geometries.NewBoxBufferGeometry(cubeSize, cubeSize, cubeSize, 1, 1, 1)
	// cubeMat := materials.NewMeshPhongMaterial(map[string]interface{}{
	cubeMat := materials.NewMeshStandardMaterial(map[string]interface{}{
		"color": "#8AC",
	})
	cubeMesh := threejs.NewMesh(cubeGeo, cubeMat)
	cubeMesh.Position().Set2(cubeSize+1, cubeSize/2, 0)
	scene.Add(cubeMesh)

	// Objects : Sphere
	const (
		sphereRadius          = 3.0
		sphereWidthDivisions  = 32
		sphereHeightDivisions = 16
	)
	sphereGeo := geometries.NewSphereBufferGeometry(sphereRadius, sphereWidthDivisions, sphereHeightDivisions)
	// sphereMat := materials.NewMeshPhongMaterial(map[string]interface{}{
	sphereMat := materials.NewMeshStandardMaterial(map[string]interface{}{
		"color": "#CA8",
	})
	sphereMesh := threejs.NewMesh(sphereGeo, sphereMat)
	sphereMesh.Position().Set2(-sphereRadius-1, sphereRadius+2, 0)
	scene.Add(sphereMesh)

	// Lights

	// Ambient Light
	// const (
	// 	lightColor     = threejs.ColorValue(0xffffff)
	// 	lightIntensity = threejs.LightIntensity(1)
	// )
	// light := lights.NewAmbientLight(lightColor, lightIntensity)
	// scene.AddLight(light)

	// Hemisphere Light
	// const (
	// 	skyColor       = threejs.ColorValue(0xB1E1FF)
	// 	groundColor    = threejs.ColorValue(0xB97A20)
	// 	lightIntensity = threejs.LightIntensity(1)
	// )
	// light := lights.NewHemisphereLight(skyColor, groundColor, lightIntensity)
	// scene.AddLight(light)

	// Directional Light / Point Light / Spot Light
	const (
		lightColor     = threejs.ColorValue(0xffffff)
		lightIntensity = threejs.LightIntensity(1)
		width          = 12.0
		height         = 4.0
	)
	// light := lights.NewDirectionalLight(lightColor, lightIntensity)
	light := lights.NewPointLight(lightColor, lightIntensity, 0, 1)
	// light := lights.NewSpotLight(lightColor, lightIntensity)
	// light := exlights.NewRectAreaLight(lightColor, lightIntensity, width, height)
	light.SetPower(800)
	light.SetDecay(2)
	light.SetDistance(0)
	light.Position().Set2(0, 10, 0)
	// light.Rotation().SetX(-1.5708)
	// light.Target().Position().Set2(-5, 0, 0)
	scene.AddLight(light)
	// scene.Add(light.Target())

	// helper := lights.NewDirectionalLightHelper(light)
	helper := lights.NewPointLightHelper(light)
	// helper := lights.NewSpotLightHelper(light)
	// helper := exlights.NewRectAreaLightHelper(light)
	// light.Add(helper)
	scene.Add(helper)

	// Objects
	c.objects = make([]threejs.Object3D, 0, 1)

	// build GUI
	jsfnUpdate := c.callbacks.Register(func(this js.Value, args []js.Value) interface{} {
		// light.Target().UpdateMatrixWorld(true)
		helper.Update()

		return nil
	})
	jsfnUpdate.Invoke()

	params := map[string]interface{}{
		"color": int(lightColor),
	}
	gui.AddColor(params, "color").Name("color").OnChange(c.callbacks.Register(
		func(this js.Value, args []js.Value) interface{} {
			col := args[0].Int()

			light.SetColor(threejs.NewColorFromColorValue(threejs.ColorValue(col)))
			return nil
		},
	))
	// gui.Add(light.JSValue(), "intensity", 0, 10)
	// gui.Add(light.JSValue(), "distance", 0, 40).OnChange(jsfnUpdate)
	// gui.Add(light.JSValue(), "angle", -3.14, 3.14).OnChange(jsfnUpdate)
	// gui.Add(light.JSValue(), "penumbra", 0.0, 1.0).OnChange(jsfnUpdate)
	// gui.Add(light.JSValue(), "width", 0, 20).OnChange(jsfnUpdate)
	// gui.Add(light.JSValue(), "height", 0, 20).OnChange(jsfnUpdate)
	// gui.Add(light.Rotation().JSValue(), "x", -math.Pi, math.Pi).Name("x rotation").OnChange(jsfnUpdate)
	// gui.Add(light.Rotation().JSValue(), "y", -math.Pi, math.Pi).Name("y rotation").OnChange(jsfnUpdate)
	// gui.Add(light.Rotation().JSValue(), "z", -math.Pi, math.Pi).Name("z rotation").OnChange(jsfnUpdate)
	gui.Add(light.JSValue(), "decay", 0.0, 4.0).OnChange(jsfnUpdate)
	gui.Add(light.JSValue(), "power", 0, 1220).OnChange(jsfnUpdate)

	c.makeXYZGUI(gui, light.Position(), "position", jsfnUpdate)
	// c.makeXYZGUI(gui, light.Target().Position(), "target", jsfnUpdate)

}

func (c *Fundamental5) makeXYZGUI(gui datgui.GUI, v *threejs.Vector3, name string, onChangeFn js.Func) {
	folder := gui.AddFolder(name)
	folder.Add(v.JSValue(), "x", -10, 10).OnChange(onChangeFn)
	folder.Add(v.JSValue(), "y", 0, 10).OnChange(onChangeFn)
	folder.Add(v.JSValue(), "z", -10, 10).OnChange(onChangeFn)
	folder.Open()
}

// resizeRendererToDisplaySize resizes render display size.
func (c *Fundamental5) resizeRendererToDisplaySize(renderer threejs.Renderer) (sizeChanged bool) {
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
func (c *Fundamental5) render(this js.Value, args []js.Value) interface{} {

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
