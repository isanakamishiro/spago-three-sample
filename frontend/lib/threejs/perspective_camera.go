package threejs

import (
	"syscall/js"
)

// PerspectiveCamera extend: [Camera]
type PerspectiveCamera struct {
	js.Value
}

// JSValue is ...
func (pc *PerspectiveCamera) JSValue() js.Value {
	return pc.Value
}

// Aspect is ...
func (pc *PerspectiveCamera) Aspect() float64 {
	return pc.Get("aspect").Float()
}

// SetAspect is ...
func (pc *PerspectiveCamera) SetAspect(v float64) {
	pc.Set("aspect", v)
}

// CastShadow is ...
func (pc *PerspectiveCamera) CastShadow() bool {
	return pc.Get("castShadow").Bool()
}

// SetCastShadow is ...
func (pc *PerspectiveCamera) SetCastShadow(v bool) {
	pc.Set("castShadow", v)
}

// Children is ...
func (pc *PerspectiveCamera) Children() js.Value {
	return pc.Get("children")
}

// SetChildren is ...
func (pc *PerspectiveCamera) SetChildren(v js.Value) {
	pc.Set("children", v)
}

// Far is ...
func (pc *PerspectiveCamera) Far() float64 {
	return pc.Get("far").Float()
}

// SetFar is ...
func (pc *PerspectiveCamera) SetFar(v float64) {
	pc.Set("far", v)
}

// FilmGauge is ...
func (pc *PerspectiveCamera) FilmGauge() float64 {
	return pc.Get("filmGauge").Float()
}

// SetFilmGauge is ...
func (pc *PerspectiveCamera) SetFilmGauge(v float64) {
	pc.Set("filmGauge", v)
}

// FilmOffset is ...
func (pc *PerspectiveCamera) FilmOffset() float64 {
	return pc.Get("filmOffset").Float()
}

// SetFilmOffset is ...
func (pc *PerspectiveCamera) SetFilmOffset(v float64) {
	pc.Set("filmOffset", v)
}

// Focus is ...
func (pc *PerspectiveCamera) Focus() float64 {
	return pc.Get("focus").Float()
}

// SetFocus is ...
func (pc *PerspectiveCamera) SetFocus(v float64) {
	pc.Set("focus", v)
}

// Fov is ...
func (pc *PerspectiveCamera) Fov() float64 {
	return pc.Get("fov").Float()
}

// SetFov is ...
func (pc *PerspectiveCamera) SetFov(v float64) {
	pc.Set("fov", v)
}

// FrustumCulled is ...
func (pc *PerspectiveCamera) FrustumCulled() bool {
	return pc.Get("frustumCulled").Bool()
}

// SetFrustumCulled is ...
func (pc *PerspectiveCamera) SetFrustumCulled(v bool) {
	pc.Set("frustumCulled", v)
}

// ID is ...
func (pc *PerspectiveCamera) ID() int {
	return pc.Get("id").Int()
}

// SetID is ...
func (pc *PerspectiveCamera) SetID(v int) {
	pc.Set("id", v)
}

// IsCamera is ...
func (pc *PerspectiveCamera) IsCamera() bool {
	return pc.Get("isCamera").Bool()
}

// SetIsCamera is ...
func (pc *PerspectiveCamera) SetIsCamera(v bool) {
	pc.Set("isCamera", v)
}

// IsObject3D is ...
func (pc *PerspectiveCamera) IsObject3D() bool {
	return pc.Get("isObject3D").Bool()
}

// SetIsObject3D is ...
func (pc *PerspectiveCamera) SetIsObject3D(v bool) {
	pc.Set("isObject3D", v)
}

// IsPerspectiveCamera is ...
func (pc *PerspectiveCamera) IsPerspectiveCamera() bool {
	return pc.Get("isPerspectiveCamera").Bool()
}

// SetIsPerspectiveCamera is ...
func (pc *PerspectiveCamera) SetIsPerspectiveCamera(v bool) {
	pc.Set("isPerspectiveCamera", v)
}

// Layers is ...
func (pc *PerspectiveCamera) Layers() *Layers {
	return &Layers{Value: pc.Get("layers")}
}

// SetLayers is ...
func (pc *PerspectiveCamera) SetLayers(v *Layers) {
	pc.Set("layers", v.JSValue())
}

// Matrix is ...
func (pc *PerspectiveCamera) Matrix() *Matrix4 {
	return &Matrix4{Value: pc.Get("matrix")}
}

// SetMatrix is ...
func (pc *PerspectiveCamera) SetMatrix(v *Matrix4) {
	pc.Set("matrix", v.JSValue())
}

