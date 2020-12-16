package threejs

import (
	"syscall/js"
)

// Object3D extend: [EventDispatcher]
type Object3D struct {
	js.Value
}

// JSValue is ...
func (od *Object3D) JSValue() js.Value {
	return od.Value
}

// CastShadow is ...
func (od *Object3D) CastShadow() bool {
	return od.Get("castShadow").Bool()
}

// SetCastShadow is ...
func (od *Object3D) SetCastShadow(v bool) {
	od.Set("castShadow", v)
}

// Children is ...
func (od *Object3D) Children() js.Value {
	return od.Get("children")
}

// SetChildren is ...
func (od *Object3D) SetChildren(v js.Value) {
	od.Set("children", v)
}

// FrustumCulled is ...
func (od *Object3D) FrustumCulled() bool {
	return od.Get("frustumCulled").Bool()
}

// SetFrustumCulled is ...
func (od *Object3D) SetFrustumCulled(v bool) {
	od.Set("frustumCulled", v)
}

// ID is ...
func (od *Object3D) ID() int {
	return od.Get("id").Int()
}

// SetID is ...
func (od *Object3D) SetID(v int) {
	od.Set("id", v)
}

// IsObject3D is ...
func (od *Object3D) IsObject3D() bool {
	return od.Get("isObject3D").Bool()
}

// SetIsObject3D is ...
func (od *Object3D) SetIsObject3D(v bool) {
	od.Set("isObject3D", v)
}

// Layers is ...
func (od *Object3D) Layers() *Layers {
	return &Layers{Value: od.Get("layers")}
}

// SetLayers is ...
func (od *Object3D) SetLayers(v *Layers) {
	od.Set("layers", v.JSValue())
}

// Matrix is ...
func (od *Object3D) Matrix() *Matrix4 {
	return &Matrix4{Value: od.Get("matrix")}
}

// SetMatrix is ...
func (od *Object3D) SetMatrix(v *Matrix4) {
	od.Set("matrix", v.JSValue())
}

// MatrixAutoUpdate is ...
func (od *Object3D) MatrixAutoUpdate() bool {
	return od.Get("matrixAutoUpdate").Bool()
}

// SetMatrixAutoUpdate is ...
func (od *Object3D) SetMatrixAutoUpdate(v bool) {
	od.Set("matrixAutoUpdate", v)
}

// MatrixWorld is ...
func (od *Object3D) MatrixWorld() *Matrix4 {
	return &Matrix4{Value: od.Get("matrixWorld")}
}

// SetMatrixWorld is ...
func (od *Object3D) SetMatrixWorld(v *Matrix4) {
	od.Set("matrixWorld", v.JSValue())
}

// MatrixWorldNeedsUpdate is ...
func (od *Object3D) MatrixWorldNeedsUpdate() bool {
	return od.Get("matrixWorldNeedsUpdate").Bool()
}

// SetMatrixWorldNeedsUpdate is ...
func (od *Object3D) SetMatrixWorldNeedsUpdate(v bool) {
	od.Set("matrixWorldNeedsUpdate", v)
}

// ModelViewMatrix is ...
func (od *Object3D) ModelViewMatrix() *Matrix4 {
	return &Matrix4{Value: od.Get("modelViewMatrix")}
}

// SetModelViewMatrix is ...
func (od *Object3D) SetModelViewMatrix(v *Matrix4) {
	od.Set("modelViewMatrix", v.JSValue())
}

// Name is ...
func (od *Object3D) Name() string {
	return od.Get("name").String()
}

// SetName is ...
func (od *Object3D) SetName(v string) {
	od.Set("name", v)
}

// NormalMatrix is ...
func (od *Object3D) NormalMatrix() *Matrix3 {
	return &Matrix3{Value: od.Get("normalMatrix")}
}

// SetNormalMatrix is ...
func (od *Object3D) SetNormalMatrix(v *Matrix3) {
	od.Set("normalMatrix", v.JSValue())
}

// OnAfterRender is ...
func (od *Object3D) OnAfterRender() js.Value {
	return od.Get("onAfterRender")
}

// SetOnAfterRender is ...
func (od *Object3D) SetOnAfterRender(v js.Value) {
	od.Set("onAfterRender", v)
}

// OnBeforeRender is ...
func (od *Object3D) OnBeforeRender() js.Value {
	return od.Get("onBeforeRender")
}

// SetOnBeforeRender is ...
func (od *Object3D) SetOnBeforeRender(v js.Value) {
	od.Set("onBeforeRender", v)
}

// Parent is ...
func (od *Object3D) Parent() *Object3D {
	return &Object3D{Value: od.Get("parent")}
}

// SetParent is ...
func (od *Object3D) SetParent(v *Object3D) {
	od.Set("parent", v.JSValue())
}

// Position is ...
func (od *Object3D) Position() *Vector3 {
	return &Vector3{Value: od.Get("position")}
}

// SetPosition is ...
func (od *Object3D) SetPosition(v *Vector3) {
	od.Set("position", v.JSValue())
}

// Quaternion is ...
func (od *Object3D) Quaternion() *Quaternion {
	return &Quaternion{Value: od.Get("quaternion")}
}

