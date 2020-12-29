package threejs

import "syscall/js"

// Texture is ...
type Texture interface {
	EventDispatcher
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
