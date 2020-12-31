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

//go:generate spago generate -c Fundamental8 -p views threejs_fundamental8.html

// Fundamental8  ...
type Fundamental8 struct {
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

	renderRequested bool
}

// NewFundamental8 is ...
func NewFundamental8() *Fundamental8 {

	return &Fundamental8{
		callbacks: threejs.NewCallbackRepository(),
	}
}

// Mount is ...
func (c *Fundamental8) Mount() {

	// shortcut of js.Funcs
	c.renderFunc = c.callbacks.Register(c.render)

	// initialize objects
	c.initSceneAndRenderer()

	// first render
	threejs.RequestAnimationFrame(c.renderFunc)

}

// Unmount ...
func (c *Fundamental8) Unmount() {

	// Release all js.Funcs
	c.callbacks.ReleaseAll()

}

func (c *Fundamental8) initSceneAndRenderer() {

	// Renderer
	renderer := threejs.NewWebGLRenderer(map[string]interface{}{
		"canvas":    js.Global().Get("document").Call("querySelector", "#c"),
		"antialias": true,
		// "alpha":     true,
	})
	renderer.ShadowMap().SetEnabled(true)
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
		fov    = 45
		aspect = 2
		near   = 0.1
		far    = 100
	)
	camera := cameras.NewPerspectiveCamera(fov, aspect, near, far)
	camera.Position().Set2(0, 10, 20)
	c.camera = camera

	// Controls
	control := controls.NewOrbitControls(c.camera, c.renderer.DomElement())
	control.Target().Set2(0, 5, 0)
	control.Update()
	c.control = control

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
	planeMat := materials.NewMeshPhongMaterial(map[string]interface{}{
		// planeMat := materials.NewMeshStandardMaterial(map[string]interface{}{
		"map":  texture,
		"side": threejs.DoubleSide.JSValue(),
	})
	planeMesh := threejs.NewMesh(planeGeo, planeMat)
	planeMesh.Rotation().SetX(math.Pi * -0.5)
	planeMesh.SetReceiveShadow(true)
	scene.Add(planeMesh)

	// Objects : Cube
	const cubeSize = 4.0
	cubeGeo := geometries.NewBoxBufferGeometry(cubeSize, cubeSize, cubeSize, 1, 1, 1)
	cubeMat := materials.NewMeshPhongMaterial(map[string]interface{}{
		// cubeMat := materials.NewMeshStandardMaterial(map[string]interface{}{
		"color": "#8AC",
	})
	cubeMesh := threejs.NewMesh(cubeGeo, cubeMat)
	cubeMesh.Position().Set2(cubeSize+1, cubeSize/2, 0)
	cubeMesh.SetCastShadow(true)
	cubeMesh.SetReceiveShadow(true)
	scene.Add(cubeMesh)

	// Objects : Sphere
	const (
		sphereRadius          = 3.0
		sphereWidthDivisions  = 32
		sphereHeightDivisions = 16
	)
	sphereGeo := geometries.NewSphereBufferGeometry(sphereRadius, sphereWidthDivisions, sphereHeightDivisions)
	sphereMat := materials.NewMeshPhongMaterial(map[string]interface{}{
		// sphereMat := materials.NewMeshStandardMaterial(map[string]interface{}{
		"color": "#CA8",
	})
	sphereMesh := threejs.NewMesh(sphereGeo, sphereMat)
	sphereMesh.Position().Set2(-sphereRadius-1, sphereRadius+2, 0)
	sphereMesh.SetCastShadow(true)
	sphereMesh.SetReceiveShadow(true)
	scene.Add(sphereMesh)

	// Lights

	// Directional Light
	const (
		lightColor     = threejs.ColorValue(0xffffff)
		lightIntensity = threejs.LightIntensity(1)
		width          = 12.0
		height         = 4.0
	)
	light := lights.NewDirectionalLight(lightColor, lightIntensity)
	light.SetCastShadow(true)
	light.Position().Set2(0, 10, 0)
	light.Target().Position().Set2(-4, 0, -4)
	scene.AddLight(light)
	scene.Add(light.Target())

	cameraHelper := cameras.NewCameraHelper(light.Shadow().Camera())
	scene.Add(cameraHelper)

	helper := lights.NewDirectionalLightHelper(light)
	scene.Add(helper)

	// Objects
	c.objects = make([]threejs.Object3D, 0, 1)

	// build GUI
	jsfnUpdate := c.callbacks.Register(func(this js.Value, args []js.Value) interface{} {
		light.Target().UpdateMatrixWorld(true)
		helper.Update()

		return nil
	})
	jsfnUpdate.Invoke()

	params := map[string]interface{}{
		"color": int(lightColor),
		"width": light.Shadow().Camera().Right() * 2,
	}
	gui.AddColor(params, "color").Name("color").OnChange(c.callbacks.Register(
		func(this js.Value, args []js.Value) interface{} {
			col := args[0].Int()

			light.SetColor(threejs.NewColorFromColorValue(threejs.ColorValue(col)))
			return nil
		},
	))
	gui.Add(light.JSValue(), "intensity", 0, 10)
	folder := gui.AddFolder("Shadow Camera")
	folder.Open()
	folder.Add(js.ValueOf(params), "width", 1, 100).Name("width")
	// gui.Add(light.JSValue(), "decay", 0.0, 4.0).OnChange(jsfnUpdate)
	// gui.Add(light.JSValue(), "power", 0, 1220).OnChange(jsfnUpdate)

	c.makeXYZGUI(gui, light.Position(), "position", jsfnUpdate)
	c.makeXYZGUI(gui, light.Target().Position(), "target", jsfnUpdate)

}

func (c *Fundamental8) makeXYZGUI(gui datgui.GUI, v *threejs.Vector3, name string, onChangeFn js.Func) {
	folder := gui.AddFolder(name)
	folder.Add(v.JSValue(), "x", -10, 10).OnChange(onChangeFn)
	folder.Add(v.JSValue(), "y", 0, 10).OnChange(onChangeFn)
	folder.Add(v.JSValue(), "z", -10, 10).OnChange(onChangeFn)
	folder.Open()
}

// resizeRendererToDisplaySize resizes render display size.
func (c *Fundamental8) resizeRendererToDisplaySize(renderer threejs.Renderer) (sizeChanged bool) {
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
func (c *Fundamental8) render(this js.Value, args []js.Value) interface{} {

	if sizeChanged := c.resizeRendererToDisplaySize(c.renderer); sizeChanged {
		canvas := c.renderer.DomElement()
		clientWidth := canvas.Get("clientWidth").Float()
		clientHeight := canvas.Get("clientHeight").Float()
		c.camera.(cameras.PerspectiveCamera).SetAspect(clientWidth / clientHeight)
		c.camera.(cameras.PerspectiveCamera).UpdateProjectionMatrix()
	}

	// time := args[0].Float() * 0.001
	// for i, obj := range c.objects {
	// 	speed := 0.2 + float64(i)*0.1
	// 	rot := time * speed

	// 	obj.Rotation().SetX(rot)
	// 	obj.Rotation().SetY(rot)
	// }

	c.control.Update()
	c.renderer.Render(c.scene, c.camera)

	c.stats.Update()

	threejs.RequestAnimationFrame(c.renderFunc)

	return nil
}
