package main

// AppState contains global state
type AppState struct {
	CurrentMode        Mode
	LogViews           LogViews
	Files              []*File
	Categories         Categories
	StatusBar          StatusBar
	SelectCategoryMode SelectCategoryMode
	HandleKeypress     func(string)
	textBuffer         string
}

// InitState sets up the status bar
type InitState struct {
}

// Apply the InitState
func (action InitState) Apply(state *AppState) {
	state.CurrentMode = NewNormalMode()
}
