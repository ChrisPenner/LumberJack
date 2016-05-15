package main

import "testing"

func TestInitFiles(t *testing.T) {
	store = NewStore()
	state := NewAppState()
	state.CommandLineArgs = []string{"One", "Two"}
	newState := initFiles{}.Apply(state)
	_, hasFile1 := newState.Files["One"]
	_, hasFile2 := newState.Files["Two"]
	if !hasFile1 || !hasFile2 {
		t.Fail()
	}
}

func TestAppendLine(t *testing.T) {
	state := NewAppState()
	state.Files = map[string]File{"1": File{Name: "1"}}
	newState := AppendLine{FileName: "1", Line: "MyLine"}.Apply(state)
	file := newState.Files["1"]
	if file.Lines[0] != "MyLine" {
		t.Fail()
	}
}
