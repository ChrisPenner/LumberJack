package main

import "testing"

func fixtureState() AppState {
	state := NewAppState([]string{"One", "Two", "Three"}, 10, 10)
	state.CurrentMode = normal
	return state
}

func TestNewAppStateSetsNormalMode(t *testing.T) {
	actual := NewAppState([]string{"One"}, 10, 10).CurrentMode
	expected := normal
	if actual != expected {
		t.Fail()
	}
}

func TestNewAppStateSetsFilesMap(t *testing.T) {
	m := NewAppState([]string{"One", "Two"}, 10, 10).Files
	if len(m) != 2 {
		t.Fail()
	}
}

func TestNewAppStateSetsCategories(t *testing.T) {
	fileNames := []string{"one", "two"}
	state := NewAppState(fileNames, 10, 10)
	_, hasKey1 := state.Files["one"]
	_, hasKey2 := state.Files["two"]
	if len(state.Files) != 2 || !hasKey1 || !hasKey2 {
		t.Fail()
	}
}

func TestNewAppStateSetsViewFiles(t *testing.T) {
	fileNames := []string{"One", "Two"}
	state := NewAppState(fileNames, 10, 10)
	views := state.LogViews
	if views[0].FileName != "One" || views[1].FileName != "Two" || views[2].FileName != "One" || views[3].FileName != "Two" {
		t.Fail()
	}
}
