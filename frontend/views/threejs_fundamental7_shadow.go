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
	"app/frontend/lib/threejs/mathutils"
	"app/frontend/lib/threejs/stats"

	"math"

	"syscall/js"

	"github.com/nobonobo/spago"
)

//go:generate spago generate -c Fundamental7 -p views threejs_fundamental7.html

type shadowObject struct {
	base       threejs.Object3D
	sphereMesh threejs.Mesh
	shadowMesh threejs.Mesh
	y          float64
}

// Fundamental7  ...
type Fundamental7 struct {
	spago.Core

	canvas js.Value

	camera threejs.Camera
	// cameraHelper *cameras.CameraHelper

	scene             threejs.Scene
	renderer          threejs.Renderer
	sphereShadowBases []shadowObject

	control controls.OrbitControls

	gui   datgui.GUI
	stats stats.Stats

	callbacks threejs.CallbackRepository

	renderID   threejs.RenderID
	renderFunc js.Func

	renderRequested bool
}

// NewFundamental7 is ...
func NewFundamental7() *Fundamental7 {

	return &Fundamental7{
		callbacks: threejs.NewCallbackRepository(),
	}
}

// Mount is ...
func (c *Fundamental7) Mount() {

	// shortcut of js.Funcs
	c.renderFunc = c.callbacks.Register(c.render)

	// initialize objects
	c.initSceneAndRenderer()

	// first render
	threejs.RequestAnimationFrame(c.renderFunc)

}

// Unmount ...
func (c *Fundamental7) Unmount() {

	// Release all js.Funcs
	c.callbacks.ReleaseAll()

}

func (c *Fundamental7) initSceneAndRenderer() {

	// Renderer
	canvas := js.Global().Get("document").Call("querySelector", "#c")
	renderer := threejs.NewWebGLRenderer(map[string]interface{}{
		"canvas":    canvas,
		"antialias": true,
	})
	renderer.SetPhysicallyCorrectLights(true)
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
	scene.SetBackgroundColor(threejs.NewColorFromColorName("white"))
	c.scene = scene

	// Camera1 / Control1
	const (
		fov    = 45
		aspect = 2
		near   = 0.1
		far    = 100
	)
	camera := cameras.NewPerspectiveCamera(fov, aspect, near, far)
	camera.Position().Set2(0, 10, 20)
	camera.LookAtXYZ(0, 0, 0)
	c.camera = camera

	// Controls
	control := controls.NewOrbitControls(c.camera, c.renderer.DomElement())
	control.Target().Set2(0, 5, 0)
	control.Update()
	c.control = control

	// Texture
	loader := loaders.NewTextureLoader()

	const planeSize = 40

	texture := loader.LoadSimply("https://threejsfundamentals.org/threejs/resources/images/checker.png")
	texture.SetWrapS(threejs.RepeatWrapping)
	texture.SetWrapT(threejs.RepeatWrapping)
	texture.SetMagFilter(threejs.NearestFilter)
	repeats := planeSize / 2.0
	texture.Repeat().Set2(repeats, repeats)

	// Objects : Plane
	planeGeo := geometries.NewPlaneBufferGeometry(planeSize, planeSize, 1, 1)
	planeMat := materials.NewMeshBasicMaterial(map[string]interface{}{
		"map":  texture,
		"side": threejs.DoubleSide.JSValue(),
	})
	planeMat.Color().SetRGB(1.5, 1.5, 1.5)
	planeMesh := threejs.NewMesh(planeGeo, planeMat)
	planeMesh.Rotation().SetX(math.Pi * -0.5)
	scene.Add(planeMesh)

	shadowTexture := loader.LoadSimply("https://threejsfundamentals.org/threejs/resources/images/roundshadow.png")

	// Objects : Sphere
	const (
		sphereRadius          = 1.0
		sphereWidthDivisions  = 32
		sphereHeightDivisions = 16
		numSpheres            = 15
	)
	sphereGeo := geometries.NewSphereBufferGeometry(sphereRadius, sphereWidthDivisions, sphereHeightDivisions)
	shadowGeo := geometries.NewPlaneBufferGeometry(1, 1, 1, 1)

	// Objects
	c.sphereShadowBases = make([]shadowObject, 0)

	for i := 0; i < numSpheres; i++ {
		base := threejs.NewObject3D()
		scene.Add(base)

		// add the shadow to the base
		// note: we make a new material for each sphere
		// so we can set that sphere's material transparency
		// separately.
		shadowMat := materials.NewMeshBasicMaterial(map[string]interface{}{
			"map":         shadowTexture,
			"transparent": true,  // so we can see the ground
			"depthWrite":  false, // so we don't have to sort
		})
		shadowMesh := threejs.NewMesh(shadowGeo, shadowMat)
		shadowMesh.Position().SetY(0.001)
		shadowMesh.Rotation().SetX(math.Pi * -0.5)
		shadowSize := sphereRadius * 4.0
		shadowMesh.Scale().Set2(shadowSize, shadowSize, shadowSize)
		base.Add(shadowMesh)

		// add the sphere to the base
		u := float64(i) / numSpheres
		sphereMat := materials.NewMeshPhongMaterial(map[string]interface{}{})
		sphereMat.Color().SetHSL(u, 1, 0.75)
		sphereMesh := threejs.NewMesh(sphereGeo, sphereMat)
		sphereMesh.Position().Set2(0, sphereRadius+2, 0)
		base.Add(sphereMesh)

		c.sphereShadowBases = append(c.sphereShadowBases, shadowObject{
			base:       base,
			sphereMesh: sphereMesh,
			shadowMesh: shadowMesh,
			y:          sphereMesh.Position().Y(),
		})

	}

	// Lights
	{
		const (
			skyColor       = threejs.ColorValue(0xb1e1ff)
			groundColor    = threejs.ColorValue(0xb97a20)
			lightIntensity = threejs.LightIntensity(2)
		)
		light := lights.NewHemisphereLight(skyColor, groundColor, lightIntensity)
		scene.AddLight(light)
	}
	{
		const (
			lightColor     = threejs.ColorValue(0xffffff)
			lightIntensity = threejs.LightIntensity(1)
		)
		light := lights.NewDirectionalLight(lightColor, lightIntensity)
		light.Position().Set2(0, 10, 5)
		light.Target().Position().Set2(-5, 0, 0)
		scene.AddLight(light)
		scene.Add(light.Target())
		helper := lights.NewDirectionalLightHelper(light)
		scene.Add(helper)
	}

	// build GUI
	// gui.Add(camera.JSValue(), "fov", 0, 180)

}

