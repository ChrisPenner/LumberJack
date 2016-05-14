package main

import ui "github.com/gizak/termui"
import tail "github.com/hpcloud/tail"
import "strings"

const statusBarHeight = 1
const categoriesHeight = 1
const numColumns = 12

type appState struct {
	LogViews   LogViews
	Categories Categories
	StatusBar  StatusBar
}

func addTail(fileName string, callback func(string)) {
	t, err := tail.TailFile(fileName, tail.Config{
		Follow: true,
		Logger: tail.DiscardingLogger,
		// MustExist: true,
	})
	if err != nil {
		panic(err)
	}
	for line := range t.Lines {
		callback(line.Text)
	}
}

func render(state *appState) {
	ui.Body.Align()
	ui.Render(ui.Body)
}

func initializeUI() {
	err := ui.Init()
	if err != nil {
		panic(err)
	}
}

func logViewHeight() int {
	return ui.TermHeight() - categoriesHeight - statusBarHeight - 1
}

func initializeLogViews(state *appState) {
	file1 := File{Lines: []string{"Line 1", "Line 2"}}
	file2 := File{Lines: []string{"Other 1", "Other 2"}}
	state.LogViews = LogViews{Files: []File{file1, file2}}
}

func initializeCategories(state *appState) {
	state.Categories = Categories{Items: []string{"Category 1", "Category 2"}}
}

func initializeStatusBar(state *appState) {
	statusBar := StatusBar{Text: "StatusBar!!"}
	state.StatusBar = statusBar
}

func displayCategories(categories []string) *ui.Par {
	categoriesPar := ui.NewPar("Categories: " + strings.Join(categories[:], ", "))
	categoriesPar.Border = false
	categoriesPar.Height = 1
	return categoriesPar
}

func initializeBody(state *appState) {
	ui.Body.AddRows(
		state.Categories.Display(),
		state.LogViews.Display(logViewHeight()),
		state.StatusBar.Display(),
	)
}

func main() {
	initializeUI()
	defer ui.Close()

	state := new(appState)
	initializeCategories(state)
	initializeStatusBar(state)
	initializeLogViews(state)
	initializeBody(state)
	render(state)
	ui.Handle("/sys/kbd/C-c", func(ui.Event) {
		ui.StopLoop()
	})
	ui.Handle("/sys/kbd/", func(e ui.Event) {
		// keyPress := e.Data.(ui.EvtKbd).KeyStr
		// status.Text = keyPress
		// render()
	})
	// ui.Handle("/sys/wnd/resize", func(e ui.Event) {
	// 	// newHeight := e.Data.(ui.EvtWnd).Height
	// 	newWidth := e.Data.(ui.EvtWnd).Width
	// 	// newHeight := e.Data.(ui.EvtWnd).Height
	// 	status.Text = "Resized"
	// 	ui.Body.Width = newWidth
	// 	ui.Body.Align()
	// 	// ui.Render(ui.Body)
	// 	rerender(state)
	// })
	ui.Loop()
}
