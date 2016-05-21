package main

import "testing"

func TestChangeSelection(t *testing.T) {
	state := NewAppState([]string{"1", "2"}, 10)
	state.layout = 2
	actions := make(chan Action, 100)

	state.selected = 1
	newState := ChangeSelection{Direction: left}.Apply(state, actions)
	if newState.selected != 0 {
		t.Error("Left from right-side")
	}

	state.selected = 0
	newState = ChangeSelection{Direction: left}.Apply(state, actions)
	if newState.selected != 0 {
		t.Error("Left from left-side")
	}

	state.selected = 0
	newState = ChangeSelection{Direction: right}.Apply(state, actions)
	if newState.selected != 1 {
		t.Error("Right from left-side")
	}

	state.selected = 1
	newState = ChangeSelection{Direction: right}.Apply(state, actions)
	if newState.selected != 1 {
		t.Error("Right from right-side")
	}
}
