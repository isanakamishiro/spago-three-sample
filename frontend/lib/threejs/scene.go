package threejs

import (
	"syscall/js"
)

// Scene is ...
type Scene interface {
	JSValue() js.Value
	AutoUpdate() bool
	SetAutoUpdate(v bool)
	Background() js.Value
	SetBackground(v js.Value)
	CastShadow() bool
	SetCastShadow(v bool)
	Children() js.Value
	SetChildren(v js.Value)
	Fog() js.Value
	SetFog(v js.Value)
	FrustumCulled() bool
	SetFrustumCulled(v bool)
	ID() int
	SetID(v int)
	IsObject3D() bool
	SetIsObject3D(v bool)
	Layers() *Layers
	SetLayers(v *Layers)
	Matrix() *Matrix4
	SetMatrix(v *Matrix4)
	MatrixAutoUpdate() bool
	SetMatrixAutoUpdate(v bool)
	MatrixWorld() *Matrix4
	SetMatrixWorld(v *Matrix4)
	MatrixWorldNeedsUpdate() bool
	SetMatrixWorldNeedsUpdate(v bool)
	ModelViewMatrix() *Matrix4
	SetModelViewMatrix(v *Matrix4)
	Name() string
	SetName(v string)
	NormalMatrix() *Matrix3
	SetNormalMatrix(v *Matrix3)
	OnAfterRender() js.Value
	SetOnAfterRender(v js.Value)
	OnBeforeRender() js.Value
	SetOnBeforeRender(v js.Value)
	OverrideMaterial() Material
	SetOverrideMaterial(v Material)
	Parent() *Object3D
	SetParent(v *Object3D)
	Position() *Vector3
	SetPosition(v *Vector3)
	Quaternion() *Quaternion
	SetQuaternion(v *Quaternion)
	ReceiveShadow() bool
	SetReceiveShadow(v bool)
	RenderOrder() float64
	SetRenderOrder(v float64)
	Rotation() *Euler
	SetRotation(v *Euler)
	Scale() *Vector3
	SetScale(v *Vector3)
	Type() string
	SetType(v string)
	Up() *Vector3
	SetUp(v *Vector3)
	UserData() js.Value
	SetUserData(v js.Value)
	UUID() string
	SetUUID(v string)
	Visible() bool
	SetVisible(v bool)
	DefaultMatrixAutoUpdate() bool
	SetDefaultMatrixAutoUpdate(v bool)
	DefaultUp() *Vector3
	SetDefaultUp(v *Vector3)
	Add(object js.Value) Scene
	AddEventListener(typ string, listener js.Value)
	ApplyMatrix(matrix *Matrix4)
	ApplyQuaternion(quaternion *Quaternion) Scene
	Clone(recursive bool) Scene
	Copy(source Scene, recursive bool) Scene
	DispatchEvent(event js.Value)
	GetObjectByID(id int) *Object3D
	GetObjectByName(name string) *Object3D
	GetObjectByProperty(name string, value string) *Object3D
	GetWorldDirection(target *Vector3) *Vector3
	GetWorldPosition(target *Vector3) *Vector3
	GetWorldQuaternion(target *Quaternion) *Quaternion
	GetWorldScale(target *Vector3) *Vector3
	HasEventListener(typ string, listener js.Value) bool
	LocalToWorld(vector *Vector3) *Vector3
	LookAt(vector *Vector3, y float64, z float64)
	// Raycast(raycaster *Raycaster, intersects js.Value)
	Remove(object js.Value) Scene
	RemoveEventListener(typ string, listener js.Value)
	RotateOnAxis(axis *Vector3, angle float64) Scene
	RotateOnWorldAxis(axis *Vector3, angle float64) Scene
	RotateX(angle float64) Scene
	RotateY(angle float64) Scene
	RotateZ(angle float64) Scene
	SetRotationFromAxisAngle(axis *Vector3, angle float64)
	SetRotationFromEuler(euler *Euler)
	SetRotationFromMatrix(m *Matrix4)
	SetRotationFromQuaternion(q *Quaternion)
	ToJSON(meta js.Value) js.Value
	TranslateOnAxis(axis *Vector3, distance float64) Scene
	TranslateX(distance float64) Scene
	TranslateY(distance float64) Scene
	TranslateZ(distance float64) Scene
	Traverse(callback js.Value)
	TraverseAncestors(callback js.Value)
	TraverseVisible(callback js.Value)
	UpdateMatrix()
	UpdateMatrixWorld(force bool)
	UpdateWorldMatrix(updateParents bool, updateChildren bool)
	WorldToLocal(vector *Vector3) *Vector3
}

