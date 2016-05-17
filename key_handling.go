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
		case "<up>", "k":
			actions <- Scroll{Direction: up, NumLines: 1}
		case "<down>", "j":
			actions <- Scroll{Direction: down, NumLines: 1}
		case "b", "C-u":
			actions <- Scroll{Direction: up, NumLines: state.termHeight / 2}
		case "f", "C-d":
			actions <- Scroll{Direction: down, NumLines: state.termHeight / 2}
		default:
			state.StatusBar.Text = key
		}
	case selectCategoryMode:
		switch key {
		case "C-8":
			actions <- Backspace{}
		case "<enter>":
			bestMatch, ok := state.Categories.getBestMatch(state)
			if ok {
				actions <- SelectCategory{FileName: bestMatch}
			}
			fallthrough
		case "<escape>":
			actions <- ChangeMode{Mode: normalMode}
			state.selectCategoryBuffer.Text = ""

		default:
			actions <- TypeKey{Key: convertKey(key)}
		}
	default:
		panic(state.CurrentMode)
	}
	return state
}
