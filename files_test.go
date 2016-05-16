package main

import "testing"

func TestInitFiles(t *testing.T) {
	fileNames := []string{"One", "Two"}
	state := NewAppState(fileNames)
	_, hasFile1 := state.Files["One"]
	_, hasFile2 := state.Files["Two"]
	if !hasFile1 || !hasFile2 {
		t.Fail()
	}
}

func TestAppendLine(t *testing.T) {
	fileNames := []string{"One", "Two"}
	state := NewAppState(fileNames)
	store := NewStore()
	state.Files = map[string]File{"One": File{Name: "One"}}
	newState := AppendLine{FileName: "One", Line: "MyLine"}.Apply(state, store.Actions)
	file := newState.Files["One"]
	if file.Lines[0] != "MyLine" {
		t.Fail()
	}
}

func TestAddWatchers(t *testing.T) {
	fileNames := []string{"One", "Two"}
	store := NewStore()
	addWatchers(fileNames, store.Actions)
	a1 := <-store.Actions
	a2 := <-store.Actions
	w1, ok1 := a1.(WatchFile)
	w2, ok2 := a2.(WatchFile)
	if !ok1 || !ok2 || w1.FileName != "One" || w2.FileName != "Two" {
		t.Fail()
	}
}
