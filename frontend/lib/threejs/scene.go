package threejs

import (
	"syscall/js"
)

// Scene is ...
type Scene interface {
	Object3D

	AutoUpdate() bool
	SetAutoUpdate(v bool)
	Background() js.Value
	SetBackground(v js.Value)

	Add(v Object3D)
	AddLight(v Light)
	AddMesh(v Mesh)
}

// sceneImpl extend: [Object3D]
type sceneImpl struct {
	Object3D
}

// NewScene is factory method for Scene.
func NewScene() Scene {

	return &sceneImpl{
		NewDefaultObject3D(
			GetJsObject("Scene").New(),
		),
	}
}

// AutoUpdate is ...
func (ss *sceneImpl) AutoUpdate() bool {
	return ss.JSValue().Get("autoUpdate").Bool()
}

// SetAutoUpdate is ...
func (ss *sceneImpl) SetAutoUpdate(v bool) {
	ss.JSValue().Set("autoUpdate", v)
}

// Background is ...
func (ss *sceneImpl) Background() js.Value {
	return ss.JSValue().Get("background")
}

// SetBackground is ...
func (ss *sceneImpl) SetBackground(v js.Value) {
	ss.JSValue().Set("background", v)
}

func (ss *sceneImpl) Add(v Object3D) {
	ss.JSValue().Call("add", v.JSValue())
}

func (ss *sceneImpl) AddLight(v Light) {
	ss.JSValue().Call("add", v.JSValue())
}

func (ss *sceneImpl) AddMesh(v Mesh) {
	ss.JSValue().Call("add", v.JSValue())
}
