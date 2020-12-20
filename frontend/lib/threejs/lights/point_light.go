package lights

import (
	"app/frontend/lib/threejs"
)

// PointLight is a light that gets emitted from a single point in all directions.
// A common use case for this is to replicate the light emitted from a bare lightbulb.
//
// This light can cast shadows - see PointLightShadow page for details.
type PointLight interface {
	threejs.Light
}

type pointLightImp struct {
	threejs.Light
}

// NewPointLight Creates a new PointLight.
// color - (optional) hexadecimal color of the light. Default is 0xffffff (white).
// intensity - (optional) numeric value of the light's strength/intensity. Default is 1.
// distance - Maximum range of the light. Default is 0 (no limit).
// decay - The amount the light dims along the distance of the light. Default is 1. For physically correct lighting, set this to 2.
func NewPointLight(
	color threejs.ColorValue,
	intensity threejs.LightIntensity,
	distance float64,
	decay float64,
) AmbientLight {

	return &ambientLightImp{
		threejs.NewDefaultLightFromJSValue(
			threejs.GetJsObject("PointLight").New(
				float64(color),
				float64(intensity),
				distance,
				decay,
			),
		),
	}
}
