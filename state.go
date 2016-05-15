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
func NewAppState() *AppState {
	state := new(AppState)
	state.CurrentMode = normalMode
	state.Files = make(map[string]File)
	return state
}
