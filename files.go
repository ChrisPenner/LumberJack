package main

import "os"
import tail "github.com/hpcloud/tail"

// File contains the lines of a given file
type File struct {
	Name   string
	Lines  []string
	Active bool
}

// InitFiles sets up the status bar
type InitFiles struct {
}

// Apply the InitFiles
func (action InitFiles) Apply(state *AppState) {
	for _, fileName := range os.Args[1:] {
		newFile := &File{Name: fileName}
		state.Files = append(state.Files, newFile)
		state.LogViews.Files = state.Files
		go addTail(fileName, func(newLine string) {
			store.Actions <- AppendLine{File: newFile, Line: newLine}
		})
	}
}

// AppendLine to file
type AppendLine struct {
	File *File
	Line string
}

// Apply the AppendLine
func (action AppendLine) Apply(state *AppState) {
	action.File.Lines = append(action.File.Lines, action.Line)
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
