package views

import (
	"github.com/hexops/vecty"
	"github.com/hexops/vecty/elem"
	"github.com/hexops/vecty/event"
	"soypat.github.io/app/models"
	"soypat.github.io/app/store/actions"
	"soypat.github.io/app/store/dispatcher"
)

type Landing struct {
	vecty.Core

	Pages []models.Page `vecty:"prop"`
}

func (l *Landing) Render() vecty.ComponentOrHTML {
	var items vecty.List
	for i, item := range l.Pages {
		i := i // escape loop var
		items = append(items, elem.ListItem(
			elem.Anchor(
				vecty.Markup(
					vecty.Attribute("href", "#"),
					event.Click(func(e *vecty.Event) {
						dispatcher.Dispatch(&actions.PageSelect{View: actions.ViewPage, PageIdx: i})
					})),
				vecty.Text(item.LinkTitle),
			),
		))
	}
	return elem.Div(
		elem.Heading1(vecty.Text("Whittileaks 4.0")),
		elem.Heading3(vecty.Text("La universidad 4.0 llega a whittileaks con todos sus beneficios.")),
		elem.Paragraph(vecty.Text("Dirigir quejas "),
			elem.Anchor(
				vecty.Markup(vecty.Attribute("href", "#salu3")),
				vecty.Text("aquí."),
			)),
		elem.UnorderedList(items),
		elem.Paragraph(vecty.Text("Frontendistas interesados en mejorar -esto- un cacho -> Reach me at pwhittingslow{-at-}itba{dot}edu{dot}ar. Warning: no javascript allowed"), elem.Strong(vecty.Text(", there can be only wasm."))),
	)
}
