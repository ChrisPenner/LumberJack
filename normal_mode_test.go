package main

import "testing"

func TestChangeSelection(t *testing.T) {
	state := NewAppState([]string{"1", "2"}, 10, 10)
	state.layout = 2

	state.selected = 1
	state = state.changeSelection(left)
	if state.selected != 0 {
		t.Error("Left from right-side")
	}

	state.selected = 0
	state = state.changeSelection(left)
	if state.selected != 0 {
		t.Error("Left from left-side")
	}

	state.selected = 0
	state = state.changeSelection(right)
	if state.selected != 1 {
		t.Error("Right from left-side")
	}

	state.selected = 1
	state = state.changeSelection(right)
	if state.selected != 1 {
		t.Error("Right from right-side")
	}
}

func TestSelectingFilterMode(t *testing.T) {
	state := NewAppState([]string{"1", "2"}, 10, 10)
	state.layout = 2
	state.selected = 1
	state.showFilters = true

	state = state.changeSelection(right)

	if state.selected != 1 || state.CurrentMode != filterMode {
		t.Error(state.CurrentMode)
	}
}

func TestSelectingFilterNotVisible(t *testing.T) {
	state := NewAppState([]string{"1", "2"}, 10, 10)
	state.layout = 2
	state.selected = 1
	state.showFilters = false

	state = state.changeSelection(right)

	if state.selected != 1 || state.CurrentMode != normal {
		t.Error(state.CurrentMode)
	}
}
