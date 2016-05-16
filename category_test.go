package main

import "testing"

func TestInitCategoriesAction(t *testing.T) {
	state := NewAppState()
	store := NewStore()
	state.Files["one"] = File{}
	state.Files["two"] = File{}
	newState := initCategories{}.Apply(state, store.Actions)
	_, hasKey1 := state.Files["one"]
	_, hasKey2 := state.Files["two"]
	if len(newState.Files) != 2 || !hasKey1 || !hasKey2 {
		t.Fail()
	}
}
