package main

func (state AppState) changeLayout(choice int) AppState {
	state.layout = choice
	if state.selected >= state.layout {
		state.selected = state.layout - 1
	}
	state = state.clearUnreadCounts()
	return state
}
