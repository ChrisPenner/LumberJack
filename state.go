package main

import "os"

// AppState contains global state
type AppState struct {
	CurrentMode          string
	CommandLineArgs      []string
	LogViews             LogViews
	Files                map[string]File
	Categories           Categories
	StatusBar            StatusBar
	HandleKeypress       func(string)
	selectCategoryBuffer TextBuffer
	selected             int
}

// NewAppState constructs and appstate
func NewAppState() AppState {
	files := make(map[string]File)
	state := AppState{CurrentMode: normalMode, Files: files}
	state.CommandLineArgs = os.Args[1:]
	return state
}
