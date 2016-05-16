package main

const normalMode = "normalMode"
const left = "left"
const right = "right"

// ChangeSelection Action
type ChangeSelection struct {
	Direction string
}

// Apply ChangeSelection
func (action ChangeSelection) Apply(state AppState, actions chan<- Action) AppState {
	switch action.Direction {
	case left:
		if state.selected > 0 {
			state.selected--
		}
	case right:
		if state.selected < len(state.LogViews)-1 {
			state.selected++
		}
	}
	return state
}
