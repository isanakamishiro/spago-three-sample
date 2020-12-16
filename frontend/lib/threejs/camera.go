package threejs

import (
	"syscall/js"
)

// Camera is camera object for three.js
type Camera interface {
	JSValue() js.Value
	CastShadow() bool
	SetCastShadow(v bool)
	Children() js.Value
	SetChildren(v js.Value)
	FrustumCulled() bool
	SetFrustumCulled(v bool)
	ID() int
	SetID(v int)
	IsCamera() bool
	SetIsCamera(v bool)
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
	MatrixWorldInverse() *Matrix4
	SetMatrixWorldInverse(v *Matrix4)
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
	Parent() *Object3D
	SetParent(v *Object3D)
	Position() *Vector3
	SetPosition(v *Vector3)
	ProjectionMatrix() *Matrix4
	SetProjectionMatrix(v *Matrix4)
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
	Add(object js.Value) Camera
	AddEventListener(typ string, listener js.Value)
	ApplyMatrix(matrix *Matrix4)
	ApplyQuaternion(quaternion *Quaternion) Camera
	Clone(recursive bool) Camera
	Copy(source Camera, recursive bool) Camera
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
	Remove(object js.Value) Camera
	RemoveEventListener(typ string, listener js.Value)
	RotateOnAxis(axis *Vector3, angle float64) Camera
	RotateOnWorldAxis(axis *Vector3, angle float64) Camera
	RotateX(angle float64) Camera
	RotateY(angle float64) Camera
	RotateZ(angle float64) Camera
	SetRotationFromAxisAngle(axis *Vector3, angle float64)
	SetRotationFromEuler(euler *Euler)
	SetRotationFromMatrix(m *Matrix4)
	SetRotationFromQuaternion(q *Quaternion)
	ToJSON(meta js.Value) js.Value
	TranslateOnAxis(axis *Vector3, distance float64) Camera
	TranslateX(distance float64) Camera
	TranslateY(distance float64) Camera
	TranslateZ(distance float64) Camera
	Traverse(callback js.Value)
	TraverseAncestors(callback js.Value)
	TraverseVisible(callback js.Value)
	UpdateMatrix()
	UpdateMatrixWorld(force bool)
	UpdateWorldMatrix(updateParents bool, updateChildren bool)
	WorldToLocal(vector *Vector3) *Vector3
}

// CameraImpl extend: [Object3D]
type CameraImpl struct {
	js.Value
}

// JSValue is ...
func (cc *CameraImpl) JSValue() js.Value {
	return cc.Value
}

// CastShadow is ...
func (cc *CameraImpl) CastShadow() bool {
	return cc.Get("castShadow").Bool()
}

// SetCastShadow is ...
func (cc *CameraImpl) SetCastShadow(v bool) {
	cc.Set("castShadow", v)
}

// Children is ...
func (cc *CameraImpl) Children() js.Value {
	return cc.Get("children")
}

// SetChildren is ...
func (cc *CameraImpl) SetChildren(v js.Value) {
	cc.Set("children", v)
}

// FrustumCulled is ...
func (cc *CameraImpl) FrustumCulled() bool {
	return cc.Get("frustumCulled").Bool()
}

// SetFrustumCulled is ...
func (cc *CameraImpl) SetFrustumCulled(v bool) {
	cc.Set("frustumCulled", v)
}

// ID is ...
func (cc *CameraImpl) ID() int {
	return cc.Get("id").Int()
}

// SetID is ...
func (cc *CameraImpl) SetID(v int) {
	cc.Set("id", v)
}

// IsCamera is ...
func (cc *CameraImpl) IsCamera() bool {
	return cc.Get("isCamera").Bool()
}

// SetIsCamera is ...
func (cc *CameraImpl) SetIsCamera(v bool) {
	cc.Set("isCamera", v)
}

// IsObject3D is ...
func (cc *CameraImpl) IsObject3D() bool {
	return cc.Get("isObject3D").Bool()
}

