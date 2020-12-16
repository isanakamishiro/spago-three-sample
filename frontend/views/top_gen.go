package views

import (
	"github.com/nobonobo/spago"
)

// Render ...
func (c *Top) Render() spago.HTML {
	return spago.Tag("body", 
		spago.Tag("button", 			
			spago.Event("click", c.OnClick),
			spago.T(`Hello world`),
		),
	)
}
