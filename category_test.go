package main

import "testing"

func TestEnterSelectsCategory(t *testing.T) {
	fileNames := []string{"One", "Two", "Three"}
	state := NewAppState(fileNames)
	store := NewStore()
	state.CurrentMode = selectCategoryMode
	state.selectCategoryBuffer.Text = "On"
	newState := KeyPress{Key: "<enter>"}.Apply(state, store.Actions)
	action := <-store.Actions
	selectCategory, ok := action.(SelectCategory)
	if !ok || selectCategory.FileName != "One" || newState.LogViews.viewNames[0] != "One" {
		t.Fail()
	}
}
