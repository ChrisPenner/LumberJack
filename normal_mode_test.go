package main

import "testing"

func TestChangeSelection(t *testing.T) {
	state := NewAppState([]string{"1", "2"})
	store := NewStore()

	state.selected = 1
	newState := ChangeSelection{Direction: left}.Apply(state, store.Actions)
	if newState.selected != 0 {
		t.Error("Left from right-side")
	}

	state.selected = 0
	newState = ChangeSelection{Direction: left}.Apply(state, store.Actions)
	if newState.selected != 0 {
		t.Error("Left from left-side")
	}

	state.selected = 0
	newState = ChangeSelection{Direction: right}.Apply(state, store.Actions)
	if newState.selected != 1 {
		t.Error("Right from left-side")
	}

	state.selected = 1
	newState = ChangeSelection{Direction: right}.Apply(state, store.Actions)
	if newState.selected != 1 {
		t.Error("Right from right-side")
	}
}
