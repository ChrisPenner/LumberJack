package main

import ui "github.com/gizak/termui"

const statusBarHeight = 1
const categoriesHeight = 1
const numColumns = 12

var store *Store

func logViewHeight() int {
	return ui.TermHeight() - categoriesHeight - statusBarHeight
}

// Render the application as a function of state
func Render(state AppState) {
	ui.Body.Rows = []*ui.Row{
		state.Categories.Display(),
		state.LogViews.Display(state.Files, logViewHeight()),
		state.StatusBar.Display(state.CurrentMode),
	}
	ui.Body.Width = ui.TermWidth()
	ui.Body.Align()
	ui.Render(ui.Body)
	switch state.CurrentMode {
	case selectCategoryMode:
		renderSelectCategoryModal(state)
	}
}

func initUI() {
	err := ui.Init()
	if err != nil {
		panic(err)
	}
}

func main() {
	initUI()
	defer ui.Close()

	state := NewAppState()
	store = NewStore()
	go store.ReduceLoop(state)

	store.Actions <- initFiles{}
	store.Actions <- initLogViews{}
	store.Actions <- initCategories{}

	ui.Handle("/sys/kbd/C-c", func(ui.Event) {
		ui.StopLoop()
	})
	ui.Handle("/sys/kbd", func(e ui.Event) {
		key := e.Data.(ui.EvtKbd).KeyStr
		store.Actions <- KeyPress{Key: key}
	})
	ui.Handle("/sys/wnd/resize", func(ui.Event) {
		// Force rerender
		store.Actions <- render{}
	})
	ui.Loop()
}
