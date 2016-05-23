package main

import "testing"

func TestSwitchingFocus(t *testing.T) {
	state := NewAppState([]string{"One"}, 10)
	state.CurrentMode = normal
	actions := make(chan Action, 100)

	KeyPress{Key: "<backspace>"}.Apply(state, actions)
	action := <-actions
	changeSelection, ok := action.(ChangeSelection)
	if !ok || changeSelection.Direction != left {
		t.Fail()
	}

	KeyPress{Key: "C-l"}.Apply(state, actions)
	action = <-actions
	changeSelection, ok = action.(ChangeSelection)
	if !ok || changeSelection.Direction != right {
		t.Fail()
	}
}

func TestLayout(t *testing.T) {
	fileNames := []string{"1"}
	state := NewAppState(fileNames, 10)
	actions := make(chan Action, 100)
	state = changeLayout{2}.Apply(state, actions)
	if state.layout != 2 {
		t.Fail()
	}
}