// SetIsObject3D is ...
func (cc *CameraImpl) SetIsObject3D(v bool) {
	cc.Set("isObject3D", v)
}

// Layers is ...
func (cc *CameraImpl) Layers() *Layers {
	return &Layers{Value: cc.Get("layers")}
}

// SetLayers is ...
func (cc *CameraImpl) SetLayers(v *Layers) {
	cc.Set("layers", v.JSValue())
}

// Matrix is ...
func (cc *CameraImpl) Matrix() *Matrix4 {
	return &Matrix4{Value: cc.Get("matrix")}
}

// SetMatrix is ...
func (cc *CameraImpl) SetMatrix(v *Matrix4) {
	cc.Set("matrix", v.JSValue())
}

// MatrixAutoUpdate is ...
func (cc *CameraImpl) MatrixAutoUpdate() bool {
	return cc.Get("matrixAutoUpdate").Bool()
}

// SetMatrixAutoUpdate is ...
func (cc *CameraImpl) SetMatrixAutoUpdate(v bool) {
	cc.Set("matrixAutoUpdate", v)
}

// MatrixWorld is ...
func (cc *CameraImpl) MatrixWorld() *Matrix4 {
	return &Matrix4{Value: cc.Get("matrixWorld")}
}

// SetMatrixWorld is ...
func (cc *CameraImpl) SetMatrixWorld(v *Matrix4) {
	cc.Set("matrixWorld", v.JSValue())
}

// MatrixWorldInverse is ...
func (cc *CameraImpl) MatrixWorldInverse() *Matrix4 {
	return &Matrix4{Value: cc.Get("matrixWorldInverse")}
}

// SetMatrixWorldInverse is ...
func (cc *CameraImpl) SetMatrixWorldInverse(v *Matrix4) {
	cc.Set("matrixWorldInverse", v.JSValue())
}

// MatrixWorldNeedsUpdate is ...
func (cc *CameraImpl) MatrixWorldNeedsUpdate() bool {
	return cc.Get("matrixWorldNeedsUpdate").Bool()
}

// SetMatrixWorldNeedsUpdate is ...
func (cc *CameraImpl) SetMatrixWorldNeedsUpdate(v bool) {
	cc.Set("matrixWorldNeedsUpdate", v)
}

// ModelViewMatrix is ...
func (cc *CameraImpl) ModelViewMatrix() *Matrix4 {
	return &Matrix4{Value: cc.Get("modelViewMatrix")}
}

// SetModelViewMatrix is ...
func (cc *CameraImpl) SetModelViewMatrix(v *Matrix4) {
	cc.Set("modelViewMatrix", v.JSValue())
}

// Name is ...
func (cc *CameraImpl) Name() string {
	return cc.Get("name").String()
}

// SetName is ...
func (cc *CameraImpl) SetName(v string) {
	cc.Set("name", v)
}

// NormalMatrix is ...
func (cc *CameraImpl) NormalMatrix() *Matrix3 {
	return &Matrix3{Value: cc.Get("normalMatrix")}
}

// SetNormalMatrix is ...
func (cc *CameraImpl) SetNormalMatrix(v *Matrix3) {
	cc.Set("normalMatrix", v.JSValue())
}

// OnAfterRender is ...
func (cc *CameraImpl) OnAfterRender() js.Value {
	return cc.Get("onAfterRender")
}

// SetOnAfterRender is ...
func (cc *CameraImpl) SetOnAfterRender(v js.Value) {
	cc.Set("onAfterRender", v)
}

// OnBeforeRender is ...
func (cc *CameraImpl) OnBeforeRender() js.Value {
	return cc.Get("onBeforeRender")
}

// SetOnBeforeRender is ...
func (cc *CameraImpl) SetOnBeforeRender(v js.Value) {
	cc.Set("onBeforeRender", v)
}

// Parent is ...
func (cc *CameraImpl) Parent() *Object3D {
	return &Object3D{Value: cc.Get("parent")}
}

