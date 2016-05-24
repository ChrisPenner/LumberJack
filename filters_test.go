package main

import "testing"

func TestFiltering(t *testing.T) {
	file := File{"one", "twox", "three", "xfour", "fivex"}
	filters := []filter{
		filter{textBuffer: textBuffer{text: "x"}, active: true},
	}
	filtered := file.filter(filters, 2, 1)
	expected := File{"twox", "xfour", "fivex"}
	if len(filtered) != 3 || filtered[0] != expected[0] || filtered[1] != expected[1] || filtered[2] != expected[2] {
		t.Error(filtered)
	}
}

func TestSelectingFilters(t *testing.T) {
	state := NewAppState([]string{"one"}, 10, 10)
	state.filters = filters{
		filter{},
		filter{},
	}
	state.CurrentMode = filterMode
	state = KeyPress{Key: "j"}.Apply(state)
	if state.selectedFilter != 1 {
		t.Fail()
	}
	state = KeyPress{Key: "k"}.Apply(state)
	if state.selectedFilter != 0 {
		t.Fail()
	}
}

func TestSelectingFiltersTooFar(t *testing.T) {
	state := NewAppState([]string{"one"}, 10, 10)
	state.filters = filters{
		filter{},
	}
	state.CurrentMode = filterMode
	state = KeyPress{Key: "k"}.Apply(state)
	if state.selectedFilter != 0 {
		t.Fail()
	}
}

func TestSelectingEmptyFilter(t *testing.T) {
	state := NewAppState([]string{"one"}, 10, 10)
	state.filters = filters{
		filter{},
	}
	state.CurrentMode = filterMode
	state = KeyPress{Key: "j"}.Apply(state)
	if state.selectedFilter != 1 {
		t.Fail()
	}
}

func TestToggleFilter(t *testing.T) {
	state := NewAppState([]string{"one"}, 10, 10)
	state.filters = filters{
		filter{active: true},
	}

	state = state.toggleFilter(0)
	if state.filters[0].active != false {
		t.Fail()
	}

	state = state.toggleFilter(0)
	if state.filters[0].active != true {
		t.Fail()
	}
}

func TestSelectingBackToNormalMode(t *testing.T) {
	state := NewAppState([]string{"one"}, 10, 10)
	state.layout = 3
	state.selected = 0
	state.CurrentMode = filterMode

	state = KeyPress{Key: "<backspace>"}.Apply(state)

	if state.selected != 2 || state.CurrentMode != normal {
		t.Error(state.selected)
	}
}

func TestAddingNewFilter(t *testing.T) {
	state := NewAppState([]string{"one"}, 10, 10)
	state.filters = filters{
		filter{},
	}
	state.selectedFilter = 1
	state.CurrentMode = filterMode
	state = KeyPress{Key: "<enter>"}.Apply(state)
	if len(state.filters) != 2 {
		t.Fail()
	}
}

func TestTabFocusesFilterMode(t *testing.T) {
	state := NewAppState([]string{"one"}, 10, 10)
	state.CurrentMode = normal
	state = KeyPress{Key: "<tab>"}.Apply(state)
	if state.CurrentMode != filterMode {
		t.Fail()
	}
}
