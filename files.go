package main

import "os"
import tail "github.com/hpcloud/tail"

// File contains the lines of a given file
type File struct {
	Name   string
	Lines  []string
	Active bool
}

func initFiles(state *AppState) {
	for _, fileName := range os.Args[1:] {
		newFile := &File{Name: fileName}
		state.Files = append(state.Files, newFile)
		state.LogViews.Files = state.Files
		go addTail(fileName, func(newLine string) {
			state.Lock()
			newFile.Lines = append(newFile.Lines, newLine)
			renderFlag = true
			state.Unlock()
		})
	}
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
