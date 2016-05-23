package main

import "testing"

func TestChangeMode(t *testing.T) {
	state := NewAppState([]string{"One"}, 10, 10)
	state.CurrentMode = normal
	newState := state.changeMode(selectCategory)
	if newState.CurrentMode != selectCategory {
		t.Fail()
	}
}
