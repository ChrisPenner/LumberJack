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
	actions := make(chan Action, 100)
	state = Scroll{Direction: up, NumLines: 3}.Apply(state, actions)
	if state.getSelectedView().offSet != 3 {
		t.Fail()
	}
	state = Scroll{Direction: down, NumLines: 2}.Apply(state, actions)
	if state.getSelectedView().offSet != 1 {
		t.Fail()
	}
}

func TestScrollDownPastEnd(t *testing.T) {
	state := NewAppState([]string{"One"}, 5)
	// Termheight is 5, logview height will be 3
	state.Files["One"] = []string{"1", "2", "3", "4", "5"}
	state.selected = 0
	actions := make(chan Action, 100)
	state = Scroll{Direction: down, NumLines: 1}.Apply(state, actions)
	if state.getSelectedView().offSet != 0 {
		t.Fail()
	}
}

func TestScrollUpTooHigh(t *testing.T) {
	state := NewAppState([]string{"One"}, 5)
	// Termheight is 5, logview height will be 3
	state.Files["One"] = []string{"1", "2", "3", "4", "5"}
	state.selected = 0
	actions := make(chan Action, 100)
	state = Scroll{Direction: up, NumLines: 30}.Apply(state, actions)
	if state.getSelectedView().offSet != 4 {
		t.Fail()
	}
}

func TestScrollToBottom(t *testing.T) {
	state := NewAppState([]string{"1"}, 1)
	state.Files["1"] = []string{"1", "2", "3", "4", "5"}
	state.LogViews[state.selected].offSet = 4
	actions := make(chan Action, 100)
	state = KeyPress{Key: "g"}.Apply(state, actions)
	action := <-actions
	scrollAction, ok := action.(Scroll)
	if !ok || scrollAction.Direction != bottom {
		t.Error("keypress didn't trigger scroll to bottom")
	}
	state = Scroll{Direction: bottom}.Apply(state, actions)
	if state.getSelectedView().offSet != 0 {
		t.Fail()
	}
}
