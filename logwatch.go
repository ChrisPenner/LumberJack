package main

import ui "github.com/gizak/termui"
import tail "github.com/hpcloud/tail"
import "strings"
import "os"
import "time"

const statusBarHeight = 1
const categoriesHeight = 1
const numColumns = 12

var renderFlag = false
var updateChan = make(chan func(*appState))

type appState struct {
	LogViews   LogViews
	Categories Categories
	StatusBar  StatusBar
}

func logViewHeight() int {
	return ui.TermHeight() - categoriesHeight - statusBarHeight
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
			updateChan <- func(state *appState) {
				newFile.Lines = append(newFile.Lines, newLine)
				renderFlag = true
			}
		})
	}
}

func updateRenderLoop(state *appState, update chan func(*appState)) {
	for {
		// Confusing way to guarantee thread safety on state edits
		debounce := time.After(50 * time.Millisecond)
		for {
			select {
			case <-debounce:
				break
			case f := <-update:
				f(state)
				continue
			}
			break
		}
		if !renderFlag {
			continue
		}
		ui.Body.Rows = []*ui.Row{
			state.Categories.Display(),
			state.LogViews.Display(logViewHeight()),
			state.StatusBar.Display(),
		}
		ui.Body.Width = ui.TermWidth()
		ui.Body.Align()
		ui.Render(ui.Body)
		renderFlag = false
	}
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
	renderFlag = true
	// ui.Handle("/sys/mouse/click", func(e ui.Event) {
	// 	state.StatusBar.Text = e.Path
	// 	renderFlag = true
	// 	// keyPress := e.Data.(ui.EvtKbd).KeyStr
	// 	// status.Text = keyPress
	// 	// renderFlag = true
	// })
	ui.Handle("/sys/kbd/C-c", func(ui.Event) {
		ui.StopLoop()
	})
	// ui.Handle("/sys/kbd/", func(e ui.Event) {
	// 	keyPress := e.Data.(ui.EvtKbd).KeyStr
	// 	state.StatusBar.Text = keyPress
	// 	renderFlag = true
	// })
	// ui.Handle("/sys/mouse", func(e ui.Event) {
	// 	state.StatusBar.Text = "Click!"
	// 	renderFlag = true
	// })
	// ui.Handle("/sys/wnd/resize", func(ui.Event) {
	// 	state.StatusBar.Text = "Resize!"
	// 	renderFlag = true
	// })
	go updateRenderLoop(state, updateChan)
	ui.Loop()
}
