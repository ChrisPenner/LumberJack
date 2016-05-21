package main

import "testing"

func TestChangeMode(t *testing.T) {
	state := NewAppState([]string{"One"}, 10)
	state.CurrentMode = normal
	actions := make(chan Action, 100)
	newState := ChangeMode{Mode: selectCategory}.Apply(state, actions)
	if newState.CurrentMode != selectCategory {
		t.Fail()
	}
}
