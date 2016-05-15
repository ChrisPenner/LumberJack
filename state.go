package main

import "sync"

func initState(state *AppState) {
	state.CurrentMode = NewNormalMode(state)
}

// AppState contains global state
type AppState struct {
	sync.Mutex
	CurrentMode        Mode
	LogViews           LogViews
	Files              []*File
	Categories         Categories
	StatusBar          StatusBar
	SelectCategoryMode SelectCategoryMode
	HandleKeypress     func(string)
	textBuffer         string
}