// MatrixAutoUpdate is ...
func (pc *PerspectiveCamera) MatrixAutoUpdate() bool {
	return pc.Get("matrixAutoUpdate").Bool()
}

// SetMatrixAutoUpdate is ...
func (pc *PerspectiveCamera) SetMatrixAutoUpdate(v bool) {
	pc.Set("matrixAutoUpdate", v)
}

// MatrixWorld is ...
func (pc *PerspectiveCamera) MatrixWorld() *Matrix4 {
	return &Matrix4{Value: pc.Get("matrixWorld")}
}

// SetMatrixWorld is ...
func (pc *PerspectiveCamera) SetMatrixWorld(v *Matrix4) {
	pc.Set("matrixWorld", v.JSValue())
}

// MatrixWorldInverse is ...
func (pc *PerspectiveCamera) MatrixWorldInverse() *Matrix4 {
	return &Matrix4{Value: pc.Get("matrixWorldInverse")}
}

// SetMatrixWorldInverse is ...
func (pc *PerspectiveCamera) SetMatrixWorldInverse(v *Matrix4) {
	pc.Set("matrixWorldInverse", v.JSValue())
}

// MatrixWorldNeedsUpdate is ...
func (pc *PerspectiveCamera) MatrixWorldNeedsUpdate() bool {
	return pc.Get("matrixWorldNeedsUpdate").Bool()
}

// SetMatrixWorldNeedsUpdate is ...
func (pc *PerspectiveCamera) SetMatrixWorldNeedsUpdate(v bool) {
	pc.Set("matrixWorldNeedsUpdate", v)
}

// ModelViewMatrix is ...
func (pc *PerspectiveCamera) ModelViewMatrix() *Matrix4 {
	return &Matrix4{Value: pc.Get("modelViewMatrix")}
}

// SetModelViewMatrix is ...
func (pc *PerspectiveCamera) SetModelViewMatrix(v *Matrix4) {
	pc.Set("modelViewMatrix", v.JSValue())
}

// Name is ...
func (pc *PerspectiveCamera) Name() string {
	return pc.Get("name").String()
}

// SetName is ...
func (pc *PerspectiveCamera) SetName(v string) {
	pc.Set("name", v)
}

// Near is ...
func (pc *PerspectiveCamera) Near() float64 {
	return pc.Get("near").Float()
}

// SetNear is ...
func (pc *PerspectiveCamera) SetNear(v float64) {
	pc.Set("near", v)
}

// NormalMatrix is ...
func (pc *PerspectiveCamera) NormalMatrix() *Matrix3 {
	return &Matrix3{Value: pc.Get("normalMatrix")}
}

// SetNormalMatrix is ...
func (pc *PerspectiveCamera) SetNormalMatrix(v *Matrix3) {
	pc.Set("normalMatrix", v.JSValue())
}

// OnAfterRender is ...
func (pc *PerspectiveCamera) OnAfterRender() js.Value {
	return pc.Get("onAfterRender")
}

// SetOnAfterRender is ...
func (pc *PerspectiveCamera) SetOnAfterRender(v js.Value) {
	pc.Set("onAfterRender", v)
}

// OnBeforeRender is ...
func (pc *PerspectiveCamera) OnBeforeRender() js.Value {
	return pc.Get("onBeforeRender")
}

// SetOnBeforeRender is ...
func (pc *PerspectiveCamera) SetOnBeforeRender(v js.Value) {
	pc.Set("onBeforeRender", v)
}

// Parent is ...
func (pc *PerspectiveCamera) Parent() *Object3D {
	return &Object3D{Value: pc.Get("parent")}
}

// SetParent is ...
func (pc *PerspectiveCamera) SetParent(v *Object3D) {
	pc.Set("parent", v.JSValue())
}

// Position is ...
func (pc *PerspectiveCamera) Position() *Vector3 {
	return &Vector3{Value: pc.Get("position")}
}

// SetPosition is ...
func (pc *PerspectiveCamera) SetPosition(v *Vector3) {
	pc.Set("position", v.JSValue())
}

// ProjectionMatrix is ...
func (pc *PerspectiveCamera) ProjectionMatrix() *Matrix4 {
	return &Matrix4{Value: pc.Get("projectionMatrix")}
}

// SetProjectionMatrix is ...
func (pc *PerspectiveCamera) SetProjectionMatrix(v *Matrix4) {
	pc.Set("projectionMatrix", v.JSValue())
}

