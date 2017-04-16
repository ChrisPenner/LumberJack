package main

func (state *AppState) findNext(direction direction) *AppState {
	switch direction {
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
