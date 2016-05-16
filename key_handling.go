package main

// KeyPress sends a keypress
type KeyPress struct {
	Key string
}

// Apply the KeyPress
func (action KeyPress) Apply(state AppState, actions chan<- Action) AppState {
	key := action.Key
	switch state.CurrentMode {
	case normalMode:
		switch key {
		case "<enter>":
			actions <- ChangeMode{Mode: selectCategoryMode}
		case "<backspace>":
			// Actually c-h
			actions <- ChangeSelection{Direction: left}
		case "C-l":
			actions <- ChangeSelection{Direction: right}
		default:
			state.StatusBar.Text = key
		}
	case selectCategoryMode:
		switch key {
		case "C-8":
			actions <- Backspace{}
		case "<enter>":
			actions <- ChangeMode{Mode: normalMode}
		default:
			actions <- TypeKey{Key: convertKey(key)}
		}
	default:
		panic(state.CurrentMode)
	}
	return state
}
