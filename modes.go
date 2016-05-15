package main

// ChangeMode changes modes
type ChangeMode struct {
	Mode string
}

// Apply the ChangeMode
func (action ChangeMode) Apply(state AppState) AppState {
	state.CurrentMode = action.Mode
	return state
}
