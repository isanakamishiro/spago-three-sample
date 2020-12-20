package lights

import "app/frontend/lib/threejs"

// DirectionalLight is a light that gets emitted in a specific direction.
// This light will behave as though it is infinitely far away and the rays produced from it are all parallel.
// The common use case for this is to simulate daylight;
// the sun is far enough away that its position can be considered to be infinite, and all light rays coming from it are parallel.
//
// This light can cast shadows - see the DirectionalLightShadow page for details.
type DirectionalLight interface {
	threejs.Light

	SetPosition(v *threejs.Vector3)
	Shadow() DirectionalLightShadow
}

type directionalLightImp struct {
	threejs.Light
}

// NewDirectionalLight is factory method for DirectionalLight
// color - Numeric value of the RGB component of the color. Default is 0xffffff.
// intensity - Numeric value of the light's strength/intensity. Default is 1.
func NewDirectionalLight(color threejs.ColorValue, intensity threejs.LightIntensity) DirectionalLight {
	return &directionalLightImp{
		threejs.NewDefaultLightFromJSValue(
			threejs.GetJsObject("DirectionalLight").New(float64(color), float64(intensity)),
		),
	}
}

// Shadow is a DirectionalLightShadow used to calculate shadows for this light.
func (l *directionalLightImp) Shadow() DirectionalLightShadow {
	return NewDirectionalLightShadowFromJSValue(l.JSValue().Get("shadow"))
}

func (l *directionalLightImp) SetPosition(v *threejs.Vector3) {
	l.JSValue().Set("position", v.JSValue())
}
