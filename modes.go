package main

type mode int

const (
	normal mode = iota
	selectCategory
	search
	modifierMode
	editModifier
)

func (state AppState) changeMode(mode mode) AppState {
	state.CurrentMode = mode
	switch mode {
	case search:
		sb := state.searchBuffer
		sb.text = ""
		state.searchBuffer = sb
		state.searchIndex = 0
	}
	return state
}
