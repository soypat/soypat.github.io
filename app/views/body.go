package views

import (
	"soypat.github.io/app/store"
	"soypat.github.io/app/store/actions"
	"soypat.github.io/app/store/dispatcher"

	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/event"
)

type Body struct {
	vecty.Core
	Ctx actions.Context `vecty:"prop"`
}

func (b *Body) Render() vecty.ComponentOrHTML {
	var mainContent vecty.MarkupOrChild
	switch b.Ctx.Page {
	case actions.ViewLanding:
		mainContent = elem.Div(
			&Landing{
				Pages: store.Pages,
			},
		)

	case actions.ViewPage:
		act := b.Ctx.Action.(*actions.PageSelect)
		mainContent = &page{
			Page: store.Pages[act.PageIdx],
		}

	default:
		panic("unknown Page")
	}
	return elem.Body(
		vecty.If(b.Ctx.Referrer != nil, elem.Div(
			elem.Button(
				vecty.Markup(event.Click(b.backButton)),
				vecty.Text("Back"),
			))),
		elem.Div(
			whittileaksLogo(35),
			vecty.Markup(
				vecty.Style("font-size", "35px"),
			),
			vecty.Text("WhittiLeaks"),
		),
		mainContent,
	)
}

func (b *Body) backButton(*vecty.Event) {
	dispatcher.Dispatch(&actions.Back{})
}

func whittileaksLogo(px int) *vecty.HTML {
	return elem.Image(
		vecty.Markup(
			vecty.Attribute("width", px),
			vecty.Attribute("height", px),
			vecty.Attribute("src", "assets/img/whittileaks.png"),
		),
	)
}
