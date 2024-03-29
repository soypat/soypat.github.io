package store

import (
	"soypat.github.io/app/models"
	"soypat.github.io/app/store/actions"
	"soypat.github.io/app/store/storeutil"
)

var (
	Ctx   actions.Context
	Pages []models.Page

	Listeners = storeutil.NewListenerRegistry()
)

func OnAction(action interface{}) {
	switch a := action.(type) {
	case *actions.PageSelect:
		oldCtx := Ctx
		Ctx = actions.Context{
			Page:     a.View,
			Referrer: &oldCtx,
			Action:   a,
		}

	case *actions.Back:
		Ctx = *Ctx.Referrer

	default:
		panic("unknown action selected!")
	}

	Listeners.Fire(action)
}
