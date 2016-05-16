package main

import ui "github.com/gizak/termui"

// LogViews is a list of viewnames
type LogViews []string
type initLogViews struct{}

// Display returns a Row object representing all of the logViews
func (viewNames LogViews) Display(state AppState) *ui.Row {
	listBlocks := []*ui.List{}
	files := state.Files
	height := logViewHeight()
	for _, name := range viewNames {
		file := files[name]
		logView := file.Display(height)
		logView.BorderLeft = false
		logView.BorderFg = ui.ColorWhite
		listBlocks = append(listBlocks, logView)
	}
	if len(listBlocks) > 0 {
		listBlocks[0].BorderLeft = true
		listBlocks[state.selected].BorderFg = ui.ColorMagenta
	}
	logViewColumns := []*ui.Row{}
	numColumnsEach := 6 //numColumns / 1 //len(state.logViews)
	for _, logViewBlock := range listBlocks {
		logViewColumns = append(logViewColumns, ui.NewCol(numColumnsEach, 0, logViewBlock))
	}
	return ui.NewRow(logViewColumns...)
}
