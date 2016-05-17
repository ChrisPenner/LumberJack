package main

import tail "github.com/hpcloud/tail"

// Files list
type Files map[string]File

// File contains the lines of a given file
type File []string

func (state AppState) getSelectedFileName() string {
	return state.LogViews[state.selected].FileName
}

func (state AppState) getSelectedView() LogView {
	return state.LogViews[state.selected]
}

func (state AppState) getSelectedFile() File {
	return state.Files[state.getSelectedFileName()]
}

func (state AppState) getFile(fileName string) File {
	return state.Files[fileName]
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
	file = append(file, action.Line)
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
