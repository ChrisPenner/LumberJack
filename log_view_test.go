package main

import "testing"

func TestGetFileSliceInRange(t *testing.T) {
	file := File{"1", "2", "3", "4", "5", "6"}
	view := LogView{offSet: 2}
	slice := file.getVisibleSlice(view, 3)
	if len(slice) != 3 || slice[0] != "2" {
		t.Fail()
	}
}

func TestGetFileSlicePastStart(t *testing.T) {
	file := File{"1", "2", "3", "4", "5", "6"}
	view := LogView{offSet: 2}
	slice := file.getVisibleSlice(view, 6)
	if len(slice) != 6 || slice[0] != "1" {
		t.Fail()
	}
}

func TestGetFileSliceMoreVisibleThanLines(t *testing.T) {
	file := File{"1", "2"}
	view := LogView{offSet: 0}
	slice := file.getVisibleSlice(view, 6)
	if len(slice) != 2 || slice[0] != "1" {
		t.Fail()
	}
}

func TestScroll(t *testing.T) {
	state := NewAppState([]string{"One"}, 5)
	// Termheight is 5, logview height will be 3
	state.Files["One"] = []string{"1", "2", "3", "4", "5"}
	state.selected = 0
	state = state.scroll(up, 3)
	if state.getSelectedView().offSet != 3 {
		t.Fail()
	}
	state = state.scroll(down, 2)
	if state.getSelectedView().offSet != 1 {
		t.Fail()
	}
}

func TestScrollDownPastEnd(t *testing.T) {
	state := NewAppState([]string{"One"}, 5)
	// Termheight is 5, logview height will be 3
	state.Files["One"] = []string{"1", "2", "3", "4", "5"}
	state.selected = 0
	state = state.scroll(down, 10)
	if state.getSelectedView().offSet != 0 {
		t.Fail()
	}
}

func TestScrollUpTooHigh(t *testing.T) {
	state := NewAppState([]string{"One"}, 5)
	// Termheight is 5, logview height will be 3
	state.Files["One"] = []string{"1", "2", "3", "4", "5"}
	state.selected = 0
	state = state.scroll(up, 30)
	if state.getSelectedView().offSet != 4 {
		t.Fail()
	}
}

func TestScrollToBottom(t *testing.T) {
	state := NewAppState([]string{"1"}, 1)
	state.Files["1"] = []string{"1", "2", "3", "4", "5"}
	state.LogViews[state.selected].offSet = 4
	state = state.scroll(bottom, 0)
	if state.getSelectedView().offSet != 0 {
		t.Fail()
	}
}

func TestToggleWrapping(t *testing.T) {
	state := NewAppState([]string{"One"}, 10)
	orig := state.wrap
	state = KeyPress{Key: "w"}.Apply(state)
	if state.wrap != (!orig) {
		t.Fail()
	}
}
