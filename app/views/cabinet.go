package views

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"soypat.github.io/app/models"
)

type cabinet struct {
	vecty.Core

	Cabinet models.Cabinet `vecty:"prop"`
}

func (c *cabinet) Render() vecty.ComponentOrHTML {
	var files vecty.List
	for _, f := range c.Cabinet.Files {
		files = append(files, &file{File: f})
	}
	return elem.Div(
		elem.Heading4(vecty.Text(c.Cabinet.Title)),
		elem.UnorderedList(files),
	)
}
