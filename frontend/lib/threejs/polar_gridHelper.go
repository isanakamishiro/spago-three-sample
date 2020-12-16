package threejs

import (
	"syscall/js"
)

// PolarGridHelper extend: []
type PolarGridHelper struct {
	js.Value
}

// JSValue is ..
func (pgh *PolarGridHelper) JSValue() js.Value {
	return pgh.Value
}

// Rotation is ...
func (pgh *PolarGridHelper) Rotation() *Euler {
	return &Euler{Value: pgh.Get("rotation")}
}

// SetRotation is ...
func (pgh *PolarGridHelper) SetRotation(v *Euler) {
	pgh.Set("rotation", v.JSValue())
}

// Position is ...
func (pgh *PolarGridHelper) Position() *Vector3 {
	return &Vector3{Value: pgh.Get("position")}
}

// SetPosition is ...
func (pgh *PolarGridHelper) SetPosition(v *Vector3) {
	pgh.Set("position", v.JSValue())
}
