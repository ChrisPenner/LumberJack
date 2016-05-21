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
	filters              filters
	showFilters          bool
	layout               int
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
	for i := 0; i < 4; i++ {
		viewNames = append(viewNames, fileNames[i%len(fileNames)])
	}

	var views []LogView
	for _, fileName := range viewNames {
		views = append(views, LogView{FileName: fileName})
	}
	state.LogViews = views

	state.Categories = fileNames

	state.filters = filters{
		filter{pattern: "INFO"},
		filter{pattern: "WARNING"},
	}

	state.layout = 1
	return state
}
