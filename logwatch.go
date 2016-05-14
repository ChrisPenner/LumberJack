package main

import ui "github.com/gizak/termui"
import tail "github.com/hpcloud/tail"
import "strings"

const statusBarHeight = 1
const categoriesHeight = 1
const numColumns = 12

type appState struct {
	logViews   [][]string
	categories []string
	statusBar  *ui.Par
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

func initialize() {
	err := ui.Init()
	if err != nil {
		panic(err)
	}
}

func logViewHeight() int {
	return ui.TermHeight() - categoriesHeight - statusBarHeight - 1
}

func initializeLogViews(state *appState) {
	logView1 := []string{"Line 1", "Line 2"}
	logView2 := []string{"Other Line 0", "Other line 2"}
	state.logViews = [][]string{logView1, logView2}
}

func initializeCategories(state *appState) {
	state.categories = []string{"Category 1", "Category 2"}
}

func initializeStatusBar(state *appState) {
	statusBar := ui.NewPar("StatusBar!!")
	statusBar.Border = false
	state.statusBar = statusBar
}

func displayCategories(categories []string) *ui.Par {
	categoriesPar := ui.NewPar("Categories: " + strings.Join(categories[:], ", "))
	categoriesPar.Border = false
	categoriesPar.Height = 1
	return categoriesPar
}

func displayLogViews(logViews [][]string) []*ui.List {
	listBlocks := []*ui.List{}
	for _, lines := range logViews {
		list := ui.NewList()
		list.Items = lines
		list.Height = logViewHeight()
		listBlocks = append(listBlocks, list)
	}
	return listBlocks
}

func initializeBody(state *appState) {
	logViewColumns := []*ui.Row{}
	numColumnsEach := 6 //numColumns / 1 //len(state.logViews)
	listObjects := displayLogViews(state.logViews)
	for _, logViewList := range listObjects {
		logViewColumns = append(logViewColumns, ui.NewCol(numColumnsEach, 0, logViewList))
	}

	categoryPar := displayCategories(state.categories)

	ui.Body.AddRows(
		ui.NewRow(ui.NewCol(12, 0, categoryPar)),
		ui.NewRow(logViewColumns...),
		ui.NewRow(ui.NewCol(12, 0, state.statusBar)),
	)
}

func main() {
	initialize()
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
