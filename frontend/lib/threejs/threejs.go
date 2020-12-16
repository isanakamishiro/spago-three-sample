package threejs

import "syscall/js"

var module js.Value

// ThreeJsModule Module is module object for Three.js
// type ThreeJsModule js.Value

// ThreeJs is core type for Three.js
type ThreeJs struct {
	module js.Value
}

// NewThreeJs is constructor for ThreeJs.
func NewThreeJs(module js.Value) *ThreeJs {
	return &ThreeJs{module: module}
}

// NewClock is factory method for Clock.
func (t *ThreeJs) NewClock(autoStart bool) Clock {
	return &ClockImp{Value: t.get("Clock").New(autoStart)}
}

// NewColor is factory method for Color.
func (t *ThreeJs) NewColor(color *Color) *Color {
	return &Color{Value: t.get("Color").New(color)}
}

// NewColor2 is factory method for Color.
func (t *ThreeJs) NewColor2(r float64, g float64, b float64) *Color {
	return &Color{Value: t.get("Color").New(r, g, b)}
}

// NewWebGLRenderer is factory method for WebGLRenderer.
func (t *ThreeJs) NewWebGLRenderer(parameters WebGLRendererParameters) *WebGLRenderer {
	return &WebGLRenderer{Value: t.get("WebGLRenderer").New(parameters)}
}

// NewCamera is factory method for Camera.
func (t *ThreeJs) NewCamera() *CameraImpl {
	return &CameraImpl{Value: t.get("Camera").New()}
}

// NewPerspectiveCamera is factory method for PerspectiveCamera.
func (t *ThreeJs) NewPerspectiveCamera(fov float64, aspect float64, near float64, far float64) *PerspectiveCamera {
	return &PerspectiveCamera{Value: t.get("PerspectiveCamera").New(fov, aspect, near, far)}
}

// NewEuler is factory method for Euler.
func (t *ThreeJs) NewEuler(x float64, y float64, z float64, order string) *Euler {
	return &Euler{Value: t.get("Euler").New(x, y, z, order)}
}

// NewLayers is factory method for Layers.
func (t *ThreeJs) NewLayers() *Layers {
	return &Layers{Value: t.get("Layers").New()}
}

// NewObject3D is factory method for Object3D.
func (t *ThreeJs) NewObject3D() *Object3D {
	return &Object3D{Value: t.get("Object3D").New()}
}

// NewScene is factory method for Scene.
func (t *ThreeJs) NewScene() *SceneImpl {
	return &SceneImpl{Value: t.get("Scene").New()}
}

// NewMaterial is factory method for Material.
func (t *ThreeJs) NewMaterial() *MaterialImpl {
	return &MaterialImpl{Value: t.get("Material").New()}
}

// NewMesh is factory method for MeshImpl.
func (t *ThreeJs) NewMesh(geometry js.Wrapper, material Material) *MeshImpl {
	return &MeshImpl{Value: t.get("Mesh").New(geometry.JSValue(), material.JSValue())}
}

// NewMeshNormalMaterial is factory method for MeshNormalMaterial.
func (t *ThreeJs) NewMeshNormalMaterial(parameters js.Value) *MeshNormalMaterial {
	return &MeshNormalMaterial{Value: t.get("MeshNormalMaterial").New(parameters)}
}

// NewMeshLambertMaterial is factory method for MeshLambertMaterial.
func (t *ThreeJs) NewMeshLambertMaterial(parameters MeshLambertMaterialParameters) *MeshLambertMaterial {
	return &MeshLambertMaterial{Value: t.get("MeshLambertMaterial").New(parameters)}
}

// NewMeshStandardMaterial is factory method for MeshStandardMaterial.
func (t *ThreeJs) NewMeshStandardMaterial(parameters MeshStandardMaterialParameters) *MeshStandardMaterial {
	return &MeshStandardMaterial{Value: t.get("MeshStandardMaterial").New(parameters)}
}

// NewMeshPhongMaterial is ...
func (t *ThreeJs) NewMeshPhongMaterial(parameters MeshPhongMaterialParameters) *MeshPhongMaterial {
	return &MeshPhongMaterial{Value: t.get("MeshPhongMaterial").New(parameters)}
}

// NewGeometry is factory method for MeshNormalMaterial.
func (t *ThreeJs) NewGeometry() *GeometryImpl {
	return &GeometryImpl{Value: t.get("Geometry").New()}
}

