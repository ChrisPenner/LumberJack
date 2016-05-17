package main

import "testing"

func TestEnterSwitchesModesInNormalMode(t *testing.T) {
	state := NewAppState([]string{}, 10)
	state.CurrentMode = normalMode
	actions := make(chan Action, 100)
	KeyPress{Key: "<enter>"}.Apply(state, actions)
	action := <-actions
	changeMode, ok := action.(ChangeMode)
	if !ok || changeMode.Mode != selectCategoryMode {
		t.Fail()
	}
}

func TestSwitchingFocus(t *testing.T) {
	state := NewAppState([]string{}, 10)
	state.CurrentMode = normalMode
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

func TestEnterSwitchesModesInSelectCategory(t *testing.T) {
	state := NewAppState([]string{}, 10)
	state.CurrentMode = selectCategoryMode
	actions := make(chan Action, 100)
	KeyPress{Key: "<enter>"}.Apply(state, actions)
	action := <-actions
	changeMode, ok := action.(ChangeMode)
	if !ok || changeMode.Mode != normalMode {
		t.Fail()
	}
}

func TestEscapeExitsCategoryModeWithoutSelecting(t *testing.T) {
	state := NewAppState([]string{"one", "two"}, 10)
	state.CurrentMode = selectCategoryMode
	state.selectCategoryBuffer.Text = "on"
	actions := make(chan Action, 100)
	KeyPress{Key: "<escape>"}.Apply(state, actions)
	action := <-actions
	changeMode, ok := action.(ChangeMode)
	if !ok || changeMode.Mode != normalMode {
		t.Fail()
	}
}

func TestEnterSelectsCategoryOfBestMatch(t *testing.T) {
	state := NewAppState([]string{"one", "two"}, 10)
	state.CurrentMode = selectCategoryMode
	state.selectCategoryBuffer.Text = "wo" // Submatch
	actions := make(chan Action, 100)
	KeyPress{Key: "<enter>"}.Apply(state, actions)
	action := <-actions
	selectCategory, ok := action.(SelectCategory)
	if !ok || (selectCategory != SelectCategory{FileName: "two"}) {
		t.Fail()
	}
	action = <-actions
	changeMode, ok := action.(ChangeMode)
	if !ok || changeMode.Mode != normalMode {
		t.Fail()
	}
}

func TestKeyPressAddsTypeKeyInSelectCategoryMode(t *testing.T) {
	state := NewAppState([]string{}, 10)
	state.CurrentMode = selectCategoryMode
	actions := make(chan Action, 100)
	KeyPress{Key: "a"}.Apply(state, actions)
	action := <-actions
	typeKey, ok := action.(TypeKey)
	if !ok || typeKey.Key != "a" {
		t.Fail()
	}
}

func TestKeyPressAddsBackspaceInSelectCategoryMode(t *testing.T) {
	state := NewAppState([]string{}, 10)
	state.CurrentMode = selectCategoryMode
	actions := make(chan Action, 100)
	KeyPress{Key: "C-8"}.Apply(state, actions)
	action := <-actions
	_, ok := action.(Backspace)
	if !ok {
		t.Fail()
	}
}
