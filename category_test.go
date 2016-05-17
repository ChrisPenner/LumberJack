package main

import "testing"

func TestEnterSelectsCategory(t *testing.T) {
	fileNames := []string{"One", "Two", "Three"}
	state := NewAppState(fileNames, 10)
	actions := make(chan Action, 100)
	state.CurrentMode = selectCategoryMode
	state.selectCategoryBuffer.Text = "On"
	newState := KeyPress{Key: "<enter>"}.Apply(state, actions)
	action := <-actions
	selectCategory, ok := action.(SelectCategory)
	if !ok || selectCategory.FileName != "One" || newState.LogViews[0].FileName != "One" {
		t.Fail()
	}
}