// NewBufferGeometry is factory method for BufferGeometry.
func (t *ThreeJs) NewBufferGeometry() *BufferGeometryImpl {
	return &BufferGeometryImpl{Value: t.get("BufferGeometry").New()}
}

// NewBoxBufferGeometry is factory method for BoxBufferGeometry.
func (t *ThreeJs) NewBoxBufferGeometry(width float64, height float64, depth float64, widthSegments int, heightSegments int, depthSegments int) *BoxBufferGeometry {
	return &BoxBufferGeometry{Value: t.get("BoxBufferGeometry").New(width, height, depth, widthSegments, heightSegments, depthSegments)}
}

// NewBoxGeometry is factory method for BoxGeometry.
func (t *ThreeJs) NewBoxGeometry(width float64, height float64, depth float64, widthSegments int, heightSegments int, depthSegments int) *BoxGeometry {
	return &BoxGeometry{Value: t.get("BoxGeometry").New(width, height, depth, widthSegments, heightSegments, depthSegments)}
}

// NewPlaneBufferGeometry is ...
func (t *ThreeJs) NewPlaneBufferGeometry(width float64, height float64, widthSegments int, heightSegments int) *PlaneBufferGeometry {
	return &PlaneBufferGeometry{Value: t.get("PlaneBufferGeometry").New(width, height, widthSegments, heightSegments)}
}

// NewMatrix3 is factory method for Matrix3.
func (t *ThreeJs) NewMatrix3() *Matrix3 {
	return &Matrix3{Value: t.get("Matrix3").New()}
}

// NewMatrix4 is factory method for Matrix4.
func (t *ThreeJs) NewMatrix4() *Matrix4 {
	return &Matrix4{Value: t.get("Matrix4").New()}
}

// NewVector2 is ...
func (t *ThreeJs) NewVector2(x float64, y float64) *Vector2 {
	return &Vector2{Value: t.get("Vector2").New(x, y)}
}

// NewVector3 is ...
func (t *ThreeJs) NewVector3(x float64, y float64, z float64) *Vector3 {
	return &Vector3{Value: t.get("Vector3").New(x, y, z)}
}

// NewVector4 is ...
func (t *ThreeJs) NewVector4(x float64, y float64, z float64, w float64) *Vector4 {
	return &Vector4{Value: t.get("Vector4").New(x, y, z, w)}
}

// NewQuaternion is factory method for Quaternion.
func (t *ThreeJs) NewQuaternion(x float64, y float64, z float64, w float64) *Quaternion {
	return &Quaternion{Value: t.get("Quaternion").New(x, y, z, w)}
}

// NewAmbientLightFromColor is factory method for AmbientLight.
func (t *ThreeJs) NewAmbientLightFromColor(color *Color, intensity float64) *AmbientLight {
	return &AmbientLight{Value: t.get("AmbientLight").New(color, intensity)}
}

// NewAmbientLight is factory method for AmbientLight.
func (t *ThreeJs) NewAmbientLight(colorValue int) *AmbientLight {
	return &AmbientLight{Value: t.get("AmbientLight").New(colorValue)}
}

// NewDirectionalLightFromColor is factory method for DirectionalLight.
func (t *ThreeJs) NewDirectionalLightFromColor(color *Color, intensity float64) *DirectionalLight {
	return &DirectionalLight{Value: t.get("DirectionalLight").New(color, intensity)}
}

// NewDirectionalLight is factory method for DirectionalLight.
func (t *ThreeJs) NewDirectionalLight(colorValue int) *DirectionalLight {
	return &DirectionalLight{Value: t.get("DirectionalLight").New(colorValue)}
}

// NewPolarGridHelper is factory method for PolarGridHelper.
func (t *ThreeJs) NewPolarGridHelper(radius float64, radials float64) *PolarGridHelper {
	return &PolarGridHelper{Value: t.get("PolarGridHelper").New(radius, radials)}
}

// Object3DIdCount is ...
func (t *ThreeJs) Object3DIdCount() int {
	return t.get("Object3DIdCount").Int()
}

// SetObject3DIdCount is ...
func (t *ThreeJs) SetObject3DIdCount(v int) {
	t.set("Object3DIdCount", v)
}

// get is getter for JavaScript object of ThreeJs.
func (t *ThreeJs) get(key string) js.Value {
	return t.module.Get(key)
}

// set is setter for JavaScript object of ThreeJs.
func (t *ThreeJs) set(key string, v interface{}) {
	t.module.Set(key, v)
}
