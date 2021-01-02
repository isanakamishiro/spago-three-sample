package main

import (
	"app/frontend/views"
	"syscall/js"

	"github.com/nobonobo/spago"
	"github.com/nobonobo/spago/router"
)

func main() {

	// spago.AddStylesheet("./assets/app.css")
	loadScript("./assets/threejs/ex/js/libs/ammo.wasm.js")

	r := router.New()
	r.Handle("/", func(key string) {
		spago.SetTitle("Top")
		spago.RenderBody(views.NewTop())
	})
	r.Handle("/f1", func(key string) {
		spago.SetTitle("Fundamental1:Simple Cubes")
		spago.RenderBody(views.NewFundamental1())
	})
	r.Handle("/f2", func(key string) {
		spago.SetTitle("Fundamental2:Interactive Cubes")
		spago.RenderBody(views.NewFundamental2())
	})
	r.Handle("/f3", func(key string) {
		spago.SetTitle("Fundamental3:")
		spago.RenderBody(views.NewFundamental3())
	})
	r.Handle("/f4", func(key string) {
		spago.SetTitle("Fundamental4:")
		spago.RenderBody(views.NewFundamental4())
	})
	r.Handle("/f5", func(key string) {
		spago.SetTitle("Fundamental5:")
		spago.RenderBody(views.NewFundamental5())
	})
	r.Handle("/f6", func(key string) {
		spago.SetTitle("Fundamental6:")
		spago.RenderBody(views.NewFundamental6())
	})
	r.Handle("/f7", func(key string) {
		spago.SetTitle("Fundamental7:Shadow1")
		spago.RenderBody(views.NewFundamental7())
	})
	r.Handle("/f8", func(key string) {
		spago.SetTitle("Fundamental8:Shadow2")
		spago.RenderBody(views.NewFundamental8())
	})
	r.Handle("/f9", func(key string) {
		spago.SetTitle("Fundamental9:Fog")
		spago.RenderBody(views.NewFundamental9())
	})
	r.Handle("/miku1", func(key string) {
		spago.SetTitle("Fundamental9:Fog")
		spago.RenderBody(views.NewMikuView())
	})

	if err := r.Start(); err != nil {
		println(err)
		spago.RenderBody(router.NotFoundPage())
	}
	// ThreeJs experiment

	// spago.RenderBody(views.NewMikuView())
	// spago.RenderBody(views.NewTop())
	// spago.RenderBody(views.NewFundamental1())
	// spago.RenderBody(views.NewFundamental2())
	// spago.RenderBody(views.NewFundamental3())
	// spago.RenderBody(views.NewFundamental4())
	// spago.RenderBody(views.NewFundamental5())
	// spago.RenderBody(views.NewFundamental6())
	// spago.RenderBody(views.NewFundamental7())
	// spago.RenderBody(views.NewFundamental8())
	// spago.RenderBody(views.NewFundamental9())

	// spago.RenderBody(views.NewTestView())

	select {}
}

// loadScript synchronous javascript loader
func loadScript(url string) {

	document := js.Global().Get("document")

	ch := make(chan bool)
	script := document.Call("createElement", "script")
	script.Set("src", url)
	var fn js.Func
	fn = js.FuncOf(func(this js.Value, args []js.Value) interface{} {
		defer fn.Release()
		close(ch)
		return nil
	})
	script.Call("addEventListener", "load", fn)
	document.Get("head").Call("appendChild", script)
	<-ch
}
