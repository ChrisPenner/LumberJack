package main

// KeyPress sends a keypress
type KeyPress struct {
	Key string
}

// Apply the KeyPress
func (action KeyPress) Apply(state *AppState) {
	key := action.Key
	switch state.CurrentMode {
	case normalMode:
		switch key {
		case "<enter>":
			store.Actions <- ChangeMode{Mode: selectCategoryMode}
		}
	case selectCategoryMode:
		switch key {
		case "C-8":
			store.Actions <- Backspace{}
		case "<enter>":
			store.Actions <- ChangeMode{Mode: normalMode}
			store.Actions <- Backspace{}
		default:
			store.Actions <- TypeKey{Key: convertKey(key)}
		}
	default:
		panic(state.CurrentMode)
	}
}
