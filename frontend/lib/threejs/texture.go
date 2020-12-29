package threejs

import "syscall/js"

// Texture is ...
type Texture interface {
	EventDispatcher

	SetWrapS(s Wrapping)
	SetWrapT(t Wrapping)
	SetMagFilter(v TextureFilter)
	SetMinFilter(v TextureFilter)

	Repeat() *Vector2
}

type textureImp struct {
	js.Value
}

// NewDefaultTextureFromJSValue creates Texture from js.Value
func NewDefaultTextureFromJSValue(v js.Value) Texture {
	return &textureImp{
		Value: v,
	}
}

// AddEventListener is ...
func (c *textureImp) AddEventListener(typ string, listener js.Value) {
	c.Call("addEventListener", typ, listener)
}

// RemoveEventListener is ...
func (c *textureImp) RemoveEventListener(typ string, listener js.Value) {
	c.Call("removeEventListener", typ, listener)
}

// HasEventListener is ...
func (c *textureImp) HasEventListener(typ string, listener js.Value) bool {
	return c.Call("hasEventListener", typ, listener).Bool()
}

// DispatchEvent is ...
func (c *textureImp) DispatchEvent(event js.Value) {
	c.Call("dispatchEvent", event)
}

// SetWrapS is ...
func (c *textureImp) SetWrapS(s Wrapping) {
	c.Set("wrapS", s.JSValue())
}

// SetWrapT is ...
func (c *textureImp) SetWrapT(t Wrapping) {
	c.Set("wrapT", t.JSValue())
}

// SetMagFilter is ...
func (c *textureImp) SetMagFilter(v TextureFilter) {
	c.Set("magFilter", v.JSValue())
}

// SetMinFilter is ...
func (c *textureImp) SetMinFilter(v TextureFilter) {
	c.Set("minFilter", v.JSValue())
}

// Repeat is ...
func (c *textureImp) Repeat() *Vector2 {
	return &Vector2{
		Value: c.Get("repeat"),
	}
}
