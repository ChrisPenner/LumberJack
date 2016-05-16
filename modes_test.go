package main

import "testing"

func TestChangeMode(t *testing.T) {
	state := NewAppState([]string{})
	state.CurrentMode = normalMode
	store := NewStore()
	newState := ChangeMode{Mode: selectCategoryMode}.Apply(state, store.Actions)
	if newState.CurrentMode != selectCategoryMode {
		t.Fail()
	}
}