// SetParent is ...
func (cc *CameraImpl) SetParent(v *Object3D) {
	cc.Set("parent", v.JSValue())
}

// Position is ...
func (cc *CameraImpl) Position() *Vector3 {
	return &Vector3{Value: cc.Get("position")}
}

// SetPosition is ...
func (cc *CameraImpl) SetPosition(v *Vector3) {
	cc.Set("position", v.JSValue())
}

// ProjectionMatrix is ...
func (cc *CameraImpl) ProjectionMatrix() *Matrix4 {
	return &Matrix4{Value: cc.Get("projectionMatrix")}
}

// SetProjectionMatrix is ...
func (cc *CameraImpl) SetProjectionMatrix(v *Matrix4) {
	cc.Set("projectionMatrix", v.JSValue())
}

// Quaternion is ...
func (cc *CameraImpl) Quaternion() *Quaternion {
	return &Quaternion{Value: cc.Get("quaternion")}
}

// SetQuaternion is ...
func (cc *CameraImpl) SetQuaternion(v *Quaternion) {
	cc.Set("quaternion", v.JSValue())
}

// ReceiveShadow is ...
func (cc *CameraImpl) ReceiveShadow() bool {
	return cc.Get("receiveShadow").Bool()
}

// SetReceiveShadow is ...
func (cc *CameraImpl) SetReceiveShadow(v bool) {
	cc.Set("receiveShadow", v)
}

// RenderOrder is ...
func (cc *CameraImpl) RenderOrder() float64 {
	return cc.Get("renderOrder").Float()
}

// SetRenderOrder is ...
func (cc *CameraImpl) SetRenderOrder(v float64) {
	cc.Set("renderOrder", v)
}

// Rotation is ...
func (cc *CameraImpl) Rotation() *Euler {
	return &Euler{Value: cc.Get("rotation")}
}

// SetRotation is ...
func (cc *CameraImpl) SetRotation(v *Euler) {
	cc.Set("rotation", v.JSValue())
}

// Scale is ...
func (cc *CameraImpl) Scale() *Vector3 {
	return &Vector3{Value: cc.Get("scale")}
}

// SetScale is ...
func (cc *CameraImpl) SetScale(v *Vector3) {
	cc.Set("scale", v.JSValue())
}

// Type is ...
func (cc *CameraImpl) Type() string {
	return cc.Get("type").String()
}

// SetType is ...
func (cc *CameraImpl) SetType(v string) {
	cc.Set("type", v)
}

// Up is ...
func (cc *CameraImpl) Up() *Vector3 {
	return &Vector3{Value: cc.Get("up")}
}

// SetUp is ...
func (cc *CameraImpl) SetUp(v *Vector3) {
	cc.Set("up", v.JSValue())
}

// UserData is ...
func (cc *CameraImpl) UserData() js.Value {
	return cc.Get("userData")
}

// SetUserData is ...
func (cc *CameraImpl) SetUserData(v js.Value) {
	cc.Set("userData", v)
}

// UUID is ...
func (cc *CameraImpl) UUID() string {
	return cc.Get("uuid").String()
}

// SetUUID is ...
func (cc *CameraImpl) SetUUID(v string) {
	cc.Set("uuid", v)
}

// Visible is ...
func (cc *CameraImpl) Visible() bool {
	return cc.Get("visible").Bool()
}

// SetVisible is ...
func (cc *CameraImpl) SetVisible(v bool) {
	cc.Set("visible", v)
}

// DefaultMatrixAutoUpdate is ...
func (cc *CameraImpl) DefaultMatrixAutoUpdate() bool {
	return cc.Get("DefaultMatrixAutoUpdate").Bool()
}

// SetDefaultMatrixAutoUpdate is ...
func (cc *CameraImpl) SetDefaultMatrixAutoUpdate(v bool) {
	cc.Set("DefaultMatrixAutoUpdate", v)
}

