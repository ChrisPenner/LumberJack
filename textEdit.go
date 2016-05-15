package main

// Backspace action
type Backspace struct {
}

// Apply the Backspace
func (action Backspace) Apply(state *AppState) {
	selectCategoryMode, ok := state.CurrentMode.(SelectCategoryMode)
	if ok {
		text := selectCategoryMode.Buffer.Text
		if len(text) > 0 {
			text = text[:len(text)-1]
		}
		selectCategoryMode.Buffer.Text = text
		state.CurrentMode = selectCategoryMode
	}
}

// TypeKey types a key
type TypeKey struct {
	Key string
}

// Apply the Keystroke
func (action TypeKey) Apply(state *AppState) {
	selectCategoryMode, ok := state.CurrentMode.(SelectCategoryMode)
	if ok {
		text := selectCategoryMode.Buffer.Text
		text = text + action.Key
		selectCategoryMode.Buffer.Text = text
		state.CurrentMode = selectCategoryMode
	}
}

// TextBuffer provides an abstraction over editing text
type TextBuffer struct {
	Cursor int
	Text   string
}

func convertKey(key string) string {
	switch key {
	case "<space>":
		return " "
	default:
		return key
	}
}
