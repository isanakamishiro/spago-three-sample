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
	// spago.RenderBody(views.NewFundamental3())
	// spago.RenderBody(views.NewFundamental4())
	// spago.RenderBody(views.NewFundamental5())
	spago.RenderBody(views.NewFundamental6())

	select {}
}
