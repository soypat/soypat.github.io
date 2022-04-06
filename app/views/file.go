package views

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"soypat.github.io/app/models"
)

type file struct {
	vecty.Core

	File models.File `vecty:"prop"`
}

func (c *file) Render() vecty.ComponentOrHTML {
	const dlBaseURL = "https://sites.google.com"
	var links vecty.List
	for i, l := range c.File.Links {
		var linkTitle string
		switch i {
		case 0:
			linkTitle = "Ver"
		case 1:
			linkTitle = "Descargar"
			l = dlBaseURL + l
		default:
			linkTitle = "Link"
		}
		links = append(links, elem.Span(
			elem.Anchor(
				vecty.Text(linkTitle),
				vecty.Markup(
					vecty.Attribute("href", l),
				),
			),
			elem.Span(vecty.Text("\t")),
		),
		)
	}
	return elem.Div(
		vecty.Markup(
			vecty.Style("margin", "1em"),
		),
		elem.Span(elem.Strong(vecty.Text(c.File.Title()+"\t"))),
		elem.Span(vecty.Text(c.File.Description())),
		elem.Span(links),
		elem.Span(vecty.Text("\t\t")),
	)
}