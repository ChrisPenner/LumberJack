package main

import "testing"

func TestEnterSearchMode(t *testing.T) {
	state := NewAppState([]string{}, 10)
	state.CurrentMode = normal
	actions := make(chan Action, 100)
	KeyPress{Key: "?"}.Apply(state, actions)
	action := <-actions
	changeMode, ok := action.(ChangeMode)
	if !ok || changeMode.Mode != search {
		t.Fail()
	}

	state.CurrentMode = normal
	KeyPress{Key: "/"}.Apply(state, actions)
	action = <-actions
	changeMode, ok = action.(ChangeMode)
	if !ok || changeMode.Mode != search {
		t.Fail()
	}
}

func TestExitSearchModeEscape(t *testing.T) {
	state := NewAppState([]string{}, 10)
	state.CurrentMode = search
	actions := make(chan Action, 100)
	KeyPress{Key: "<escape>"}.Apply(state, actions)
	action := <-actions
	changeMode, ok := action.(ChangeMode)
	if !ok || changeMode.Mode != normal {
		t.Fail()
	}
}

func TestExitSearchModeEnter(t *testing.T) {
	state := NewAppState([]string{}, 10)
	state.CurrentMode = search
	actions := make(chan Action, 100)
	KeyPress{Key: "<enter>"}.Apply(state, actions)
	action := <-actions
	changeMode, ok := action.(ChangeMode)
	if !ok || changeMode.Mode != normal {
		t.Fail()
	}
}

func TestSearchAddsTypeKey(t *testing.T) {
	state := NewAppState([]string{}, 10)
	state.CurrentMode = search
	actions := make(chan Action, 100)
	KeyPress{Key: "a"}.Apply(state, actions)
	action := <-actions
	typeKey, ok := action.(typeKey)
	if !ok || typeKey.key != "a" {
		t.Fail()
	}
}

func TestTypeKeySetsViewOffsetForSearch(t *testing.T) {
	state := NewAppState([]string{"1"}, 2)
	state.Files["1"] = File{"abc1efg", "test", "other", "things", "out"}
	state.CurrentMode = search
	actions := make(chan Action, 100)
	state = typeKey{key: "1"}.Apply(state, actions)
	if state.searchBuffer.text != "1" {
		t.Fail()
	}
	if state.getSelectedView().offSet != 4 {
		t.Fail()
	}
}

func TestIncrementalSearch(t *testing.T) {
	state := NewAppState([]string{"1"}, 2)
	state.Files["1"] = File{"banana", "test", "bana", "ban", "one"}
	state.CurrentMode = search
	state.searchBuffer.text = "ba"
	actions := make(chan Action, 100)
	state = typeKey{key: "n"}.Apply(state, actions)
	if state.getSelectedView().offSet != 1 {
		t.Fail()
	}
	state = typeKey{key: "a"}.Apply(state, actions)
	if state.getSelectedView().offSet != 2 {
		t.Fail()
	}
	state = typeKey{key: "n"}.Apply(state, actions)
	if state.getSelectedView().offSet != 4 {
		t.Fail()
	}
}

func TestTriggerFindNext(t *testing.T) {
	state := NewAppState([]string{}, 10)
	state.CurrentMode = normal
	actions := make(chan Action, 100)
	state = KeyPress{Key: "n"}.Apply(state, actions)
	action := <-actions
	findNextA, ok := action.(findNext)
	if !ok || findNextA.direction != up {
		t.Error("n didn't work properly")
	}

	state = KeyPress{Key: "N"}.Apply(state, actions)
	action = <-actions
	findNextA, ok = action.(findNext)
	if !ok || findNextA.direction != down {
		t.Error("N didn't work properly")
	}

}

func TestFindNextMatch(t *testing.T) {
	state := NewAppState([]string{"1"}, 1)
	state.Files["1"] = File{"banana", "test", "bana", "ban", "one"}
	state.CurrentMode = normal
	state.searchBuffer.text = "ba"
	actions := make(chan Action, 100)
	state = findNext{direction: up}.Apply(state, actions)
	if state.searchIndex != 1 {
		t.Error("incorrect search index")
	}
	if state.getSelectedView().offSet != 2 {
		t.Error("incorrect view offset")
	}
	state = findNext{direction: up}.Apply(state, actions)
	if state.getSelectedView().offSet != 4 {
		t.Error("incorrect view offset2")
	}
	if state.searchIndex != 2 {
		t.Error("incorrect search index2")
	}
}

func TestClearsSearchBufferAndIndexOnEnter(t *testing.T) {
	state := NewAppState([]string{}, 10)
	state.searchBuffer.text = "asdf"
	state.searchIndex = 3
	state.CurrentMode = normal
	actions := make(chan Action, 100)
	state = ChangeMode{Mode: search}.Apply(state, actions)
	if state.searchBuffer.text != "" {
		t.Error("Failed to clear buffer")
	}
	if state.searchIndex != 0 {
		t.Error("Failed to clear searchIndex")
	}
}
