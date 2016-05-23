package main

import "testing"

func TestChangeSelection(t *testing.T) {
	state := NewAppState([]string{"1", "2"}, 10)
	state.layout = 2

	state.selected = 1
	newState := state.changeSelection(left)
	if newState.selected != 0 {
		t.Error("Left from right-side")
	}

	state.selected = 0
	newState = state.changeSelection(left)
	if newState.selected != 0 {
		t.Error("Left from left-side")
	}

	state.selected = 0
	newState = state.changeSelection(right)
	if newState.selected != 1 {
		t.Error("Right from left-side")
	}

	state.selected = 1
	newState = state.changeSelection(right)
	if newState.selected != 1 {
		t.Error("Right from right-side")
	}
}
