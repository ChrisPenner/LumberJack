package main

import "sync"
import "strconv"

func initState(state *appState) {
	state.HandleKeypress = state.CommandModeHandleKey
}

type appState struct {
	sync.Mutex
	CurrentModal        Modal
	LogViews            LogViews
	Categories          Categories
	StatusBar           StatusBar
	SelectCategoryModal SelectCategoryModal
	HandleKeypress      func(string)
	textBuffer          string
}

func (state *appState) CommandModeHandleKey(key string) {
	state.Lock()
	defer state.Unlock()
	switch key {
	case "<enter>":
		state.CurrentModal = SelectCategoryModal{}
		state.HandleKeypress = state.typingModeHandleKey
	}
	if n, err := strconv.Atoi(key); err == nil {
		state.StatusBar.Text = key
		state.LogViews.Select(n - 1)
	}
	renderFlag = true
}

func (state *appState) typingModeHandleKey(key string) {
	state.Lock()
	defer state.Unlock()
	switch key {
	case "<enter>":
		state.CurrentModal.Done(state)
	case "C-8":
		// Backspace
		if len(state.textBuffer) > 0 {
			state.textBuffer = state.textBuffer[:len(state.textBuffer)-1]
		}
	case "<space>":
		key = " "
		fallthrough
	default:
		state.textBuffer = state.textBuffer + key
	}
	if n, err := strconv.Atoi(key); err == nil {
		state.StatusBar.Text = key
		state.LogViews.Select(n - 1)
	}
	renderFlag = true
}
