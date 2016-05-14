package main

import ui "github.com/gizak/termui"

// File contains the lines of a given file
type File struct {
	Name  string
	Lines []string
}

// Display returns a list object representing the file
func (f File) Display(height int) *ui.List {
	list := ui.NewList()
	list.Height = height
	list.Items = f.Lines
	return list
}

// LogViews is a list of Files
type LogViews struct {
	Files []*File
}

// Display returns a Row object representing all of the logViews
func (lv LogViews) Display(height int) *ui.Row {
	listBlocks := []*ui.List{}
	for _, file := range lv.Files {
		listBlocks = append(listBlocks, file.Display(height))
	}

	logViewColumns := []*ui.Row{}
	numColumnsEach := 6 //numColumns / 1 //len(state.logViews)
	for _, logViewBlock := range listBlocks {
		logViewColumns = append(logViewColumns, ui.NewCol(numColumnsEach, 0, logViewBlock))
	}
	return ui.NewRow(logViewColumns...)
}
