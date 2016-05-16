package main

import "sort"

// AppState contains global state
type AppState struct {
	CurrentMode          string
	LogViews             LogViews
	Files                map[string]File
	Categories           Categories
	StatusBar            StatusBar
	HandleKeypress       func(string)
	selectCategoryBuffer TextBuffer
	selected             int
}

// NewAppState constructs and appstate
func NewAppState(fileNames []string) AppState {
	sort.Strings(fileNames)
	files := make(map[string]File)
	state := AppState{CurrentMode: normalMode, Files: files}

	for _, fileName := range fileNames {
		newFile := File{Name: fileName}
		state.Files[fileName] = newFile
	}

	if len(fileNames) < 3 {
		state.LogViews = fileNames[:]
	} else {
		state.LogViews = fileNames[:2]
	}

	state.Categories = Categories{Items: fileNames}
	return state
}
