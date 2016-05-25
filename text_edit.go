package main

func (state AppState) typeKey(key string) AppState {
	switch state.CurrentMode {
	case selectCategory:
		state.selectCategoryBuffer = state.selectCategoryBuffer.typeKey(key)
	case search:
		state.searchBuffer = state.searchBuffer.typeKey(key)
		view := state.getSelectedView()
		state.LogViews[state.selected] = view.scrollToSearch(state)
	case editModifier:
		newBuffer := state.modifiers[state.selectedMod].typeKey(key)
		state.modifiers[state.selectedMod].buffer = newBuffer
	}
	return state
}

// buffer provides an abstraction over editing text
type buffer struct {
	text string
}

func (t buffer) typeKey(key string) buffer {
	key = convertKey(key)
	switch key {
	case "<BS>":
		// Backspace
		if len(t.text) > 0 {
			t.text = t.text[:len(t.text)-1]
		}
	default:
		t.text = t.text + key
	}
	return t
}

func convertKey(key string) string {
	switch key {
	case "<space>":
		return " "
	case "C-8":
		return "<BS>"
	default:
		// Just ignore weird control sequences
		if len(key) > 1 {
			return ""
		}
		return key
	}
}
