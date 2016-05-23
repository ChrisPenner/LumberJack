package main

import ui "github.com/gizak/termui"
import "os"

const statusBarHeight = 1
const categoriesHeight = 1

func logViewHeight(termHeight int) int {
	return termHeight - categoriesHeight - statusBarHeight
}

type resize struct {
	height int
	width  int
}

func (action resize) Apply(state AppState) AppState {
	state.termHeight = action.height
	state.termWidth = action.width
	return state
}

// Render the application as a function of state
func Render(state AppState) {
	mainColumns := state.LogViews.display(state)
	if state.showFilters {
		filterColumn := state.filters.display(state)
		mainColumns = append(mainColumns, filterColumn)
	}
	mainRow := ui.NewRow(mainColumns...)
	ui.Body.Rows = []*ui.Row{
		state.Categories.Display(),
		mainRow,
		state.StatusBar.display(state),
	}
	ui.Body.Width = ui.TermWidth()
	ui.Body.Align()
	ui.Render(ui.Body)
	if state.CurrentMode == selectCategory {
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
	state := NewAppState(fileNames, ui.TermHeight(), ui.TermWidth())
	addWatchers(fileNames, store.Actions)
	go store.ReduceLoop(state)

	ui.Handle("/sys/kbd/C-c", func(ui.Event) {
		ui.StopLoop()
	})
	ui.Handle("/sys/kbd", func(e ui.Event) {
		key := e.Data.(ui.EvtKbd).KeyStr
		store.Actions <- KeyPress{Key: key}
	})
	ui.Handle("/sys/wnd/resize", func(e ui.Event) {
		wndEvent := e.Data.(ui.EvtWnd)
		// Force rerender
		store.Actions <- resize{height: wndEvent.Height, width: wndEvent.Width}
	})
	ui.Loop()
}