// SceneImpl extend: [Object3D]
type SceneImpl struct {
	js.Value
}

// JSValue is ...
func (ss *SceneImpl) JSValue() js.Value {
	return ss.Value
}

// AutoUpdate is ...
func (ss *SceneImpl) AutoUpdate() bool {
	return ss.Get("autoUpdate").Bool()
}

// SetAutoUpdate is ...
func (ss *SceneImpl) SetAutoUpdate(v bool) {
	ss.Set("autoUpdate", v)
}

// Background is ...
func (ss *SceneImpl) Background() js.Value {
	return ss.Get("background")
}

// SetBackground is ...
func (ss *SceneImpl) SetBackground(v js.Value) {
	ss.Set("background", v)
}

// CastShadow is ...
func (ss *SceneImpl) CastShadow() bool {
	return ss.Get("castShadow").Bool()
}

// SetCastShadow is ...
func (ss *SceneImpl) SetCastShadow(v bool) {
	ss.Set("castShadow", v)
}

// Children is ...
func (ss *SceneImpl) Children() js.Value {
	return ss.Get("children")
}

// SetChildren is ...
func (ss *SceneImpl) SetChildren(v js.Value) {
	ss.Set("children", v)
}

// Fog is ...
func (ss *SceneImpl) Fog() js.Value {
	return ss.Get("fog")
}

// SetFog is ...
func (ss *SceneImpl) SetFog(v js.Value) {
	ss.Set("fog", v)
}

// FrustumCulled is ...
func (ss *SceneImpl) FrustumCulled() bool {
	return ss.Get("frustumCulled").Bool()
}

// SetFrustumCulled is ...
func (ss *SceneImpl) SetFrustumCulled(v bool) {
	ss.Set("frustumCulled", v)
}

// ID is ...
func (ss *SceneImpl) ID() int {
	return ss.Get("id").Int()
}

// SetID is ...
func (ss *SceneImpl) SetID(v int) {
	ss.Set("id", v)
}

// IsObject3D is ...
func (ss *SceneImpl) IsObject3D() bool {
	return ss.Get("isObject3D").Bool()
}

// SetIsObject3D is ...
func (ss *SceneImpl) SetIsObject3D(v bool) {
	ss.Set("isObject3D", v)
}

// Layers is ...
func (ss *SceneImpl) Layers() *Layers {
	return &Layers{Value: ss.Get("layers")}
}

// SetLayers is ...
func (ss *SceneImpl) SetLayers(v *Layers) {
	ss.Set("layers", v.JSValue())
}

// Matrix is ...
func (ss *SceneImpl) Matrix() *Matrix4 {
	return &Matrix4{Value: ss.Get("matrix")}
}

// SetMatrix is ...
func (ss *SceneImpl) SetMatrix(v *Matrix4) {
	ss.Set("matrix", v.JSValue())
}

// MatrixAutoUpdate is ...
func (ss *SceneImpl) MatrixAutoUpdate() bool {
	return ss.Get("matrixAutoUpdate").Bool()
}

// SetMatrixAutoUpdate is ...
func (ss *SceneImpl) SetMatrixAutoUpdate(v bool) {
	ss.Set("matrixAutoUpdate", v)
}

// MatrixWorld is ...
func (ss *SceneImpl) MatrixWorld() *Matrix4 {
	return &Matrix4{Value: ss.Get("matrixWorld")}
}

// SetMatrixWorld is ...
func (ss *SceneImpl) SetMatrixWorld(v *Matrix4) {
	ss.Set("matrixWorld", v.JSValue())
}

// MatrixWorldNeedsUpdate is ...
func (ss *SceneImpl) MatrixWorldNeedsUpdate() bool {
	return ss.Get("matrixWorldNeedsUpdate").Bool()
}

// SetMatrixWorldNeedsUpdate is ...
func (ss *SceneImpl) SetMatrixWorldNeedsUpdate(v bool) {
	ss.Set("matrixWorldNeedsUpdate", v)
}

// ModelViewMatrix is ...
func (ss *SceneImpl) ModelViewMatrix() *Matrix4 {
	return &Matrix4{Value: ss.Get("modelViewMatrix")}
}

