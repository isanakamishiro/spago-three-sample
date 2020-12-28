package main

import (
	"app/frontend/views"

	"github.com/nobonobo/spago"
)

func main() {

	spago.AddStylesheet("./assets/app.css")

	// ThreeJs experiment

	// spago.RenderBody(views.NewMikuView())
	// spago.RenderBody(views.NewTop())
	// spago.RenderBody(views.NewFundamental1())
	// spago.RenderBody(views.NewFundamental2())
	spago.RenderBody(views.NewFundamental3())

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
