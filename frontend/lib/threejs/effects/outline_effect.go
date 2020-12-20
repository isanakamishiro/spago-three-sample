package effects

import (
	"app/frontend/lib/threejs"
	"syscall/js"
)

// OutlineEffect is ...
type OutlineEffect struct {
	js.Value
}

// NewOutlineEffect is ...
func NewOutlineEffect(outlineEffectModule js.Value, renderer threejs.Renderer) *OutlineEffect {
	return &OutlineEffect{Value: outlineEffectModule.Get("OutlineEffect").New(renderer.JSValue())}
}

// Render is ...
func (oe *OutlineEffect) Render(scene threejs.Scene, camera threejs.Camera) {
	oe.Call("render", scene.JSValue(), camera.JSValue())
}

// SetSize is ...
func (oe *OutlineEffect) SetSize(width float64, height float64, updateStyle bool) {
	oe.Call("setSize", width, height, updateStyle)
}
