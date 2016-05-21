package main

type mode int

const (
	normal mode = iota
	selectCategory
	search
)

// ChangeMode changes modes
type ChangeMode struct {
	Mode mode
}

// Apply the ChangeMode
func (action ChangeMode) Apply(state AppState, actions chan<- Action) AppState {
	state.CurrentMode = action.Mode
	switch action.Mode {
	case search:
		sb := state.searchBuffer
		sb.text = ""
		state.searchBuffer = sb
	}
	return state
}
