package views

import (
	"strings"

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
					vecty.Attribute("href", "#"+strings.ReplaceAll(item.Title, "-", "")),
					event.Click(func(e *vecty.Event) {
						dispatcher.Dispatch(&actions.PageSelect{View: actions.ViewPage, PageIdx: i})
					})),
				vecty.Text(item.LinkTitle),
			),
		))
	}
	return elem.Div(
		elem.Heading1(vecty.Text("Whittileaks 3.9")),
		elem.Paragraph(
			vecty.Text("⚠️⚠️El sitio está pasando por una fase de reestructuración."),
			elem.Anchor(
				vecty.Markup(vecty.Attribute("href", "https://drive.google.com/drive/folders/0BxlGAHNyMIneZVIwTVlmYU45ck0?resourcekey=0-BEt5b_on6_crCYH4E-7DZg&usp=drive_link")),
				vecty.Text("Puede encontrar links a los drives aquí con el contenido de whittileaks."),
			),
		),
		elem.UnorderedList(items),
		elem.Paragraph(vecty.Text("Frontendistas interesados en mejorar -esto- un cacho -> Reach me at pwhittingslow{-at-}itba{dot}edu{dot}ar. Warning: no javascript allowed"), elem.Strong(vecty.Text(", there can be only wasm."))),
	)
}
