package main

import "testing"

func TestEnterSearchMode(t *testing.T) {
	state := NewAppState([]string{"One"}, 10, 10)
	state.CurrentMode = normal
	state = KeyPress{Key: "?"}.Apply(state)
	if state.CurrentMode != search {
		t.Fail()
	}

	state.CurrentMode = normal
	state = KeyPress{Key: "/"}.Apply(state)
	if state.CurrentMode != search {
		t.Fail()
	}
}

func TestExitSearchModeEscape(t *testing.T) {
	state := NewAppState([]string{"One"}, 10, 10)
	state.CurrentMode = search
	state = KeyPress{Key: "<escape>"}.Apply(state)
	if state.CurrentMode != normal {
		t.Fail()
	}
}

func TestExitSearchModeEnter(t *testing.T) {
	state := NewAppState([]string{"One"}, 10, 10)
	state.CurrentMode = search
	state = KeyPress{Key: "<enter>"}.Apply(state)
	if state.CurrentMode != normal {
		t.Fail()
	}
}

func TestSearchTyping(t *testing.T) {
	state := NewAppState([]string{"One"}, 10, 10)
	state.CurrentMode = search
	state = KeyPress{Key: "a"}.Apply(state)
	if state.searchBuffer.text != "a" {
		t.Fail()
	}
}

func TestTypeKeySetsViewOffsetForSearch(t *testing.T) {
	state := NewAppState([]string{"1"}, 2, 10)
	state.Files["1"] = file{lines: lines{"abc1efg", "test", "other", "things", "out"}}
	state.CurrentMode = search
	state = state.typeKey("1")
	if state.searchBuffer.text != "1" {
		t.Fail()
	}
	if state.getSelectedView().offSet != 4 {
		t.Fail()
	}
}

func TestIncrementalSearch(t *testing.T) {
	state := NewAppState([]string{"1"}, 2, 10)
	state.Files["1"] = file{lines: lines{"banana", "test", "bana", "ban", "one"}}
	state.CurrentMode = search
	state.searchBuffer.text = "ba"
	state = state.typeKey("n")
	if state.getSelectedView().offSet != 1 {
		t.Fail()
	}
	state = state.typeKey("a")
	if state.getSelectedView().offSet != 2 {
		t.Fail()
	}
	state = state.typeKey("n")
	if state.getSelectedView().offSet != 4 {
		t.Fail()
	}
}

func TestFindNextMatch(t *testing.T) {
	state := NewAppState([]string{"1"}, 1, 10)
	state.Files["1"] = file{lines: lines{"banana", "test", "bana", "ban", "one"}}
	state.CurrentMode = normal
	state.searchBuffer.text = "ba"
	state = state.findNext(up)
	if state.searchIndex != 1 {
		t.Error("incorrect search index")
	}
	if state.getSelectedView().offSet != 2 {
		t.Error("incorrect view offset")
	}
	state = state.findNext(up)
	if state.getSelectedView().offSet != 4 {
		t.Error("incorrect view offset2")
	}
	if state.searchIndex != 2 {
		t.Error("incorrect search index2")
	}
}

func TestClearsSearchBufferAndIndexOnEnter(t *testing.T) {
	state := NewAppState([]string{"One"}, 10, 10)
	state.searchBuffer.text = "asdf"
	state.searchIndex = 3
	state.CurrentMode = normal
	state = state.changeMode(search)
	if state.searchBuffer.text != "" {
		t.Error("Failed to clear buffer")
	}
	if state.searchIndex != 0 {
		t.Error("Failed to clear searchIndex")
	}
}
