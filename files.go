package main

import tail "github.com/hpcloud/tail"
import ui "github.com/gizak/termui"

// Files list
type Files map[string]File

// File contains the lines of a given file
type File struct {
	Name   string
	Lines  []string
	Active bool
}

// WatchFile Action
type WatchFile struct {
	FileName string
}

// Apply WatchFile
func (action WatchFile) Apply(state AppState, actions chan<- Action) AppState {
	addNewLine := func(fileName string, newLine string) {
		actions <- AppendLine{FileName: fileName, Line: newLine}
	}
	go tailFile(action.FileName, addNewLine)
	return state
}

// Display returns a list object representing the file
func (f File) Display(height int) *ui.List {
	list := ui.NewList()
	list.Height = height
	if f.Active {
		list.BorderFg = ui.ColorWhite
	} else {
		list.BorderFg = ui.ColorYellow
	}
	list.BorderLabel = f.Name
	sliceStart := len(f.Lines) - (height - 2)
	if sliceStart < 0 {
		sliceStart = 0
	}
	list.Items = f.Lines[sliceStart:]
	return list
}

func addWatchers(fileNames []string, actions chan<- Action) {
	for _, fileName := range fileNames {
		actions <- WatchFile{FileName: fileName}
	}
}

// AppendLine to file
type AppendLine struct {
	FileName string
	Line     string
}

// Apply the AppendLine
func (action AppendLine) Apply(state AppState, actions chan<- Action) AppState {
	file := state.Files[action.FileName]
	file.Lines = append(file.Lines, action.Line)
	state.Files[action.FileName] = file
	return state
}

func tailFile(fileName string, callback func(string, string)) {
	t, err := tail.TailFile(fileName, tail.Config{
		Follow:    true,
		Logger:    tail.DiscardingLogger,
		MustExist: true,
	})
	if err != nil {
		panic(err)
	}
	for line := range t.Lines {
		callback(fileName, line.Text)
	}
}