// SetModelViewMatrix is ...
func (ss *SceneImpl) SetModelViewMatrix(v *Matrix4) {
	ss.Set("modelViewMatrix", v.JSValue())
}

// Name is ...
func (ss *SceneImpl) Name() string {
	return ss.Get("name").String()
}

// SetName is ...
func (ss *SceneImpl) SetName(v string) {
	ss.Set("name", v)
}

// NormalMatrix is ...
func (ss *SceneImpl) NormalMatrix() *Matrix3 {
	return &Matrix3{Value: ss.Get("normalMatrix")}
}

// SetNormalMatrix is ...
func (ss *SceneImpl) SetNormalMatrix(v *Matrix3) {
	ss.Set("normalMatrix", v.JSValue())
}

// OnAfterRender is ...
func (ss *SceneImpl) OnAfterRender() js.Value {
	return ss.Get("onAfterRender")
}

// SetOnAfterRender is ...
func (ss *SceneImpl) SetOnAfterRender(v js.Value) {
	ss.Set("onAfterRender", v)
}

// OnBeforeRender is ...
func (ss *SceneImpl) OnBeforeRender() js.Value {
	return ss.Get("onBeforeRender")
}

// SetOnBeforeRender is ...
func (ss *SceneImpl) SetOnBeforeRender(v js.Value) {
	ss.Set("onBeforeRender", v)
}

// OverrideMaterial is ...
func (ss *SceneImpl) OverrideMaterial() Material {
	return &MaterialImpl{Value: ss.Get("overrideMaterial")}
}

// SetOverrideMaterial is ...
func (ss *SceneImpl) SetOverrideMaterial(v Material) {
	ss.Set("overrideMaterial", v.JSValue())
}

// Parent is ...
func (ss *SceneImpl) Parent() *Object3D {
	return &Object3D{Value: ss.Get("parent")}
}

// SetParent is ...
func (ss *SceneImpl) SetParent(v *Object3D) {
	ss.Set("parent", v.JSValue())
}

// Position is ...
func (ss *SceneImpl) Position() *Vector3 {
	return &Vector3{Value: ss.Get("position")}
}

// SetPosition is ...
func (ss *SceneImpl) SetPosition(v *Vector3) {
	ss.Set("position", v.JSValue())
}

// Quaternion is ...
func (ss *SceneImpl) Quaternion() *Quaternion {
	return &Quaternion{Value: ss.Get("quaternion")}
}

// SetQuaternion is ...
func (ss *SceneImpl) SetQuaternion(v *Quaternion) {
	ss.Set("quaternion", v.JSValue())
}

// ReceiveShadow is ...
func (ss *SceneImpl) ReceiveShadow() bool {
	return ss.Get("receiveShadow").Bool()
}

// SetReceiveShadow is ...
func (ss *SceneImpl) SetReceiveShadow(v bool) {
	ss.Set("receiveShadow", v)
}

// RenderOrder is ...
func (ss *SceneImpl) RenderOrder() float64 {
	return ss.Get("renderOrder").Float()
}

// SetRenderOrder is ...
func (ss *SceneImpl) SetRenderOrder(v float64) {
	ss.Set("renderOrder", v)
}

// Rotation is ...
func (ss *SceneImpl) Rotation() *Euler {
	return &Euler{Value: ss.Get("rotation")}
}

// SetRotation is ...
func (ss *SceneImpl) SetRotation(v *Euler) {
	ss.Set("rotation", v.JSValue())
}

// Scale is ...
func (ss *SceneImpl) Scale() *Vector3 {
	return &Vector3{Value: ss.Get("scale")}
}

// SetScale is ...
func (ss *SceneImpl) SetScale(v *Vector3) {
	ss.Set("scale", v.JSValue())
}

// Type is ...
func (ss *SceneImpl) Type() string {
	return ss.Get("type").String()
}

// SetType is ...
func (ss *SceneImpl) SetType(v string) {
	ss.Set("type", v)
}

// Up is ...
func (ss *SceneImpl) Up() *Vector3 {
	return &Vector3{Value: ss.Get("up")}
}

// SetUp is ...
func (ss *SceneImpl) SetUp(v *Vector3) {
	ss.Set("up", v.JSValue())
}

// UserData is ...
func (ss *SceneImpl) UserData() js.Value {
	return ss.Get("userData")
}

// SetUserData is ...
func (ss *SceneImpl) SetUserData(v js.Value) {
	ss.Set("userData", v)
}

