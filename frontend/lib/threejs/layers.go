package threejs

import (
	"syscall/js"
)

// Layers extend: []
type Layers struct {
	js.Value
}

// NewLayers ...
// func NewLayers() *Layers {
// 	return &Layers{Value: get("Layers").New()}
// }

// JSValue ...
func (ll *Layers) JSValue() js.Value {
	return ll.Value
}

// Mask ...
func (ll *Layers) Mask() float64 {
	return ll.Get("mask").Float()
}

// SetMask ...
func (ll *Layers) SetMask(v float64) {
	ll.Set("mask", v)
}

// Disable ...
func (ll *Layers) Disable(channel float64) {
	ll.Call("disable", channel)
}

// Enable ...
func (ll *Layers) Enable(channel float64) {
	ll.Call("enable", channel)
}

// Set2 ...
func (ll *Layers) Set2(channel float64) {
	ll.Call("set", channel)
}

// Test ...
func (ll *Layers) Test(layers *Layers) bool {
	return ll.Call("test", layers).Bool()
}

// Toggle ...
func (ll *Layers) Toggle(channel float64) {
	ll.Call("toggle", channel)
}