// Quaternion is ...
func (pc *PerspectiveCamera) Quaternion() *Quaternion {
	return &Quaternion{Value: pc.Get("quaternion")}
}

// SetQuaternion is ...
func (pc *PerspectiveCamera) SetQuaternion(v *Quaternion) {
	pc.Set("quaternion", v.JSValue())
}

// ReceiveShadow is ...
func (pc *PerspectiveCamera) ReceiveShadow() bool {
	return pc.Get("receiveShadow").Bool()
}

// SetReceiveShadow is ...
func (pc *PerspectiveCamera) SetReceiveShadow(v bool) {
	pc.Set("receiveShadow", v)
}

// RenderOrder is ...
func (pc *PerspectiveCamera) RenderOrder() float64 {
	return pc.Get("renderOrder").Float()
}

// SetRenderOrder is ...
func (pc *PerspectiveCamera) SetRenderOrder(v float64) {
	pc.Set("renderOrder", v)
}

// Rotation is ...
func (pc *PerspectiveCamera) Rotation() *Euler {
	return &Euler{Value: pc.Get("rotation")}
}

// SetRotation is ...
func (pc *PerspectiveCamera) SetRotation(v *Euler) {
	pc.Set("rotation", v.JSValue())
}

// Scale is ...
func (pc *PerspectiveCamera) Scale() *Vector3 {
	return &Vector3{Value: pc.Get("scale")}
}

// SetScale is ...
func (pc *PerspectiveCamera) SetScale(v *Vector3) {
	pc.Set("scale", v.JSValue())
}

// Type is ...
func (pc *PerspectiveCamera) Type() string {
	return pc.Get("type").String()
}

// SetType is ...
func (pc *PerspectiveCamera) SetType(v string) {
	pc.Set("type", v)
}

// Up is ...
func (pc *PerspectiveCamera) Up() *Vector3 {
	return &Vector3{Value: pc.Get("up")}
}

// SetUp is ...
func (pc *PerspectiveCamera) SetUp(v *Vector3) {
	pc.Set("up", v.JSValue())
}

// UserData is ...
func (pc *PerspectiveCamera) UserData() js.Value {
	return pc.Get("userData")
}

// SetUserData is ...
func (pc *PerspectiveCamera) SetUserData(v js.Value) {
	pc.Set("userData", v)
}

// UUID is ...
func (pc *PerspectiveCamera) UUID() string {
	return pc.Get("uuid").String()
}

// SetUUID is ...
func (pc *PerspectiveCamera) SetUUID(v string) {
	pc.Set("uuid", v)
}

// View is ...
func (pc *PerspectiveCamera) View() js.Value {
	return pc.Get("view")
}

// SetView is ...
func (pc *PerspectiveCamera) SetView(v js.Value) {
	pc.Set("view", v)
}

// Visible is ...
func (pc *PerspectiveCamera) Visible() bool {
	return pc.Get("visible").Bool()
}

// SetVisible is ...
func (pc *PerspectiveCamera) SetVisible(v bool) {
	pc.Set("visible", v)
}

// Zoom is ...
func (pc *PerspectiveCamera) Zoom() float64 {
	return pc.Get("zoom").Float()
}

// SetZoom is ...
func (pc *PerspectiveCamera) SetZoom(v float64) {
	pc.Set("zoom", v)
}

// DefaultMatrixAutoUpdate is ...
func (pc *PerspectiveCamera) DefaultMatrixAutoUpdate() bool {
	return pc.Get("DefaultMatrixAutoUpdate").Bool()
}

// SetDefaultMatrixAutoUpdate is ...
func (pc *PerspectiveCamera) SetDefaultMatrixAutoUpdate(v bool) {
	pc.Set("DefaultMatrixAutoUpdate", v)
}

// DefaultUp is ...
func (pc *PerspectiveCamera) DefaultUp() *Vector3 {
	return &Vector3{Value: pc.Get("DefaultUp")}
}

// SetDefaultUp is ...
func (pc *PerspectiveCamera) SetDefaultUp(v *Vector3) {
	pc.Set("DefaultUp", v.JSValue())
}

// Add is ...
func (pc *PerspectiveCamera) Add(object js.Value) Camera {
	return &PerspectiveCamera{Value: pc.Call("add", object)}
}

// AddEventListener is ...
func (pc *PerspectiveCamera) AddEventListener(typ string, listener js.Value) {
	pc.Call("addEventListener", typ, listener)
}

// ApplyMatrix is ...
func (pc *PerspectiveCamera) ApplyMatrix(matrix *Matrix4) {
	pc.Call("applyMatrix", matrix.JSValue())
}

