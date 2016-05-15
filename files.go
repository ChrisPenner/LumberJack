package main

import "os"
import tail "github.com/hpcloud/tail"
import ui "github.com/gizak/termui"

// File contains the lines of a given file
type File struct {
	Name   string
	Lines  []string
	Active bool
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

type initFiles struct {
}

func (action initFiles) Apply(state AppState) AppState {
	for _, fileName := range os.Args[1:] {
		newFile := File{Name: fileName}
		state.Files[fileName] = newFile
		go addTail(fileName, func(innerFileName string, newLine string) {
			store.Actions <- AppendLine{FileName: innerFileName, Line: newLine}
		})
	}
	return state
}

// AppendLine to file
type AppendLine struct {
	FileName string
	Line     string
}

// Apply the AppendLine
func (action AppendLine) Apply(state AppState) AppState {
	file := state.Files[action.FileName]
	file.Lines = append(file.Lines, action.Line)
	state.Files[action.FileName] = file
	return state
}

func addTail(fileName string, callback func(string, string)) {
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
