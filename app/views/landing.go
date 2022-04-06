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
	return elem.UnorderedList(items)
}
