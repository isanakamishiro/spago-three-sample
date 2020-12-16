package views

import (
	"github.com/nobonobo/spago"
)

// Render ...
func (c *MikuView) Render() spago.HTML {
	return spago.Tag("body", 		
		spago.A("style", spago.S(`margin: 0px;`)),
	)
}