// DefaultUp is ...
func (cc *CameraImpl) DefaultUp() *Vector3 {
	return &Vector3{Value: cc.Get("DefaultUp")}
}

// SetDefaultUp is ...
func (cc *CameraImpl) SetDefaultUp(v *Vector3) {
	cc.Set("DefaultUp", v.JSValue())
}

// Add is ...
func (cc *CameraImpl) Add(object js.Value) Camera {
	return &CameraImpl{Value: cc.Call("add", object)}
}

// AddEventListener is ...
func (cc *CameraImpl) AddEventListener(typ string, listener js.Value) {
	cc.Call("addEventListener", typ, listener)
}

// ApplyMatrix is ...
func (cc *CameraImpl) ApplyMatrix(matrix *Matrix4) {
	cc.Call("applyMatrix", matrix.JSValue())
}

// ApplyQuaternion is ...
func (cc *CameraImpl) ApplyQuaternion(quaternion *Quaternion) Camera {
	return &CameraImpl{Value: cc.Call("applyQuaternion", quaternion)}
}

// Clone is ...
func (cc *CameraImpl) Clone(recursive bool) Camera {
	return &CameraImpl{Value: cc.Call("clone", recursive)}
}

// Copy is ...
func (cc *CameraImpl) Copy(source Camera, recursive bool) Camera {
	return &CameraImpl{Value: cc.Call("copy", source.JSValue(), recursive)}
}

// DispatchEvent is ...
func (cc *CameraImpl) DispatchEvent(event js.Value) {
	cc.Call("dispatchEvent", event)
}

// GetObjectByID is ...
func (cc *CameraImpl) GetObjectByID(id int) *Object3D {
	return &Object3D{Value: cc.Call("getObjectById", id)}
}

// GetObjectByName is ...
func (cc *CameraImpl) GetObjectByName(name string) *Object3D {
	return &Object3D{Value: cc.Call("getObjectByName", name)}
}

// GetObjectByProperty is ...
func (cc *CameraImpl) GetObjectByProperty(name string, value string) *Object3D {
	return &Object3D{Value: cc.Call("getObjectByProperty", name, value)}
}

// GetWorldDirection is ...
func (cc *CameraImpl) GetWorldDirection(target *Vector3) *Vector3 {
	return &Vector3{Value: cc.Call("getWorldDirection", target)}
}

// GetWorldPosition is ...
func (cc *CameraImpl) GetWorldPosition(target *Vector3) *Vector3 {
	return &Vector3{Value: cc.Call("getWorldPosition", target)}
}

// GetWorldQuaternion is ...
func (cc *CameraImpl) GetWorldQuaternion(target *Quaternion) *Quaternion {
	return &Quaternion{Value: cc.Call("getWorldQuaternion", target)}
}

// GetWorldScale is ...
func (cc *CameraImpl) GetWorldScale(target *Vector3) *Vector3 {
	return &Vector3{Value: cc.Call("getWorldScale", target)}
}

// HasEventListener is ...
func (cc *CameraImpl) HasEventListener(typ string, listener js.Value) bool {
	return cc.Call("hasEventListener", typ, listener).Bool()
}

// LocalToWorld is ...
func (cc *CameraImpl) LocalToWorld(vector *Vector3) *Vector3 {
	return &Vector3{Value: cc.Call("localToWorld", vector)}
}

// LookAt is ...
func (cc *CameraImpl) LookAt(vector *Vector3, y float64, z float64) {
	cc.Call("lookAt", vector, y, z)
}

// Raycast is ...
// func (cc *CameraImpl) Raycast(raycaster *Raycaster, intersects js.Value) {
// 	cc.Call("raycast", raycaster.JSValue(), intersects)
// }

// Remove is ...
func (cc *CameraImpl) Remove(object js.Value) Camera {
	return &CameraImpl{Value: cc.Call("remove", object)}
}

// RemoveEventListener is ...
func (cc *CameraImpl) RemoveEventListener(typ string, listener js.Value) {
	cc.Call("removeEventListener", typ, listener)
}

