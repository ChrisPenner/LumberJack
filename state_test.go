package main

import "testing"

func TestNewAppStateSetsNormalMode(t *testing.T) {
	actual := NewAppState().CurrentMode
	expected := normalMode
	if actual != expected {
		t.Fail()
	}
}

func TestNewAppStateSetsBlankFilesMap(t *testing.T) {
	m := NewAppState().Files
	if len(m) != 0 {
		t.Fail()
	}
}
