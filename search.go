package main

type findNext struct {
	direction direction
}

func (action findNext) Apply(state AppState, actions chan<- Action) AppState {
	switch action.direction {
	case up:
		state.searchIndex++
	case down:
		state.searchIndex--
		if state.searchIndex < 0 {
			state.searchIndex = 0
		}
	}
	state.LogViews[state.selected] = state.getSelectedView().scrollToSearch(state)
	return state
}
