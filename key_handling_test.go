package main

import "testing"

func TestSwitchingFocusTooFarLeft(t *testing.T) {
	state := NewAppState([]string{"One"}, 10, 10)
	state = KeyPress{Key: "<backspace>"}.Apply(state)
	if state.selected != 0 {
		t.Fail()
	}
}

func TestSwitchingFocusToRight(t *testing.T) {
	state := NewAppState([]string{"One"}, 10, 10)
	state.layout = 2
	state = KeyPress{Key: "C-l"}.Apply(state)
	if state.selected != 1 {
		t.Fail()
	}
}

func TestChangingLayout(t *testing.T) {
	fileNames := []string{"1"}
	state := NewAppState(fileNames, 10, 10)
	state = KeyPress{Key: "2"}.Apply(state)
	if state.layout != 2 {
		t.Fail()
	}
}

func TestSpaceTogglesmodifier(t *testing.T) {
	state := NewAppState([]string{"1"}, 10, 10)
	state.CurrentMode = modifierMode
	state.modifiers = modifiers{
		modifier{active: false},
	}
	state.selectedMod = 0

	state = KeyPress{Key: "<space>"}.Apply(state)
	if state.modifiers[state.selectedMod].active != true {
		t.Fail()
	}
}
