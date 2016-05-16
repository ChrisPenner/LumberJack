package main

import ui "github.com/gizak/termui"
import "os"

const statusBarHeight = 1
const categoriesHeight = 1
const numColumns = 12

func logViewHeight() int {
	return ui.TermHeight() - categoriesHeight - statusBarHeight
}

// Render the application as a function of state
func Render(state AppState) {
	ui.Body.Rows = []*ui.Row{
		state.Categories.Display(),
		state.LogViews.Display(state),
		state.StatusBar.Display(),
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

	store := NewStore()
	fileNames := os.Args[1:]
	state := NewAppState(fileNames)
	addWatchers(fileNames, store.Actions)
	go store.ReduceLoop(state)

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