// RotateOnAxis is ...
func (cc *CameraImpl) RotateOnAxis(axis *Vector3, angle float64) Camera {
	return &CameraImpl{Value: cc.Call("rotateOnAxis", axis, angle)}
}

// RotateOnWorldAxis is ...
func (cc *CameraImpl) RotateOnWorldAxis(axis *Vector3, angle float64) Camera {
	return &CameraImpl{Value: cc.Call("rotateOnWorldAxis", axis, angle)}
}

// RotateX is ...
func (cc *CameraImpl) RotateX(angle float64) Camera {
	return &CameraImpl{Value: cc.Call("rotateX", angle)}
}

// RotateY is ...
func (cc *CameraImpl) RotateY(angle float64) Camera {
	return &CameraImpl{Value: cc.Call("rotateY", angle)}
}

// RotateZ is ...
func (cc *CameraImpl) RotateZ(angle float64) Camera {
	return &CameraImpl{Value: cc.Call("rotateZ", angle)}
}

// SetRotationFromAxisAngle is ...
func (cc *CameraImpl) SetRotationFromAxisAngle(axis *Vector3, angle float64) {
	cc.Call("setRotationFromAxisAngle", axis.JSValue(), angle)
}

// SetRotationFromEuler is ...
func (cc *CameraImpl) SetRotationFromEuler(euler *Euler) {
	cc.Call("setRotationFromEuler", euler.JSValue())
}

// SetRotationFromMatrix is ...
func (cc *CameraImpl) SetRotationFromMatrix(m *Matrix4) {
	cc.Call("setRotationFromMatrix", m.JSValue())
}

// SetRotationFromQuaternion is ...
func (cc *CameraImpl) SetRotationFromQuaternion(q *Quaternion) {
	cc.Call("setRotationFromQuaternion", q.JSValue())
}

// ToJSON is ...
func (cc *CameraImpl) ToJSON(meta js.Value) js.Value {
	return cc.Call("toJSON", meta)
}

// TranslateOnAxis is ...
func (cc *CameraImpl) TranslateOnAxis(axis *Vector3, distance float64) Camera {
	return &CameraImpl{Value: cc.Call("translateOnAxis", axis, distance)}
}

// TranslateX is ...
func (cc *CameraImpl) TranslateX(distance float64) Camera {
	return &CameraImpl{Value: cc.Call("translateX", distance)}
}

// TranslateY is ...
func (cc *CameraImpl) TranslateY(distance float64) Camera {
	return &CameraImpl{Value: cc.Call("translateY", distance)}
}

// TranslateZ is ...
func (cc *CameraImpl) TranslateZ(distance float64) Camera {
	return &CameraImpl{Value: cc.Call("translateZ", distance)}
}

// Traverse is ...
func (cc *CameraImpl) Traverse(callback js.Value) {
	cc.Call("traverse", callback)
}

// TraverseAncestors is ...
func (cc *CameraImpl) TraverseAncestors(callback js.Value) {
	cc.Call("traverseAncestors", callback)
}

// TraverseVisible is ...
func (cc *CameraImpl) TraverseVisible(callback js.Value) {
	cc.Call("traverseVisible", callback)
}

// UpdateMatrix is ...
func (cc *CameraImpl) UpdateMatrix() {
	cc.Call("updateMatrix")
}

// UpdateMatrixWorld is ...
func (cc *CameraImpl) UpdateMatrixWorld(force bool) {
	cc.Call("updateMatrixWorld", force)
}

// UpdateWorldMatrix is ...
func (cc *CameraImpl) UpdateWorldMatrix(updateParents bool, updateChildren bool) {
	cc.Call("updateWorldMatrix", updateParents, updateChildren)
}

// WorldToLocal is ...
func (cc *CameraImpl) WorldToLocal(vector *Vector3) *Vector3 {
	return &Vector3{Value: cc.Call("worldToLocal", vector)}
}
