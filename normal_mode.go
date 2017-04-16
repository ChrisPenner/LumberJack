package main

type direction int

const (
	left direction = iota
	right
	up
	down
	bottom
)

func (state *AppState) changeSelection(direction direction) *AppState {
	switch direction {
	case left:
		if state.selected > 0 {
			state.selected--
		}
	case right:
		if state.selected < state.layout-1 {
			state.selected++
		} else if state.showMods {
			state = state.changeMode(modifierMode)
		}
	}
	return state
}
