package main

import "testing"

func TestInitFiles(t *testing.T) {
	fileNames := []string{"One", "Two"}
	state := NewAppState(fileNames, 10, 10)
	_, hasFile1 := state.Files["One"]
	_, hasFile2 := state.Files["Two"]
	if !hasFile1 || !hasFile2 {
		t.Fail()
	}
}

func TestAppendLine(t *testing.T) {
	fileNames := []string{"One", "Two"}
	state := NewAppState(fileNames, 10, 10)
	state.Files = map[string]file{"One": file{}}
	newState := AppendLine{FileName: "One", Line: "MyLine"}.Apply(state)
	file := newState.Files["One"]
	if file.lines[0] != "MyLine" {
		t.Fail()
	}
}

func TestHlAndFilteredSelectorUpdates(t *testing.T) {
	state := fixtureState()
	f := file{
		lines:                lines{"one", "two", "three"},
		filteredFileSelector: &filteredFileSelector{lastLen: 2, lastHlAndFiltered: lines{"one", "two"}},
	}
	l := f.hlAndFiltered(state)
	if len(l) != 3 || l[0] != "one" && l[1] != "one" && l[2] != "one" {
		t.Fail()
	}
}

// func TestAddWatchers(t *testing.T) {
// 	fileNames := []string{"One", "Two"}
// 	actions := make(chan Action, 100)
// 	addWatchers(fileNames, actions)
// 	if !ok1 || !ok2 || w1.FileName != "One" || w2.FileName != "Two" {
// 		t.Fail()
// 	}
// }
