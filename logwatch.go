package main

import ui "github.com/gizak/termui"
import "strings"
import "time"

const statusBarHeight = 1
const categoriesHeight = 1
const numColumns = 12

var renderFlag = true
var updateChan = make(chan func(*AppState))

func logViewHeight() int {
	return ui.TermHeight() - categoriesHeight - statusBarHeight
}

func render(state *AppState) {
	ui.Body.Rows = []*ui.Row{
		state.Categories.Display(),
		state.LogViews.Display(logViewHeight()),
		state.StatusBar.Display(),
	}
	ui.Body.Width = ui.TermWidth()
	ui.Body.Align()
	ui.Render(ui.Body)
	state.CurrentMode.Render()
}

func renderLoop(state *AppState) {
	for {
		if !renderFlag {
			time.Sleep(50 * time.Millisecond)
			continue
		}
		state.Lock()
		render(state)
		renderFlag = false
		state.Unlock()
	}
}

func initUI() {
	err := ui.Init()
	if err != nil {
		panic(err)
	}
}

func initStatusBar(state *AppState) {
	statusBar := StatusBar{Text: "StatusBar!!"}
	state.StatusBar = statusBar
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
	initState(state)
	initFiles(state)
	initCategories(state)
	initStatusBar(state)
	ui.Handle("/sys/kbd/C-c", func(ui.Event) {
		ui.StopLoop()
	})
	ui.Handle("/sys/kbd/<enter>", func(ui.Event) {
		state.Lock()
		defer state.Unlock()
		state.CurrentMode = state.CurrentMode.Next()(state)
		renderFlag = true
	})
	ui.Handle("/sys/kbd", func(e ui.Event) {
		key := e.Data.(ui.EvtKbd).KeyStr
		state.Lock()
		defer state.Unlock()
		state.CurrentMode.KeyboardHandler(key)
		state.StatusBar.Text = key
		renderFlag = true
	})
	ui.Handle("/sys/wnd/resize", func(ui.Event) {
		state.Lock()
		defer state.Unlock()
		renderFlag = true
	})
	go renderLoop(state)
	ui.Loop()
}
