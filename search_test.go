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

func TestClearsSearchBufferOnEnter(t *testing.T) {
	state := NewAppState([]string{}, 10)
	state.searchBuffer.text = "asdf"
	state.CurrentMode = normal
	actions := make(chan Action, 100)
	state = ChangeMode{Mode: search}.Apply(state, actions)
	if state.searchBuffer.text != "" {
		t.Fail()
	}
}