// SetQuaternion is ...
func (od *Object3D) SetQuaternion(v *Quaternion) {
	od.Set("quaternion", v.JSValue())
}

// ReceiveShadow is ...
func (od *Object3D) ReceiveShadow() bool {
	return od.Get("receiveShadow").Bool()
}

// SetReceiveShadow is ...
func (od *Object3D) SetReceiveShadow(v bool) {
	od.Set("receiveShadow", v)
}

// RenderOrder is ...
func (od *Object3D) RenderOrder() float64 {
	return od.Get("renderOrder").Float()
}

// SetRenderOrder is ...
func (od *Object3D) SetRenderOrder(v float64) {
	od.Set("renderOrder", v)
}

// Rotation is ...
func (od *Object3D) Rotation() *Euler {
	return &Euler{Value: od.Get("rotation")}
}

// SetRotation is ...
func (od *Object3D) SetRotation(v *Euler) {
	od.Set("rotation", v.JSValue())
}

// Scale is ...
func (od *Object3D) Scale() *Vector3 {
	return &Vector3{Value: od.Get("scale")}
}

// SetScale is ...
func (od *Object3D) SetScale(v *Vector3) {
	od.Set("scale", v.JSValue())
}

// Type is ...
func (od *Object3D) Type() string {
	return od.Get("type").String()
}

// SetType is ...
func (od *Object3D) SetType(v string) {
	od.Set("type", v)
}

// Up is ...
func (od *Object3D) Up() *Vector3 {
	return &Vector3{Value: od.Get("up")}
}

// SetUp is ...
func (od *Object3D) SetUp(v *Vector3) {
	od.Set("up", v.JSValue())
}

// UserData is ...
func (od *Object3D) UserData() js.Value {
	return od.Get("userData")
}

// SetUserData is ...
func (od *Object3D) SetUserData(v js.Value) {
	od.Set("userData", v)
}

// UUID is ...
func (od *Object3D) UUID() string {
	return od.Get("uuid").String()
}

// SetUUID is ...
func (od *Object3D) SetUUID(v string) {
	od.Set("uuid", v)
}

// Visible is ...
func (od *Object3D) Visible() bool {
	return od.Get("visible").Bool()
}

// SetVisible is ...
func (od *Object3D) SetVisible(v bool) {
	od.Set("visible", v)
}

// DefaultMatrixAutoUpdate is ...
func (od *Object3D) DefaultMatrixAutoUpdate() bool {
	return od.Get("DefaultMatrixAutoUpdate").Bool()
}

// SetDefaultMatrixAutoUpdate is ...
func (od *Object3D) SetDefaultMatrixAutoUpdate(v bool) {
	od.Set("DefaultMatrixAutoUpdate", v)
}

// DefaultUp is ...
func (od *Object3D) DefaultUp() *Vector3 {
	return &Vector3{Value: od.Get("DefaultUp")}
}

// SetDefaultUp is ...
func (od *Object3D) SetDefaultUp(v *Vector3) {
	od.Set("DefaultUp", v.JSValue())
}

// Add is ...
func (od *Object3D) Add(object js.Value) *Object3D {
	return &Object3D{Value: od.Call("add", object)}
}

// AddEventListener is ...
func (od *Object3D) AddEventListener(typ string, listener js.Value) {
	od.Call("addEventListener", typ, listener)
}

// ApplyMatrix is ...
func (od *Object3D) ApplyMatrix(matrix *Matrix4) {
	od.Call("applyMatrix", matrix.JSValue())
}

// ApplyQuaternion is ...
func (od *Object3D) ApplyQuaternion(quaternion *Quaternion) *Object3D {
	return &Object3D{Value: od.Call("applyQuaternion", quaternion)}
}

// Clone is ...
func (od *Object3D) Clone(recursive bool) *Object3D {
	return &Object3D{Value: od.Call("clone", recursive)}
}

// Copy is ...
func (od *Object3D) Copy(source *Object3D, recursive bool) *Object3D {
	return &Object3D{Value: od.Call("copy", source, recursive)}
}

// DispatchEvent is ...
func (od *Object3D) DispatchEvent(event js.Value) {
	od.Call("dispatchEvent", event)
}

// GetObjectByID is ...
func (od *Object3D) GetObjectByID(id int) *Object3D {
	return &Object3D{Value: od.Call("getObjectById", id)}
}

// GetObjectByName is ...
func (od *Object3D) GetObjectByName(name string) *Object3D {
	return &Object3D{Value: od.Call("getObjectByName", name)}
}

// GetObjectByProperty is ...
func (od *Object3D) GetObjectByProperty(name string, value string) *Object3D {
	return &Object3D{Value: od.Call("getObjectByProperty", name, value)}
}

// GetWorldDirection is ...
func (od *Object3D) GetWorldDirection(target *Vector3) *Vector3 {
	return &Vector3{Value: od.Call("getWorldDirection", target)}
}

// GetWorldPosition is ...
func (od *Object3D) GetWorldPosition(target *Vector3) *Vector3 {
	return &Vector3{Value: od.Call("getWorldPosition", target)}
}

