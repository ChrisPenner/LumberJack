package main

import "testing"

func TestChangeMode(t *testing.T) {
	state := NewAppState([]string{}, 10)
	state.CurrentMode = normalMode
	actions := make(chan Action, 100)
	newState := ChangeMode{Mode: selectCategoryMode}.Apply(state, actions)
	if newState.CurrentMode != selectCategoryMode {
		t.Fail()
	}
}
