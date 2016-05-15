package main

// AppState contains global state
type AppState struct {
	CurrentMode          string
	LogViews             LogViews
	Files                map[string]File
	Categories           Categories
	StatusBar            StatusBar
	HandleKeypress       func(string)
	selectCategoryBuffer TextBuffer
}

// NewAppState constructs and appstate
func NewAppState() AppState {
	files := make(map[string]File)
	state := AppState{CurrentMode: normalMode, Files: files}
	return state
}
