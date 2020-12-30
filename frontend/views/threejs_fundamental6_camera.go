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

//go:generate spago generate -c Fundamental6 -p views threejs_fundamental6.html

// Fundamental6  ...
type Fundamental6 struct {
	spago.Core

	canvas js.Value
	view1  js.Value
	view2  js.Value

	camera  threejs.Camera
	camera2 threejs.Camera

	cameraHelper *cameras.CameraHelper

	scene    threejs.Scene
	renderer threejs.Renderer
	objects  []threejs.Object3D

	control  controls.OrbitControls
	control2 controls.OrbitControls

	gui   datgui.GUI
	stats stats.Stats

	callbacks threejs.CallbackRepository

	renderID   threejs.RenderID
	renderFunc js.Func

	renderRequested bool
}

// NewFundamental6 is ...
func NewFundamental6() *Fundamental6 {

	return &Fundamental6{
		callbacks: threejs.NewCallbackRepository(),
	}
}

// Mount is ...
func (c *Fundamental6) Mount() {

	// shortcut of js.Funcs
	c.renderFunc = c.callbacks.Register(c.render)

	// initialize objects
	c.initSceneAndRenderer()

	// first render
	threejs.RequestAnimationFrame(c.renderFunc)

}

// Unmount ...
func (c *Fundamental6) Unmount() {

	// Release all js.Funcs
	c.callbacks.ReleaseAll()

}

func (c *Fundamental6) initSceneAndRenderer() {

	// Renderer
	canvas := js.Global().Get("document").Call("querySelector", "#c")
	renderer := threejs.NewWebGLRenderer(map[string]interface{}{
		"canvas":    canvas,
		"antialias": true,
		// "logarithmicDepthBuffer": true,
		// "alpha":     true,
	})
	c.canvas = canvas
	c.renderer = renderer

	// GUI
	gui := datgui.NewGUI()
	c.gui = gui

	// Stats
	stats := stats.NewStats()
	js.Global().Get("document").Get("body").Call("appendChild", stats.DomElement())
	c.stats = stats

	// Scene
	scene := threejs.NewScene()
	c.scene = scene

	// Camera1 / Control1
	const (
		fov    = 45
		aspect = 2
		near   = 5
		far    = 100

		left   = -1
		right  = 1
		top    = 1
		bottom = -1
	)
	// camera := cameras.NewPerspectiveCamera(fov, aspect, near, far)
	camera := cameras.NewOrthographicCamera(left, right, top, bottom, near, far)
	camera.SetZoom(0.2)
	camera.Position().Set2(0, 10, 20)
	c.camera = camera

	cameraHelper := cameras.NewCameraHelper(camera)
	scene.Add(cameraHelper)
	c.cameraHelper = cameraHelper

	// Controls
	view1Elem := js.Global().Get("document").Call("querySelector", "#view1")
	c.view1 = view1Elem

	// controls := controls.NewOrbitControls(c.camera, c.renderer.DomElement())
	control := controls.NewOrbitControls(camera, view1Elem)
	control.Target().Set2(0, 5, 0)
	control.Update()
	c.control = control

	// Camera2 / Control2
	camera2 := cameras.NewPerspectiveCamera(
		60,
		2,
		0.1,
		500,
	)
	camera2.Position().Set2(40, 10, 30)
	camera2.LookAtXYZ(0, 5, 0)
	c.camera2 = camera2

	// Control2
	view2Elem := js.Global().Get("document").Call("querySelector", "#view2")
	c.view2 = view2Elem

	control2 := controls.NewOrbitControls(camera2, view2Elem)
	control2.Target().Set2(0, 5, 0)
	control2.Update()
	c.control2 = control2

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
		"map":  texture,
		"side": threejs.DoubleSide.JSValue(),
	})
	planeMesh := threejs.NewMesh(planeGeo, planeMat)
	planeMesh.Rotation().SetX(math.Pi * -0.5)
	scene.Add(planeMesh)

	// Objects : Cube
	const cubeSize = 4.0
	cubeGeo := geometries.NewBoxBufferGeometry(cubeSize, cubeSize, cubeSize, 1, 1, 1)
	cubeMat := materials.NewMeshPhongMaterial(map[string]interface{}{
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
		numSpheres            = 20
	)
	sphereGeo := geometries.NewSphereBufferGeometry(sphereRadius, sphereWidthDivisions, sphereHeightDivisions)
	sphereMat := materials.NewMeshPhongMaterial(map[string]interface{}{
		"color": "#CA8",
	})
	sphereMesh := threejs.NewMesh(sphereGeo, sphereMat)
	sphereMesh.Position().Set2(-sphereRadius-1, sphereRadius+2, 0)
	scene.Add(sphereMesh)

	// for i := 0; i < numSpheres; i++ {
	// 	sphereMat := materials.NewMeshPhongMaterial(map[string]interface{}{
	// 		"color": "#CA8",
	// 	})
	// 	sphereMesh := threejs.NewMesh(sphereGeo, sphereMat)
	// 	sphereMesh.Position().Set2(-sphereRadius-1, sphereRadius+2, float64(i)*sphereRadius*-2.2)
	// 	scene.Add(sphereMesh)
	// }

	// Lights
	const (
		lightColor     = threejs.ColorValue(0xffffff)
		lightIntensity = threejs.LightIntensity(1)
	)
	light := lights.NewDirectionalLight(lightColor, lightIntensity)
	light.Position().Set2(0, 10, 0)
	light.Target().Position().Set2(-5, 0, 0)
	scene.AddLight(light)
	scene.Add(light.Target())

	helper := lights.NewDirectionalLightHelper(light)
	scene.Add(helper)

	// Objects
	c.objects = make([]threejs.Object3D, 0, 1)

	// build GUI
	// gui.Add(camera.JSValue(), "fov", 0, 180)
	// gui.Add(camera.JSValue(), "near", 0.000001, 50)
	// gui.Add(camera.JSValue(), "far", 0.1, 50)
	gui.Add(camera.JSValue(), "zoom", 0.01, 1).Listen()

}

