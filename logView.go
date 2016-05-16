package main

import ui "github.com/gizak/termui"
import "os"

type initLogViews struct{}

func (action initLogViews) Apply(state AppState, actions chan<- Action) AppState {
	var fileNames []string
	for _, name := range os.Args[1:] {
		fileNames = append(fileNames, name)
		if len(fileNames) == 2 {
			break
		}
	}
	state.LogViews.viewNames = fileNames
	return state
}

// LogViews is a list of Files
type LogViews struct {
	viewNames []string
}

// Display returns a Row object representing all of the logViews
func (lv LogViews) Display(files map[string]File, height int) *ui.Row {
	listBlocks := []*ui.List{}
	for _, name := range lv.viewNames {
		file := files[name]
		logView := file.Display(height)
		logView.BorderLeft = false
		listBlocks = append(listBlocks, logView)
	}
	if len(listBlocks) > 0 {
		listBlocks[0].BorderLeft = true
	}

	logViewColumns := []*ui.Row{}
	numColumnsEach := 6 //numColumns / 1 //len(state.logViews)
	for _, logViewBlock := range listBlocks {
		logViewColumns = append(logViewColumns, ui.NewCol(numColumnsEach, 0, logViewBlock))
	}
	return ui.NewRow(logViewColumns...)
}

// Select the File at index i
// func (lv LogViews) Select(i int) {
// 	if len(lv.Files) <= i || i < 0 {
// 		return
// 	}
// 	for _, file := range lv.Files {
// 		file.Active = false
// 	}
// 	lv.Files[i].Active = true
// }
