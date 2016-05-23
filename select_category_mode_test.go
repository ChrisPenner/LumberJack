package main

import "testing"

func TestEnterSelectCategoryMode(t *testing.T) {
	state := NewAppState([]string{"One"}, 10)
	state.CurrentMode = normal
	state = KeyPress{Key: "<space>"}.Apply(state)
	if state.CurrentMode != selectCategory {
		t.Fail()
	}
}

func TestEnterSwitchesSelectsCategory(t *testing.T) {
	state := NewAppState([]string{"one", "two"}, 10)
	state.CurrentMode = selectCategory
	state.selectCategoryBuffer.text = "tw"
	state = KeyPress{Key: "<enter>"}.Apply(state)
	if state.CurrentMode != normal || state.getSelectedView().FileName != "two" {
		t.Fail()
	}
}

func TestEscapeExitsCategoryModeWithoutSelecting(t *testing.T) {
	state := NewAppState([]string{"one", "two"}, 10)
	state.CurrentMode = selectCategory
	state.selectCategoryBuffer.text = "tw"
	state = KeyPress{Key: "<escape>"}.Apply(state)
	if state.CurrentMode != normal || state.getSelectedView().FileName == "two" {
		t.Fail()
	}
}

func TestEnterSelectsCategoryOfBestMatch(t *testing.T) {
	state := NewAppState([]string{"one", "two"}, 10)
	state.CurrentMode = selectCategory
	state.selectCategoryBuffer.text = "wo" // Submatch
	state = KeyPress{Key: "<enter>"}.Apply(state)
	if state.getSelectedView().FileName != "two" {
		t.Fail()
	}
}
