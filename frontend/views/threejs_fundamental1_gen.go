package views

import (
	"app/frontend/components"
	"github.com/nobonobo/spago"
)

// Render ...
func (c *Fundamental1) Render() spago.HTML {
	return spago.Tag("body", 
		spago.Tag("div", 			
			spago.A("id", spago.S(`main`)),
			spago.A("class", spago.S(`uk-cover-container uk-height-1-1`)),
			spago.Tag("canvas", 				
				spago.A("id", spago.S(`c`)),
			),
		),
		spago.C(&components.Header{}),
	)
}