// UUID is ...
func (ss *SceneImpl) UUID() string {
	return ss.Get("uuid").String()
}

// SetUUID is ...
func (ss *SceneImpl) SetUUID(v string) {
	ss.Set("uuid", v)
}

// Visible is ...
func (ss *SceneImpl) Visible() bool {
	return ss.Get("visible").Bool()
}

// SetVisible is ...
func (ss *SceneImpl) SetVisible(v bool) {
	ss.Set("visible", v)
}

// DefaultMatrixAutoUpdate is ...
func (ss *SceneImpl) DefaultMatrixAutoUpdate() bool {
	return ss.Get("DefaultMatrixAutoUpdate").Bool()
}

// SetDefaultMatrixAutoUpdate is ...
func (ss *SceneImpl) SetDefaultMatrixAutoUpdate(v bool) {
	ss.Set("DefaultMatrixAutoUpdate", v)
}

// DefaultUp is ...
func (ss *SceneImpl) DefaultUp() *Vector3 {
	return &Vector3{Value: ss.Get("DefaultUp")}
}

// SetDefaultUp is ...
func (ss *SceneImpl) SetDefaultUp(v *Vector3) {
	ss.Set("DefaultUp", v.JSValue())
}

// Add is ...
func (ss *SceneImpl) Add(object js.Value) Scene {
	return &SceneImpl{Value: ss.Call("add", object)}
}

// AddEventListener is ...
func (ss *SceneImpl) AddEventListener(typ string, listener js.Value) {
	ss.Call("addEventListener", typ, listener)
}

// ApplyMatrix is ...
func (ss *SceneImpl) ApplyMatrix(matrix *Matrix4) {
	ss.Call("applyMatrix", matrix.JSValue())
}

// ApplyQuaternion is ...
func (ss *SceneImpl) ApplyQuaternion(quaternion *Quaternion) Scene {
	return &SceneImpl{Value: ss.Call("applyQuaternion", quaternion)}
}

// Clone is ...
func (ss *SceneImpl) Clone(recursive bool) Scene {
	return &SceneImpl{Value: ss.Call("clone", recursive)}
}

// Copy is ...
func (ss *SceneImpl) Copy(source Scene, recursive bool) Scene {
	return &SceneImpl{Value: ss.Call("copy", source, recursive)}
}

// DispatchEvent is ...
func (ss *SceneImpl) DispatchEvent(event js.Value) {
	ss.Call("dispatchEvent", event)
}

// GetObjectByID is ...
func (ss *SceneImpl) GetObjectByID(id int) *Object3D {
	return &Object3D{Value: ss.Call("getObjectById", id)}
}

// GetObjectByName is ...
func (ss *SceneImpl) GetObjectByName(name string) *Object3D {
	return &Object3D{Value: ss.Call("getObjectByName", name)}
}

// GetObjectByProperty is ...
func (ss *SceneImpl) GetObjectByProperty(name string, value string) *Object3D {
	return &Object3D{Value: ss.Call("getObjectByProperty", name, value)}
}

// GetWorldDirection is ...
func (ss *SceneImpl) GetWorldDirection(target *Vector3) *Vector3 {
	return &Vector3{Value: ss.Call("getWorldDirection", target)}
}

// GetWorldPosition is ...
func (ss *SceneImpl) GetWorldPosition(target *Vector3) *Vector3 {
	return &Vector3{Value: ss.Call("getWorldPosition", target)}
}

// GetWorldQuaternion is ...
func (ss *SceneImpl) GetWorldQuaternion(target *Quaternion) *Quaternion {
	return &Quaternion{Value: ss.Call("getWorldQuaternion", target)}
}

// GetWorldScale is ...
func (ss *SceneImpl) GetWorldScale(target *Vector3) *Vector3 {
	return &Vector3{Value: ss.Call("getWorldScale", target)}
}

// HasEventListener is ...
func (ss *SceneImpl) HasEventListener(typ string, listener js.Value) bool {
	return ss.Call("hasEventListener", typ, listener).Bool()
}

// LocalToWorld is ...
func (ss *SceneImpl) LocalToWorld(vector *Vector3) *Vector3 {
	return &Vector3{Value: ss.Call("localToWorld", vector)}
}

// LookAt is ...
func (ss *SceneImpl) LookAt(vector *Vector3, y float64, z float64) {
	ss.Call("lookAt", vector, y, z)
}

