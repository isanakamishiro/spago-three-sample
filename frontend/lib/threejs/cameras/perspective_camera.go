package cameras

import (
	"app/frontend/lib/threejs"
	"syscall/js"
)

// PerspectiveCamera is Camera that uses perspective projection.
// This projection mode is designed to mimic the way the human eye sees.
// It is the most common projection mode used for rendering a 3D scene.
type PerspectiveCamera interface {
	threejs.Camera
	Aspect() float64
	SetAspect(v float64)
	UpdateProjectionMatrix()
}

// perspectiveCameraImp is implementation of PerspectiveCamera
type perspectiveCameraImp struct {
	threejs.Object3D
	// js.Value
}

// NewPerspectiveCamera is factory method for PerspectiveCamera.
func NewPerspectiveCamera(fov float64, aspect float64, near float64, far float64) PerspectiveCamera {
	return &perspectiveCameraImp{threejs.NewObject3DFromJSValue(threejs.GetJsObject("PerspectiveCamera").New(fov, aspect, near, far))}
}

// NewPerspectiveCameraFromJSValue is ...
func NewPerspectiveCameraFromJSValue(value js.Value) PerspectiveCamera {
	return &perspectiveCameraImp{threejs.NewObject3DFromJSValue(value)}
}

// Aspect is ...
func (pc *perspectiveCameraImp) Aspect() float64 {
	return pc.JSValue().Get("aspect").Float()
}

// SetAspect is ...
func (pc *perspectiveCameraImp) SetAspect(v float64) {
	pc.JSValue().Set("aspect", v)
}

// ProjectionMatrix is ...
func (pc *perspectiveCameraImp) ProjectionMatrix() *threejs.Matrix4 {
	return &threejs.Matrix4{Value: pc.JSValue().Get("projectionMatrix")}
}

// SetProjectionMatrix is ...
func (pc *perspectiveCameraImp) SetProjectionMatrix(v *threejs.Matrix4) {
	pc.JSValue().Set("projectionMatrix", v.JSValue())
}

// UpdateProjectionMatrix is ...
func (pc *perspectiveCameraImp) UpdateProjectionMatrix() {
	pc.JSValue().Call("updateProjectionMatrix")
}

// MatrixWorldInverse is the inverse of matrixWorld.
// MatrixWorld contains the Matrix which has the world transform of the Camera.
func (pc *perspectiveCameraImp) MatrixWorldInverse() *threejs.Matrix4 {
	return &threejs.Matrix4{Value: pc.JSValue().Get("matrixWorldInverse")}
}

// ProjectionMatrixInverse is the inverse of projectionMatrix.
func (pc *perspectiveCameraImp) ProjectionMatrixInverse() *threejs.Matrix4 {
	return &threejs.Matrix4{Value: pc.JSValue().Get("projectionMatrixInverse")}
}

// UpdateMatrixWorld is ...
func (pc *perspectiveCameraImp) UpdateMatrixWorld(force bool) {
	pc.JSValue().Call("updateMatrixWorld", force)
}

// IsCamera is ...
func (pc *perspectiveCameraImp) IsCamera() bool {
	return pc.JSValue().Get("isCamera").Bool()
}

// WorldDirection is ...
func (pc *perspectiveCameraImp) WorldDirection(target *threejs.Vector3) *threejs.Vector3 {
	return &threejs.Vector3{Value: pc.JSValue().JSValue().Call("getWorldDirection", target)}
}
