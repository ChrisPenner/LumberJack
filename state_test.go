package main

import "testing"

func TestNewAppStateSetsNormalMode(t *testing.T) {
	actual := NewAppState([]string{}, 10).CurrentMode
	expected := normalMode
	if actual != expected {
		t.Fail()
	}
}

func TestNewAppStateSetsBlankFilesMap(t *testing.T) {
	m := NewAppState([]string{}, 10).Files
	if len(m) != 0 {
		t.Fail()
	}
}

func TestNewAppStateSetsCategories(t *testing.T) {
	fileNames := []string{"one", "two"}
	state := NewAppState(fileNames, 10)
	_, hasKey1 := state.Files["one"]
	_, hasKey2 := state.Files["two"]
	if len(state.Files) != 2 || !hasKey1 || !hasKey2 {
		t.Fail()
	}
}

func TestNewAppStateSetsOneFile(t *testing.T) {
	fileNames := []string{"One"}
	state := NewAppState(fileNames, 10)
	viewNames := state.LogViews
	if len(viewNames) != 1 || viewNames[0].FileName != "One" {
		t.Fail()
	}
}

func TestNewAppStateTakesFirstTwoFilesAsLogViews(t *testing.T) {
	fileNames := []string{"One", "Two", "Three", "Four"}
	state := NewAppState(fileNames, 10)
	viewNames := state.LogViews
	if len(viewNames) != 2 {
		t.Fail()
	}
}

func TestNewAppStateSetsLogViews(t *testing.T) {
	fileNames := []string{"One", "Two"}
	state := NewAppState(fileNames, 10)
	viewNames := state.LogViews
	if len(viewNames) != len(fileNames) {
		t.Fail()
	}
	if viewNames[0].FileName != "One" || viewNames[1].FileName != "Two" {
		t.Fail()
	}
}