// Raycast is ...
// func (ss *SceneImpl) Raycast(raycaster *Raycaster, intersects js.Value) {
// 	ss.Call("raycast", raycaster.JSValue(), intersects)
// }

// Remove is ...
func (ss *SceneImpl) Remove(object js.Value) Scene {
	return &SceneImpl{Value: ss.Call("remove", object)}
}

// RemoveEventListener is ...
func (ss *SceneImpl) RemoveEventListener(typ string, listener js.Value) {
	ss.Call("removeEventListener", typ, listener)
}

// RotateOnAxis is ...
func (ss *SceneImpl) RotateOnAxis(axis *Vector3, angle float64) Scene {
	return &SceneImpl{Value: ss.Call("rotateOnAxis", axis, angle)}
}

// RotateOnWorldAxis is ...
func (ss *SceneImpl) RotateOnWorldAxis(axis *Vector3, angle float64) Scene {
	return &SceneImpl{Value: ss.Call("rotateOnWorldAxis", axis, angle)}
}

// RotateX is ...
func (ss *SceneImpl) RotateX(angle float64) Scene {
	return &SceneImpl{Value: ss.Call("rotateX", angle)}
}

// RotateY is ...
func (ss *SceneImpl) RotateY(angle float64) Scene {
	return &SceneImpl{Value: ss.Call("rotateY", angle)}
}

// RotateZ is ...
func (ss *SceneImpl) RotateZ(angle float64) Scene {
	return &SceneImpl{Value: ss.Call("rotateZ", angle)}
}

// SetRotationFromAxisAngle is ...
func (ss *SceneImpl) SetRotationFromAxisAngle(axis *Vector3, angle float64) {
	ss.Call("setRotationFromAxisAngle", axis.JSValue(), angle)
}

// SetRotationFromEuler is ...
func (ss *SceneImpl) SetRotationFromEuler(euler *Euler) {
	ss.Call("setRotationFromEuler", euler.JSValue())
}

// SetRotationFromMatrix is ...
func (ss *SceneImpl) SetRotationFromMatrix(m *Matrix4) {
	ss.Call("setRotationFromMatrix", m.JSValue())
}

// SetRotationFromQuaternion is ...
func (ss *SceneImpl) SetRotationFromQuaternion(q *Quaternion) {
	ss.Call("setRotationFromQuaternion", q.JSValue())
}

// ToJSON is ...
func (ss *SceneImpl) ToJSON(meta js.Value) js.Value {
	return ss.Call("toJSON", meta)
}

// TranslateOnAxis is ...
func (ss *SceneImpl) TranslateOnAxis(axis *Vector3, distance float64) Scene {
	return &SceneImpl{Value: ss.Call("translateOnAxis", axis, distance)}
}

// TranslateX is ...
func (ss *SceneImpl) TranslateX(distance float64) Scene {
	return &SceneImpl{Value: ss.Call("translateX", distance)}
}

// TranslateY is ...
func (ss *SceneImpl) TranslateY(distance float64) Scene {
	return &SceneImpl{Value: ss.Call("translateY", distance)}
}

// TranslateZ is ...
func (ss *SceneImpl) TranslateZ(distance float64) Scene {
	return &SceneImpl{Value: ss.Call("translateZ", distance)}
}

// Traverse is ...
func (ss *SceneImpl) Traverse(callback js.Value) {
	ss.Call("traverse", callback)
}

// TraverseAncestors is ...
func (ss *SceneImpl) TraverseAncestors(callback js.Value) {
	ss.Call("traverseAncestors", callback)
}

// TraverseVisible is ...
func (ss *SceneImpl) TraverseVisible(callback js.Value) {
	ss.Call("traverseVisible", callback)
}

// UpdateMatrix is ...
func (ss *SceneImpl) UpdateMatrix() {
	ss.Call("updateMatrix")
}

// UpdateMatrixWorld is ...
func (ss *SceneImpl) UpdateMatrixWorld(force bool) {
	ss.Call("updateMatrixWorld", force)
}

// UpdateWorldMatrix is ...
func (ss *SceneImpl) UpdateWorldMatrix(updateParents bool, updateChildren bool) {
	ss.Call("updateWorldMatrix", updateParents, updateChildren)
}

// WorldToLocal is ...
func (ss *SceneImpl) WorldToLocal(vector *Vector3) *Vector3 {
	return &Vector3{Value: ss.Call("worldToLocal", vector)}
}
