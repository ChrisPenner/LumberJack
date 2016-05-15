package main

import ui "github.com/gizak/termui"
import "strings"

const statusBarHeight = 1
const categoriesHeight = 1
const numColumns = 12

var store *Store

func logViewHeight() int {
	return ui.TermHeight() - categoriesHeight - statusBarHeight
}

// Render the application as a function of state
func Render(state *AppState) {
	ui.Body.Rows = []*ui.Row{
		state.Categories.Display(),
		state.LogViews.Display(logViewHeight()),
		state.StatusBar.Display(),
	}
	ui.Body.Width = ui.TermWidth()
	ui.Body.Align()
	ui.Render(ui.Body)
	state.CurrentMode.Render(state)
}

func initUI() {
	err := ui.Init()
	if err != nil {
		panic(err)
	}
}

func displayCategories(categories []string) *ui.Par {
	categoriesPar := ui.NewPar("Categories: " + strings.Join(categories[:], ", "))
	categoriesPar.Border = false
	categoriesPar.Height = 1
	return categoriesPar
}

func main() {
	initUI()
	defer ui.Close()

	state := new(AppState)

	store = NewStore()
	go store.ReduceLoop(state)
	store.Actions <- InitState{}
	store.Actions <- InitFiles{}
	store.Actions <- InitCategories{}
	store.Actions <- InitStatusBar{}

	ui.Handle("/sys/kbd/C-c", func(ui.Event) {
		ui.StopLoop()
	})
	ui.Handle("/sys/kbd", func(e ui.Event) {
		key := e.Data.(ui.EvtKbd).KeyStr
		store.Actions <- KeyPress{Key: key}
	})
	ui.Handle("/sys/wnd/resize", func(ui.Event) {
		// Force rerender
		store.Actions <- NullAction{}
	})
	ui.Loop()
}
