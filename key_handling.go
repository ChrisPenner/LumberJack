package main

// KeyPress sends a keypress
type KeyPress struct {
	Key string
}

// Apply the KeyPress
func (action KeyPress) Apply(state AppState, actions chan<- Action) AppState {
	key := action.Key
	switch state.CurrentMode {
	case normal:
		switch key {
		case "<enter>":
			actions <- ChangeMode{Mode: selectCategory}
		case "?", "/":
			actions <- ChangeMode{Mode: search}
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
	case selectCategory:
		switch key {
		case "<enter>":
			bestMatch, ok := state.Categories.getBestMatch(state)
			if ok {
				actions <- SelectCategory{FileName: bestMatch}
			}
			fallthrough
		case "<escape>":
			actions <- ChangeMode{Mode: normal}
			state.selectCategoryBuffer.text = ""

		default:
			actions <- typeKey{key: convertKey(key)}
		}
	case search:
		switch key {
		case "<enter>":
			// quit search here...
			fallthrough
		case "<escape>":
			actions <- ChangeMode{Mode: normal}
		default:
			actions <- typeKey{key: convertKey(key)}
		}
	default:
		panic("Didn't handle keypress!")
	}
	return state
}