// setScissorForElement computes scissor and sets scissor and viewport for renderer.
func (c *Fundamental6) setScissorForElement(e js.Value) float64 {
	pixelRatio := js.Global().Get("devicePixelRatio").Float()

	canvasRect := threejs.NewRectangleFromJSValue(c.canvas.Call("getBoundingClientRect"))
	elemRect := threejs.NewRectangleFromJSValue(e.Call("getBoundingClientRect"))

	// compute a canvas relative rectangle
	right := math.Min(elemRect.Right(), canvasRect.Right()) - canvasRect.Left()
	left := math.Max(0, elemRect.Left()-canvasRect.Left())
	bottom := math.Min(elemRect.Bottom(), canvasRect.Bottom()) - canvasRect.Top()
	top := math.Max(0, elemRect.Top()-canvasRect.Top())

	width := math.Min(canvasRect.Width(), right-left)
	if width <= 0 {
		width = 1
	}
	height := math.Min(canvasRect.Height(), bottom-top)
	if height <= 0 {
		height = 1
	}

	// setup the scissor to only render to that part of the canvas
	positiveYUpBottom := canvasRect.Height() - bottom

	c.renderer.SetScissor(
		int(left*pixelRatio),
		int(positiveYUpBottom*pixelRatio),
		int(width*pixelRatio),
		int(height*pixelRatio),
	)
	c.renderer.SetViewport(
		int(left*pixelRatio),
		int(positiveYUpBottom*pixelRatio),
		int(width*pixelRatio),
		int(height*pixelRatio),
	)

	// return the aspect
	return width / height
}

// resizeRendererToDisplaySize resizes render display size.
func (c *Fundamental6) resizeRendererToDisplaySize(renderer threejs.Renderer) (sizeChanged bool) {
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
func (c *Fundamental6) render(this js.Value, args []js.Value) interface{} {

	if sizeChanged := c.resizeRendererToDisplaySize(c.renderer); sizeChanged {
		// canvas := c.renderer.DomElement()
		// clientWidth := canvas.Get("clientWidth").Float()
		// clientHeight := canvas.Get("clientHeight").Float()
		// c.camera.(cameras.PerspectiveCamera).SetAspect(clientWidth / clientHeight)
		// c.camera.(cameras.PerspectiveCamera).UpdateProjectionMatrix()
	}

	// turn on the scissor
	c.renderer.SetScissorTest(true)
	// c.renderer.SetScissorTest(false)

	// render the original view
	{
		aspect := c.setScissorForElement(c.view1)
		// c.camera.(cameras.PerspectiveCamera).SetAspect(aspect)
		// c.camera.(cameras.PerspectiveCamera).UpdateProjectionMatrix()
		c.camera.(cameras.OrthographicCamera).SetLeft(-aspect)
		c.camera.(cameras.OrthographicCamera).SetRight(aspect)
		c.camera.(cameras.OrthographicCamera).UpdateProjectionMatrix()

		c.cameraHelper.Update()

		// don't draw the camera helper in the original view
		c.cameraHelper.SetVisible(false)

		c.scene.SetBackgroundColor(threejs.NewColorFromColorValue(threejs.ColorValue(0x000000)))

		c.renderer.Render(c.scene, c.camera)

	}

	// render from the 2nd camera
	{
		aspect := c.setScissorForElement(c.view2)

		// adjust the camera for this aspect
		c.camera2.(cameras.PerspectiveCamera).SetAspect(aspect)
		c.camera2.(cameras.PerspectiveCamera).UpdateProjectionMatrix()

		// don't draw the camera helper in the original view
		c.cameraHelper.SetVisible(true)

		c.scene.SetBackgroundColor(threejs.NewColorFromColorValue(threejs.ColorValue(0x000040)))

		c.renderer.Render(c.scene, c.camera2)

	}

	c.stats.Update()
	threejs.RequestAnimationFrame(c.renderFunc)

	return nil
}
