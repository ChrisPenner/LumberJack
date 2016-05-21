package main

import "sort"

// AppState contains global state
type AppState struct {
	termHeight           int
	CurrentMode          mode
	LogViews             LogViews
	Files                Files
	Categories           Categories
	StatusBar            StatusBar
	HandleKeypress       func(string)
	selected             int
	searchIndex          int
	selectCategoryBuffer textBuffer
	searchBuffer         textBuffer
	wrap                 bool
}

// NewAppState constructs and appstate
func NewAppState(fileNames []string, height int) AppState {
	sort.Strings(fileNames)
	files := make(map[string]File)
	state := AppState{
		Files:      files,
		termHeight: height,
	}

	for _, fileName := range fileNames {
		state.Files[fileName] = File{}
	}

	viewNames := fileNames[:]
	if len(fileNames) >= 3 {
		viewNames = viewNames[:2]
	}

	var views []LogView
	for _, fileName := range viewNames {
		views = append(views, LogView{FileName: fileName})
	}
	state.LogViews = views

	state.Categories = fileNames
	return state
}
