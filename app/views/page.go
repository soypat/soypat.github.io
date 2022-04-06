package views

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"soypat.github.io/app/models"
)

type page struct {
	vecty.Core

	Page models.Page `vecty:"prop"`
}

func (p *page) Render() vecty.ComponentOrHTML {
	var cabs vecty.List
	for _, cab := range p.Page.Cabinets {
		cabs = append(cabs, &cabinet{Cabinet: cab})
	}
	return elem.Div(
		elem.Heading2(vecty.Text(p.Page.Title)),
		vecty.Markup(
			vecty.Style("margin-left", "3em"),
			vecty.Attribute("id", p.Page.Title),
		),
		elem.Div(
			vecty.Markup(
				vecty.UnsafeHTML(p.Page.MainContentHTML),
			),
		),
		elem.UnorderedList(
			cabs,
		),
	)
}
