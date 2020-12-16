package threejs

import (
	"syscall/js"
)

// Euler extend: []
type Euler struct {
	js.Value
}

// JSValue is ...
func (ee *Euler) JSValue() js.Value {
	return ee.Value
}

// OnChangeCallback is ...
func (ee *Euler) OnChangeCallback() js.Value {
	return ee.Get("onChangeCallback")
}

// SetOnChangeCallback is ...
func (ee *Euler) SetOnChangeCallback(v js.Value) {
	ee.Set("onChangeCallback", v)
}

// Order is ...
func (ee *Euler) Order() string {
	return ee.Get("order").String()
}

// SetOrder is ...
func (ee *Euler) SetOrder(v string) {
	ee.Set("order", v)
}

// X is ...
func (ee *Euler) X() float64 {
	return ee.Get("x").Float()
}

// SetX is ...
func (ee *Euler) SetX(v float64) {
	ee.Set("x", v)
}

// Y is ...
func (ee *Euler) Y() float64 {
	return ee.Get("y").Float()
}

// SetY is ...
func (ee *Euler) SetY(v float64) {
	ee.Set("y", v)
}

// Z is ...
func (ee *Euler) Z() float64 {
	return ee.Get("z").Float()
}

// SetZ is ...
func (ee *Euler) SetZ(v float64) {
	ee.Set("z", v)
}

// DefaultOrder is ...
func (ee *Euler) DefaultOrder() string {
	return ee.Get("DefaultOrder").String()
}

// SetDefaultOrder is ...
func (ee *Euler) SetDefaultOrder(v string) {
	ee.Set("DefaultOrder", v)
}

// RotationOrders is ...
func (ee *Euler) RotationOrders() js.Value {
	return ee.Get("RotationOrders")
}

// SetRotationOrders is ...
func (ee *Euler) SetRotationOrders(v js.Value) {
	ee.Set("RotationOrders", v)
}

// Clone is ...
func (ee *Euler) Clone() *Euler {
	return &Euler{Value: ee.Call("clone")}
}

// Copy is ...
func (ee *Euler) Copy(euler *Euler) *Euler {
	return &Euler{Value: ee.Call("copy", euler)}
}

// Equals is ...
func (ee *Euler) Equals(euler *Euler) bool {
	return ee.Call("equals", euler).Bool()
}

// FromArray is ...
func (ee *Euler) FromArray(xyzo js.Value) *Euler {
	return &Euler{Value: ee.Call("fromArray", xyzo)}
}

// OnChange is ...
func (ee *Euler) OnChange(callback js.Value) *Euler {
	return &Euler{Value: ee.Call("onChange", callback)}
}

// Reorder is ...
func (ee *Euler) Reorder(newOrder string) *Euler {
	return &Euler{Value: ee.Call("reorder", newOrder)}
}

// Set2 is ...
func (ee *Euler) Set2(x float64, y float64, z float64, order string) *Euler {
	return &Euler{Value: ee.Call("set", x, y, z, order)}
}

// SetFromQuaternion is ...
func (ee *Euler) SetFromQuaternion(q *Quaternion, order string, update bool) *Euler {
	return &Euler{Value: ee.Call("setFromQuaternion", q, order, update)}
}

// SetFromRotationMatrix is ...
func (ee *Euler) SetFromRotationMatrix(m *Matrix4, order string, update bool) *Euler {
	return &Euler{Value: ee.Call("setFromRotationMatrix", m, order, update)}
}

// SetFromVector3 is ...
func (ee *Euler) SetFromVector3(v *Vector3, order string) *Euler {
	return &Euler{Value: ee.Call("setFromVector3", v, order)}
}

// ToArray is ...
func (ee *Euler) ToArray(array js.Value, offset int) js.Value {
	return ee.Call("toArray", array, offset)
}

// ToVector3 is ...
func (ee *Euler) ToVector3(optionalResult *Vector3) *Vector3 {
	return &Vector3{Value: ee.Call("toVector3", optionalResult)}
}
