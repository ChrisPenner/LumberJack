package main

import "testing"

func TestInitLogViews(t *testing.T) {
	state := NewAppState()
	args := []string{"One", "Two"}
	state.CommandLineArgs = args
	store := NewStore()
	newState := initLogViews{}.Apply(state, store.Actions)
	viewNames := newState.LogViews.viewNames
	if len(viewNames) != len(state.CommandLineArgs) {
		t.Fail()
	}

	if viewNames[0] != "One" || viewNames[1] != "Two" {
		t.Fail()
	}
}

func TestInitLogViewsTakesFirstTwoFiles(t *testing.T) {
	state := NewAppState()
	args := []string{"One", "Two", "Three", "Four"}
	state.CommandLineArgs = args
	store := NewStore()
	newState := initLogViews{}.Apply(state, store.Actions)
	viewNames := newState.LogViews.viewNames
	if len(viewNames) != 2 {
		t.Fail()
	}
}

func TestInitLogViewsWithOneFile(t *testing.T) {
	state := NewAppState()
	args := []string{"One"}
	state.CommandLineArgs = args
	store := NewStore()
	newState := initLogViews{}.Apply(state, store.Actions)
	viewNames := newState.LogViews.viewNames
	if len(viewNames) != 1 || viewNames[0] != "One" {
		t.Fail()
	}
}
