package main

import "testing"

func TestFiltering(t *testing.T) {
	file := File{"one", "twox", "three", "xfour", "fivex"}
	modifiers := modifiers{
		modifier{buffer: buffer{"x"}, active: true, kind: filter},
	}
	filtered := file.filter(modifiers, 2, 1)
	expected := File{"twox", "xfour", "fivex"}
	if len(filtered) != 3 || filtered[0] != expected[0] || filtered[1] != expected[1] || filtered[2] != expected[2] {
		t.Error(filtered)
	}
}

func TestSelectingmodifiers(t *testing.T) {
	state := NewAppState([]string{"one"}, 10, 10)
	state.modifiers = modifiers{
		modifier{},
		modifier{},
	}
	state.CurrentMode = modifierMode
	state = KeyPress{Key: "j"}.Apply(state)
	if state.selectedMod != 1 {
		t.Fail()
	}
	state = KeyPress{Key: "k"}.Apply(state)
	if state.selectedMod != 0 {
		t.Fail()
	}
}

func TestSelectingmodifiersTooFar(t *testing.T) {
	state := NewAppState([]string{"one"}, 10, 10)
	state.modifiers = modifiers{
		modifier{},
	}
	state.CurrentMode = modifierMode
	state = KeyPress{Key: "k"}.Apply(state)
	if state.selectedMod != 0 {
		t.Fail()
	}
}

func TestSelectingEmptyMod(t *testing.T) {
	state := NewAppState([]string{"one"}, 10, 10)
	state.modifiers = modifiers{
		modifier{},
	}
	state.CurrentMode = modifierMode
	state = KeyPress{Key: "j"}.Apply(state)
	if state.selectedMod != 1 {
		t.Fail()
	}
}

func TestToggleMod(t *testing.T) {
	state := NewAppState([]string{"one"}, 10, 10)
	state.modifiers = modifiers{
		modifier{active: true},
	}

	state = state.toggleModifier(0)
	if state.modifiers[0].active != false {
		t.Fail()
	}

	state = state.toggleModifier(0)
	if state.modifiers[0].active != true {
		t.Fail()
	}
}

func TestSelectingBackToNormalMode(t *testing.T) {
	state := NewAppState([]string{"one"}, 10, 10)
	state.layout = 3
	state.selected = 0
	state.CurrentMode = modifierMode

	state = KeyPress{Key: "<backspace>"}.Apply(state)

	if state.selected != 2 || state.CurrentMode != normal {
		t.Error(state.selected)
	}
}

func TestAddingNewFilter(t *testing.T) {
	state := NewAppState([]string{"one"}, 10, 10)
	state.modifiers = modifiers{
		modifier{},
	}
	state.selectedMod = 1
	state.CurrentMode = modifierMode
	state = KeyPress{Key: "<enter>"}.Apply(state)
	if len(state.modifiers) != 2 {
		t.Fail()
	}
}

func TestTabFocusesFilterMode(t *testing.T) {
	state := NewAppState([]string{"one"}, 10, 10)
	state.CurrentMode = normal
	state = KeyPress{Key: "<tab>"}.Apply(state)
	if state.CurrentMode != modifierMode {
		t.Fail()
	}
}

func TestHighlightsLines(t *testing.T) {
	view := File{"line one", "love two", "three"}
	highlighters := modifiers{
		modifier{
			kind:    highlighter,
			active:  true,
			fgColor: "green",
			bgColor: "white",
			buffer:  buffer{"l\\w*e"},
		},
	}
	view = view.highlight(highlighters)
	goodLength := len(view) == 3
	lineOne := "[line](fg-green,bg-white) one" == view[0]
	lineTwo := "[love](fg-green,bg-white) two" == view[1]
	lineThree := "three" == view[2]
	if !goodLength || !lineOne || !lineTwo || !lineThree {
		t.Error(view)
	}
}
