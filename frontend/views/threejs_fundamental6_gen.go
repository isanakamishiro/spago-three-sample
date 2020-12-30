package views

import (
	"github.com/nobonobo/spago"
)

// Render ...
func (c *Fundamental6) Render() spago.HTML {
	return spago.Tag("body", 
		spago.Tag("canvas", 			
			spago.A("id", spago.S(`c`)),
		),
		spago.Tag("div", 			
			spago.A("class", spago.S(`split`)),
			spago.Tag("div", 				
				spago.A("id", spago.S(`view1`)),
				spago.A("tabindex", spago.S(`1`)),
			),
			spago.Tag("div", 				
				spago.A("id", spago.S(`view2`)),
				spago.A("tabindex", spago.S(`2`)),
			),
		),
	)
}
