package main

// AppState contains global state
type AppState struct {
	CurrentMode        Mode
	LogViews           LogViews
	Files              map[string]File
	Categories         Categories
	StatusBar          StatusBar
	SelectCategoryMode SelectCategoryMode
	HandleKeypress     func(string)
}

// NewAppState constructs and appstate
func NewAppState() *AppState {
	state := new(AppState)
	state.CurrentMode = NewNormalMode()
	state.Files = make(map[string]File)
	return state
}