// ApplyQuaternion is ...
func (pc *PerspectiveCamera) ApplyQuaternion(quaternion *Quaternion) Camera {
	return &PerspectiveCamera{Value: pc.Call("applyQuaternion", quaternion)}
}

// ClearViewOffset is ...
func (pc *PerspectiveCamera) ClearViewOffset() {
	pc.Call("clearViewOffset")
}

// Clone is ...
func (pc *PerspectiveCamera) Clone(recursive bool) Camera {
	return &PerspectiveCamera{Value: pc.Call("clone", recursive)}
}

// Copy is ...
func (pc *PerspectiveCamera) Copy(source Camera, recursive bool) Camera {
	return &PerspectiveCamera{Value: pc.Call("copy", source.JSValue(), recursive)}
}

// DispatchEvent is ...
func (pc *PerspectiveCamera) DispatchEvent(event js.Value) {
	pc.Call("dispatchEvent", event)
}

// GetEffectiveFOV is ...
func (pc *PerspectiveCamera) GetEffectiveFOV() float64 {
	return pc.Call("getEffectiveFOV").Float()
}

// GetFilmHeight is ...
func (pc *PerspectiveCamera) GetFilmHeight() float64 {
	return pc.Call("getFilmHeight").Float()
}

// GetFilmWidth is ...
func (pc *PerspectiveCamera) GetFilmWidth() float64 {
	return pc.Call("getFilmWidth").Float()
}

// GetFocalLength is ...
func (pc *PerspectiveCamera) GetFocalLength() float64 {
	return pc.Call("getFocalLength").Float()
}

// GetObjectByID is ...
func (pc *PerspectiveCamera) GetObjectByID(id int) *Object3D {
	return &Object3D{Value: pc.Call("getObjectById", id)}
}

// GetObjectByName is ...
func (pc *PerspectiveCamera) GetObjectByName(name string) *Object3D {
	return &Object3D{Value: pc.Call("getObjectByName", name)}
}

// GetObjectByProperty is ...
func (pc *PerspectiveCamera) GetObjectByProperty(name string, value string) *Object3D {
	return &Object3D{Value: pc.Call("getObjectByProperty", name, value)}
}

// GetWorldDirection is ...
func (pc *PerspectiveCamera) GetWorldDirection(target *Vector3) *Vector3 {
	return &Vector3{Value: pc.Call("getWorldDirection", target)}
}

// GetWorldPosition is ...
func (pc *PerspectiveCamera) GetWorldPosition(target *Vector3) *Vector3 {
	return &Vector3{Value: pc.Call("getWorldPosition", target)}
}

// GetWorldQuaternion is ...
func (pc *PerspectiveCamera) GetWorldQuaternion(target *Quaternion) *Quaternion {
	return &Quaternion{Value: pc.Call("getWorldQuaternion", target)}
}

// GetWorldScale is ...
func (pc *PerspectiveCamera) GetWorldScale(target *Vector3) *Vector3 {
	return &Vector3{Value: pc.Call("getWorldScale", target)}
}

// HasEventListener is ...
func (pc *PerspectiveCamera) HasEventListener(typ string, listener js.Value) bool {
	return pc.Call("hasEventListener", typ, listener).Bool()
}

// LocalToWorld is ...
func (pc *PerspectiveCamera) LocalToWorld(vector *Vector3) *Vector3 {
	return &Vector3{Value: pc.Call("localToWorld", vector)}
}

// LookAt is ...
func (pc *PerspectiveCamera) LookAt(vector *Vector3, y float64, z float64) {
	pc.Call("lookAt", vector, y, z)
}

// Raycast is ...
// func (pc *PerspectiveCamera) Raycast(raycaster *Raycaster, intersects js.Value) {
// 	pc.Call("raycast", raycaster.JSValue(), intersects)
// }

// Remove is ...
func (pc *PerspectiveCamera) Remove(object js.Value) Camera {
	return &PerspectiveCamera{Value: pc.Call("remove", object)}
}

// RemoveEventListener is ...
func (pc *PerspectiveCamera) RemoveEventListener(typ string, listener js.Value) {
	pc.Call("removeEventListener", typ, listener)
}

// RotateOnAxis is ...
func (pc *PerspectiveCamera) RotateOnAxis(axis *Vector3, angle float64) Camera {
	return &PerspectiveCamera{Value: pc.Call("rotateOnAxis", axis, angle)}
}

