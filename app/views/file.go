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
			// Temporary files.
		default:
			linkTitle = "Link"
		}
		if len(c.File.Links) == 2 {
			// Temporary workaround for google site files
			l = "https://drive.google.com/drive/folders/11qxYf4nRtOPVDK0Ymi-2kuqTlKEbSkP6?usp=sharing"
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
	date := c.File.DateAdded()
	return elem.Div(
		vecty.Markup(
			vecty.Style("margin", "1em"),
		),
		elem.Span(elem.Strong(vecty.Text(c.File.Title()+"\t"))),
		elem.Strong(vecty.Text("\t-\t")),
		elem.Span(vecty.Text(c.File.Description())),
		elem.Span(links),
		vecty.If(!date.IsZero(),
			elem.Strong(vecty.Text("\t-\t")),
			elem.Span(vecty.Text("Added "+date.Format("02 Jan 2006, 15:04"))),
		),
	)
}
