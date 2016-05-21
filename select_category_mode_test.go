package main

import "testing"

func TestEnterSelectCategoryMode(t *testing.T) {
	state := NewAppState([]string{}, 10)
	state.CurrentMode = normal
	actions := make(chan Action, 100)
	KeyPress{Key: "<space>"}.Apply(state, actions)
	action := <-actions
	changeMode, ok := action.(ChangeMode)
	if !ok || changeMode.Mode != selectCategory {
		t.Fail()
	}
}

func TestEnterSwitchesToNormalModeFromSC(t *testing.T) {
	state := NewAppState([]string{}, 10)
	state.CurrentMode = selectCategory
	actions := make(chan Action, 100)
	KeyPress{Key: "<enter>"}.Apply(state, actions)
	action := <-actions
	changeMode, ok := action.(ChangeMode)
	if !ok || changeMode.Mode != normal {
		t.Fail()
	}
}

func TestEscapeExitsCategoryModeWithoutSelecting(t *testing.T) {
	state := NewAppState([]string{"one", "two"}, 10)
	state.CurrentMode = selectCategory
	state.selectCategoryBuffer.text = "on"
	actions := make(chan Action, 100)
	KeyPress{Key: "<escape>"}.Apply(state, actions)
	action := <-actions
	changeMode, ok := action.(ChangeMode)
	if !ok || changeMode.Mode != normal {
		t.Fail()
	}
}

func TestEnterSelectsCategoryOfBestMatch(t *testing.T) {
	state := NewAppState([]string{"one", "two"}, 10)
	state.CurrentMode = selectCategory
	state.selectCategoryBuffer.text = "wo" // Submatch
	actions := make(chan Action, 100)
	KeyPress{Key: "<enter>"}.Apply(state, actions)
	action := <-actions
	selectCategory, ok := action.(SelectCategory)
	if !ok || (selectCategory != SelectCategory{FileName: "two"}) {
		t.Fail()
	}
	action = <-actions
	changeMode, ok := action.(ChangeMode)
	if !ok || changeMode.Mode != normal {
		t.Fail()
	}
}

func TestKeyPressAddsTypeKeyInSelectCategoryMode(t *testing.T) {
	state := NewAppState([]string{}, 10)
	state.CurrentMode = selectCategory
	actions := make(chan Action, 100)
	KeyPress{Key: "a"}.Apply(state, actions)
	action := <-actions
	typeKey, ok := action.(typeKey)
	if !ok || typeKey.key != "a" {
		t.Fail()
	}
}
