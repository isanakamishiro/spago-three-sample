package threejs

import (
	"syscall/js"

	"github.com/nobonobo/spago"
)

const modulePath = "./assets/threejs/build/three.module.js"

var module js.Value

func init() {
	module = spago.LoadModuleAs("THREE", modulePath)
}

// GetJsObject is getter for JavaScript object of ThreeJs.
func GetJsObject(key string) js.Value {
	return module.Get(key)
}

// setJsObject is setter for JavaScript object of ThreeJs.
func setJsObject(key string, v interface{}) {
	module.Set(key, v)
}
