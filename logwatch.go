package main

import ui "github.com/gizak/termui"
import tail "github.com/hpcloud/tail"
import "strings"
import "os"

const statusBarHeight = 1
const categoriesHeight = 1
const numColumns = 12

type appState struct {
	LogViews   LogViews
	Categories Categories
	StatusBar  StatusBar
}

func logViewHeight() int {
	return ui.TermHeight() - categoriesHeight - statusBarHeight - 1
}

func addTail(fileName string, callback func(string)) {
	t, err := tail.TailFile(fileName, tail.Config{
		Follow:    true,
		Logger:    tail.DiscardingLogger,
		MustExist: true,
	})
	if err != nil {
		panic(err)
	}
	for line := range t.Lines {
		callback(line.Text)
	}
}

func initFiles(state *appState) {
	for _, fileName := range os.Args[1:] {
		newFile := new(File)
		newFile.Name = fileName
		state.LogViews.Files = append(state.LogViews.Files, newFile)
		go addTail(fileName, func(newLine string) {
			newFile.Lines = append(newFile.Lines, newLine)
			render(state)
		})
	}
}

func render(state *appState) {
	ui.Body.Rows = []*ui.Row{
		state.Categories.Display(),
		state.LogViews.Display(logViewHeight()),
		state.StatusBar.Display(),
	}
	ui.Body.Align()
	ui.Render(ui.Body)
}

func initUI() {
	err := ui.Init()
	if err != nil {
		panic(err)
	}
}

func initCategories(state *appState) {
	state.Categories = Categories{Items: []string{"Category 1", "Category 2"}}
}

func initStatusBar(state *appState) {
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

	state := new(appState)
	initFiles(state)
	initCategories(state)
	initStatusBar(state)
	render(state)
	ui.Handle("/sys/kbd/C-c", func(ui.Event) {
		render(state)
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
