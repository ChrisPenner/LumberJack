package main

type changeLayout struct {
	choice int
}

func (action changeLayout) Apply(state AppState, actions chan<- Action) AppState {
	state.layout = action.choice
	if state.selected >= state.layout {
		state.selected = state.layout - 1
	}
	return state
}
