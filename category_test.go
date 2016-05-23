package main

import "testing"

func TestEnterSelectsCategory(t *testing.T) {
	fileNames := []string{"One", "Two", "Three"}
	state := NewAppState(fileNames, 10, 10)
	state.CurrentMode = selectCategory
	state.selectCategoryBuffer.text = "Thr"
	state = KeyPress{Key: "<enter>"}.Apply(state)
	if state.getSelectedView().FileName != "Three" {
		t.Fail()
	}
}
