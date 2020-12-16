package main

import (
	"app/frontend/lib/threejs"
	"app/frontend/views"

	"github.com/nobonobo/spago"
)

func main() {

	threejs := threejs.NewThreeJs(
		spago.LoadModuleAs("THREE", "https://cdnjs.cloudflare.com/ajax/libs/three.js/r122/three.module.js"),
	)

	spago.RenderBody(views.NewMikuView(threejs))

	select {}
}

// loadScript synchronous javascript loader
// func loadScript(url string) {

// 	document := js.Global().Get("document")

// 	ch := make(chan bool)
// 	script := document.Call("createElement", "script")
// 	script.Set("src", url)
// 	var fn js.Func
// 	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
// 		defer fn.Release()
// 		close(ch)
// 		return nil
// 	})
// 	script.Call("addeventListener", "load", fn)
// 	document.Get("head").Call("appendChild", script)
// 	<-ch
// }
