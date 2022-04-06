package main

import (
	"encoding/json"
	"strings"
	"syscall/js"

	_ "embed"

	"github.com/hexops/vecty"
	"soypat.github.io/app/store"
	"soypat.github.io/app/store/dispatcher"
	"soypat.github.io/app/views"
)

//go:embed assets/data/whittileaks.json
var whittileaks string

func main() {
	// OnAction must be registered before any storage manipulation.
	dispatcher.Register(store.OnAction)
	err := json.NewDecoder(strings.NewReader(whittileaks)).Decode(&store.Pages)
	if err != nil {
		panic(err)
	}
	// attachItemsStorage()

	body := &views.Body{
		Ctx: store.Ctx,
	}
	store.Listeners.Add(body, func(interface{}) {
		body.Ctx = store.Ctx
		vecty.Rerender(body)
	})
	vecty.RenderBody(body)
}

// attachItemsStorage provides persistent local storage saved on edits so
// no data is lost due to bad connection or refreshed page.
func attachItemsStorage() {
	const key = "vecty_items"
	store.Listeners.Add(nil, func(action interface{}) {
		// if _, ok := action.(*actions.NewItem); !ok {
		// 	// Only save state upon adding an item
		// 	return
		// }
		// After item addition save state.
		b, err := json.Marshal(&store.Pages)
		if err != nil {
			panic(err)
		}
		js.Global().Get("localStorage").Set(key, string(b))
	})

	if data := js.Global().Get("localStorage").Get(key); !data.IsUndefined() {
		// Old session data found, initialize store data.
		err := json.Unmarshal([]byte(data.String()), &store.Pages)
		if err != nil {
			panic(err)
		}
	}
}
