package main

type direction int

const (
	left direction = iota
	right
	up
	down
	bottom
)

// ChangeSelection Action
type ChangeSelection struct {
	Direction direction
}

// Apply ChangeSelection
func (action ChangeSelection) Apply(state AppState, actions chan<- Action) AppState {
	switch action.Direction {
	case left:
		if state.selected > 0 {
			state.selected--
		}
	case right:
		if state.selected < state.layout-1 {
			state.selected++
		}
	}
	return state
}
