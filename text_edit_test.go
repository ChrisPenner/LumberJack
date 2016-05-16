package main

import "testing"

func TestBackSpace(t *testing.T) {
	state := NewAppState()
	store := NewStore()
	state.CurrentMode = selectCategoryMode
	state.selectCategoryBuffer.Text = "a"
	newState := Backspace{}.Apply(state, store.Actions)
	if newState.selectCategoryBuffer.Text != "" {
		t.Fail()
	}

	// Empty Textfield
	newState = Backspace{}.Apply(newState, store.Actions)
	if newState.selectCategoryBuffer.Text != "" {
		t.Fail()
	}
}

func TestTypeKey(t *testing.T) {
	state := NewAppState()
	store := NewStore()
	state.CurrentMode = selectCategoryMode
	state.selectCategoryBuffer.Text = ""
	newState := TypeKey{Key: "a"}.Apply(state, store.Actions)
	if newState.selectCategoryBuffer.Text != "a" {
		t.Fail()
	}
	newState = TypeKey{Key: "b"}.Apply(newState, store.Actions)
	if newState.selectCategoryBuffer.Text != "ab" {
		t.Fail()
	}
}

func TestConvertKey(t *testing.T) {
	if convertKey("<space>") != " " {
		t.Fail()
	}
	if convertKey("a") != "a" {
		t.Fail()
	}
}
