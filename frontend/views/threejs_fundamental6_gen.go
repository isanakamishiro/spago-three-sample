package views

import (
	"app/frontend/components"
	"github.com/nobonobo/spago"
)

// Render ...
func (c *Fundamental6) Render() spago.HTML {
	return spago.Tag("body", 
		spago.Tag("div", 			
			spago.A("id", spago.S(`main`)),
			spago.A("class", spago.S(`uk-cover-container uk-height-1-1`)),
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
			spago.Tag("div", 				
				spago.A("id", spago.S(`gui`)),
				spago.A("class", spago.S(`uk-position-small uk-position-bottom-right`)),
			),
		),
		spago.C(&components.Header{}),
	)
}