// GetWorldQuaternion is ...
func (od *Object3D) GetWorldQuaternion(target *Quaternion) *Quaternion {
	return &Quaternion{Value: od.Call("getWorldQuaternion", target)}
}

// GetWorldScale is ...
func (od *Object3D) GetWorldScale(target *Vector3) *Vector3 {
	return &Vector3{Value: od.Call("getWorldScale", target)}
}

// HasEventListener is ...
func (od *Object3D) HasEventListener(typ string, listener js.Value) bool {
	return od.Call("hasEventListener", typ, listener).Bool()
}

// LocalToWorld is ...
func (od *Object3D) LocalToWorld(vector *Vector3) *Vector3 {
	return &Vector3{Value: od.Call("localToWorld", vector)}
}

// LookAt is ...
func (od *Object3D) LookAt(vector *Vector3, y float64, z float64) {
	od.Call("lookAt", vector, y, z)
}

// Raycast is ...
// func (od *Object3D) Raycast(raycaster *Raycaster, intersects js.Value) {
// 	od.Call("raycast", raycaster.JSValue(), intersects)
// }

// Remove is ...
func (od *Object3D) Remove(object js.Value) *Object3D {
	return &Object3D{Value: od.Call("remove", object)}
}

// RemoveEventListener is ...
func (od *Object3D) RemoveEventListener(typ string, listener js.Value) {
	od.Call("removeEventListener", typ, listener)
}

// RotateOnAxis is ...
func (od *Object3D) RotateOnAxis(axis *Vector3, angle float64) *Object3D {
	return &Object3D{Value: od.Call("rotateOnAxis", axis, angle)}
}

// RotateOnWorldAxis is ...
func (od *Object3D) RotateOnWorldAxis(axis *Vector3, angle float64) *Object3D {
	return &Object3D{Value: od.Call("rotateOnWorldAxis", axis, angle)}
}

// RotateX is ...
func (od *Object3D) RotateX(angle float64) *Object3D {
	return &Object3D{Value: od.Call("rotateX", angle)}
}

// RotateY is ...
func (od *Object3D) RotateY(angle float64) *Object3D {
	return &Object3D{Value: od.Call("rotateY", angle)}
}

// RotateZ is ...
func (od *Object3D) RotateZ(angle float64) *Object3D {
	return &Object3D{Value: od.Call("rotateZ", angle)}
}

// SetRotationFromAxisAngle is ...
func (od *Object3D) SetRotationFromAxisAngle(axis *Vector3, angle float64) {
	od.Call("setRotationFromAxisAngle", axis.JSValue(), angle)
}

// SetRotationFromEuler is ...
func (od *Object3D) SetRotationFromEuler(euler *Euler) {
	od.Call("setRotationFromEuler", euler.JSValue())
}

// SetRotationFromMatrix is ...
func (od *Object3D) SetRotationFromMatrix(m *Matrix4) {
	od.Call("setRotationFromMatrix", m.JSValue())
}

// SetRotationFromQuaternion is ...
func (od *Object3D) SetRotationFromQuaternion(q *Quaternion) {
	od.Call("setRotationFromQuaternion", q.JSValue())
}

// ToJSON is ...
func (od *Object3D) ToJSON(meta js.Value) js.Value {
	return od.Call("toJSON", meta)
}

// TranslateOnAxis is ...
func (od *Object3D) TranslateOnAxis(axis *Vector3, distance float64) *Object3D {
	return &Object3D{Value: od.Call("translateOnAxis", axis, distance)}
}

// TranslateX is ...
func (od *Object3D) TranslateX(distance float64) *Object3D {
	return &Object3D{Value: od.Call("translateX", distance)}
}

// TranslateY is ...
func (od *Object3D) TranslateY(distance float64) *Object3D {
	return &Object3D{Value: od.Call("translateY", distance)}
}

// TranslateZ is ...
func (od *Object3D) TranslateZ(distance float64) *Object3D {
	return &Object3D{Value: od.Call("translateZ", distance)}
}

// Traverse is ...
func (od *Object3D) Traverse(callback js.Value) {
	od.Call("traverse", callback)
}

// TraverseAncestors is ...
func (od *Object3D) TraverseAncestors(callback js.Value) {
	od.Call("traverseAncestors", callback)
}

// TraverseVisible is ...
func (od *Object3D) TraverseVisible(callback js.Value) {
	od.Call("traverseVisible", callback)
}

// UpdateMatrix is ...
func (od *Object3D) UpdateMatrix() {
	od.Call("updateMatrix")
}

// UpdateMatrixWorld is ...
func (od *Object3D) UpdateMatrixWorld(force bool) {
	od.Call("updateMatrixWorld", force)
}

// UpdateWorldMatrix is ...
func (od *Object3D) UpdateWorldMatrix(updateParents bool, updateChildren bool) {
	od.Call("updateWorldMatrix", updateParents, updateChildren)
}

// WorldToLocal is ...
func (od *Object3D) WorldToLocal(vector *Vector3) *Vector3 {
	return &Vector3{Value: od.Call("worldToLocal", vector)}
}