// resizeRendererToDisplaySize resizes render display size.
func (c *Fundamental7) resizeRendererToDisplaySize(renderer threejs.Renderer) (sizeChanged bool) {
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
func (c *Fundamental7) render(this js.Value, args []js.Value) interface{} {

	if sizeChanged := c.resizeRendererToDisplaySize(c.renderer); sizeChanged {
		canvas := c.renderer.DomElement()
		clientWidth := canvas.Get("clientWidth").Float()
		clientHeight := canvas.Get("clientHeight").Float()
		c.camera.(cameras.PerspectiveCamera).SetAspect(clientWidth / clientHeight)
		c.camera.(cameras.PerspectiveCamera).UpdateProjectionMatrix()
	}

	time := args[0].Float() * 0.001
	for i, obj := range c.sphereShadowBases {

		// u is a value that goes from 0 to 1 as we iterate the spheres
		u := float64(i) / float64(len(c.sphereShadowBases))

		// compute a position for there base. This will move
		// both the sphere and its shadow
		speed := time * 0.2
		angle := speed + u*math.Pi*2 // * (i % 1 ? 1 : -1)
		radius := math.Sin(speed-float64(i)) * 10
		obj.base.Position().Set2(math.Cos(angle)*radius, 0, math.Sin(angle)*radius)

		// yOff is a value that goes from 0 to 1
		yOff := math.Abs(math.Sin(time*2 + float64(i)))
		obj.sphereMesh.Position().SetY(obj.y + mathutils.Lerp(-2, 3, yOff))
		obj.shadowMesh.Material().SetOpacity(mathutils.Lerp(1, 0.25, yOff))

	}

	c.renderer.Render(c.scene, c.camera)

	c.stats.Update()
	threejs.RequestAnimationFrame(c.renderFunc)

	return nil
}
