package datgui

import (
	"app/frontend/lib/threejs"
	"log"
	"syscall/js"
)

const guiModulePath = "./assets/threejs/ex/jsm/libs/dat.gui.module.js"

var guiModule js.Value

func init() {

	m := threejs.LoadModule([]string{"GUI"}, guiModulePath)
	if len(m) == 0 {
		log.Fatal("GUI module could not be loaded.")
	}
	guiModule = m[0]
}

// Controller is an "abstract" class that represents a given property of an object.
type Controller interface {
	Name(name string) Controller

	OnChange(fn js.Func) Controller
}

// GUI is A lightweight controller library for JavaScript.
// It allows you to easily manipulate variables and fire functions on the fly.
type GUI interface {
	JSValue() js.Value

	AddFolder(name string) GUI

	Add(object js.Value, property string, min float64, max float64) Controller
	AddSimply(object js.Value, property string) Controller
	AddColor(object map[string]interface{}, property string) Controller

	Open()
	Close()

	Hide()
	Show()

	// add() Controller
	// add
}

// contollerImp is default implementation of Controller
type contollerImp struct {
	js.Value
}

// guiImp is default implementation of GUI
type guiImp struct {
	js.Value
}

// NewGUI is ...
func NewGUI() GUI {
	return &guiImp{
		Value: guiModule.New(),
	}
}

// newGUIFromJSValue is ...
func newGUIFromJSValue(v js.Value) GUI {
	return &guiImp{
		Value: v,
	}
}

// newControllerFromJSValue is ...
func newControllerFromJSValue(v js.Value) Controller {
	return &contollerImp{
		Value: v,
	}
}

//
// GUI method
//

// JSValue return the javascript object
func (c *guiImp) JSValue() js.Value {
	return c.Value
}

// AddFolder adds a new color controller to the GUI.
func (c *guiImp) AddFolder(name string) GUI {
	return newGUIFromJSValue(
		c.Call("addFolder", name),
	)
}

// Add adds a new object...
func (c *guiImp) Add(object js.Value, property string, min float64, max float64) Controller {
	return newControllerFromJSValue(
		c.Call("add", object, property, min, max),
	)
}

// AddSimply adds a new object as simple.
func (c *guiImp) AddSimply(object js.Value, property string) Controller {
	return newControllerFromJSValue(
		c.Call("add", object, property),
	)
}

// AddColor adds a new color controller to the GUI.
func (c *guiImp) AddColor(object map[string]interface{}, property string) Controller {
	return newControllerFromJSValue(
		c.Call("addColor", object, property),
	)
}

// Open opens the GUI.
func (c *guiImp) Open() {
	c.Call("open")
}

// Close closes the GUI.
func (c *guiImp) Close() {
	c.Call("close")
}

// Hide hides the GUI.
func (c *guiImp) Hide() {
	c.Call("hide")
}

// Show shows the GUI.
func (c *guiImp) Show() {
	c.Call("show")
}

//
// Contoller method
//

// Name sets the name of the controller.
func (c *contollerImp) Name(name string) Controller {
	return newControllerFromJSValue(
		c.Call("name", name),
	)
}

// OnChange specify that a function fire every time someone changes the value with this Controller.
func (c *contollerImp) OnChange(fn js.Func) Controller {
	return newControllerFromJSValue(
		c.Call("onChange", fn),
	)
}
