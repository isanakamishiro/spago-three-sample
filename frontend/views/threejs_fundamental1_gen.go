package views

import (
	"github.com/nobonobo/spago"
)

// Render ...
func (c *Fundamental1) Render() spago.HTML {
	return spago.Tag("body", 
		spago.Tag("canvas", 			
			spago.A("id", spago.S(`c`)),
		),
	)
}