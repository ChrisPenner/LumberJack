package main

// Backspace action
type Backspace struct {
}

// Apply the Backspace
func (action Backspace) Apply(state AppState, actions chan<- Action) AppState {
	switch state.CurrentMode {
	case selectCategoryMode:
		text := state.selectCategoryBuffer.Text
		if len(text) > 0 {
			text = text[:len(text)-1]
		}
		state.selectCategoryBuffer.Text = text
	default:
		break
	}
	return state
}

// TypeKey types a key
type TypeKey struct {
	Key string
}

// Apply the Keystroke
func (action TypeKey) Apply(state AppState, actions chan<- Action) AppState {
	switch state.CurrentMode {
	case selectCategoryMode:
		text := state.selectCategoryBuffer.Text
		text = text + action.Key
		state.selectCategoryBuffer.Text = text

	default:
		break
	}
	return state
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
		// Just ignore weird control sequences
		if len(key) > 1 {
			return ""
		}
		return key
	}
}