// RotateOnWorldAxis is ...
func (pc *PerspectiveCamera) RotateOnWorldAxis(axis *Vector3, angle float64) Camera {
	return &PerspectiveCamera{Value: pc.Call("rotateOnWorldAxis", axis, angle)}
}

// RotateX is ...
func (pc *PerspectiveCamera) RotateX(angle float64) Camera {
	return &PerspectiveCamera{Value: pc.Call("rotateX", angle)}
}

// RotateY is ...
func (pc *PerspectiveCamera) RotateY(angle float64) Camera {
	return &PerspectiveCamera{Value: pc.Call("rotateY", angle)}
}

// RotateZ is ...
func (pc *PerspectiveCamera) RotateZ(angle float64) Camera {
	return &PerspectiveCamera{Value: pc.Call("rotateZ", angle)}
}

// SetFocalLength is ...
func (pc *PerspectiveCamera) SetFocalLength(focalLength float64) {
	pc.Call("setFocalLength", focalLength)
}

// SetLens is ...
func (pc *PerspectiveCamera) SetLens(focalLength float64, frameHeight float64) {
	pc.Call("setLens", focalLength, frameHeight)
}

// SetRotationFromAxisAngle is ...
func (pc *PerspectiveCamera) SetRotationFromAxisAngle(axis *Vector3, angle float64) {
	pc.Call("setRotationFromAxisAngle", axis.JSValue(), angle)
}

// SetRotationFromEuler is ...
func (pc *PerspectiveCamera) SetRotationFromEuler(euler *Euler) {
	pc.Call("setRotationFromEuler", euler.JSValue())
}

// SetRotationFromMatrix is ...
func (pc *PerspectiveCamera) SetRotationFromMatrix(m *Matrix4) {
	pc.Call("setRotationFromMatrix", m.JSValue())
}

// SetRotationFromQuaternion is ...
func (pc *PerspectiveCamera) SetRotationFromQuaternion(q *Quaternion) {
	pc.Call("setRotationFromQuaternion", q.JSValue())
}

// SetViewOffset is ...
func (pc *PerspectiveCamera) SetViewOffset(fullWidth float64, fullHeight float64, x float64, y float64, width float64, height float64) {
	pc.Call("setViewOffset", fullWidth, fullHeight, x, y, width, height)
}

// ToJSON is ...
func (pc *PerspectiveCamera) ToJSON(meta js.Value) js.Value {
	return pc.Call("toJSON", meta)
}

// TranslateOnAxis is ...
func (pc *PerspectiveCamera) TranslateOnAxis(axis *Vector3, distance float64) Camera {
	return &PerspectiveCamera{Value: pc.Call("translateOnAxis", axis, distance)}
}

// TranslateX is ...
func (pc *PerspectiveCamera) TranslateX(distance float64) Camera {
	return &PerspectiveCamera{Value: pc.Call("translateX", distance)}
}

// TranslateY is ...
func (pc *PerspectiveCamera) TranslateY(distance float64) Camera {
	return &PerspectiveCamera{Value: pc.Call("translateY", distance)}
}

// TranslateZ is ...
func (pc *PerspectiveCamera) TranslateZ(distance float64) Camera {
	return &PerspectiveCamera{Value: pc.Call("translateZ", distance)}
}

// Traverse is ...
func (pc *PerspectiveCamera) Traverse(callback js.Value) {
	pc.Call("traverse", callback)
}

// TraverseAncestors is ...
func (pc *PerspectiveCamera) TraverseAncestors(callback js.Value) {
	pc.Call("traverseAncestors", callback)
}

// TraverseVisible is ...
func (pc *PerspectiveCamera) TraverseVisible(callback js.Value) {
	pc.Call("traverseVisible", callback)
}

// UpdateMatrix is ...
func (pc *PerspectiveCamera) UpdateMatrix() {
	pc.Call("updateMatrix")
}

// UpdateMatrixWorld is ...
func (pc *PerspectiveCamera) UpdateMatrixWorld(force bool) {
	pc.Call("updateMatrixWorld", force)
}

// UpdateProjectionMatrix is ...
func (pc *PerspectiveCamera) UpdateProjectionMatrix() {
	pc.Call("updateProjectionMatrix")
}

// UpdateWorldMatrix is ...
func (pc *PerspectiveCamera) UpdateWorldMatrix(updateParents bool, updateChildren bool) {
	pc.Call("updateWorldMatrix", updateParents, updateChildren)
}

// WorldToLocal is ...
func (pc *PerspectiveCamera) WorldToLocal(vector *Vector3) *Vector3 {
	return &Vector3{Value: pc.Call("worldToLocal", vector)}
}
