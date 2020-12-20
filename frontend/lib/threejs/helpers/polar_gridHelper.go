package helpers

import (
	"app/frontend/lib/threejs"
)

// PolarGridHelper extend: []
type PolarGridHelper struct {
	threejs.Object3D
}

// NewPolarGridHelper is factory method for PolarGridHelper.
func NewPolarGridHelper(radius float64, radials float64) *PolarGridHelper {
	return &PolarGridHelper{
		threejs.NewDefaultObject3D(
			threejs.GetJsObject("PolarGridHelper").New(radius, radials),
		),
	}
}
