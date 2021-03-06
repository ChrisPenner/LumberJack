package main

import "testing"

func TestFiltering(t *testing.T) {
	file := newFile(lines{"one", "twox", "three", "xfour", "fivex"})
	state := NewAppState([]string{"1", "2"}, 10, 10)
	state.modifiers = modifiers{
		modifier{buffer: buffer{"x"}, active: true, kind: filter},
	}
	filtered := file.filter(state)
	expected := lines{"twox", "xfour", "fivex"}
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
	state := fixtureState()
	view := newFile(lines{"line one", "love two", "three"})
	state.modifiers = modifiers{
		modifier{
			kind:    highlighter,
			active:  true,
			fgColor: "green",
			bgColor: "white",
			buffer:  buffer{"l\\w*e"},
		},
	}
	highlighted := view.highlight(state)
	goodLength := len(highlighted) == 3
	lineOne := "[line](fg-green,bg-white) one" == highlighted[0]
	lineTwo := "[love](fg-green,bg-white) two" == highlighted[1]
	lineThree := "three" == highlighted[2]
	if !goodLength || !lineOne || !lineTwo || !lineThree {
		t.Error(highlighted)
	}
}

// Benchmarks

func BenchmarkDisplayModifiers(b *testing.B) {
	state := NewAppState([]string{"1", "2"}, 10, 10)
	state.Files["1"] = sampleFile
	state.Files["2"] = sampleFile
	state.modifiers = modifiers{
		modifier{buffer: buffer{"\\d"}, active: true, kind: filter},
	}
	for i := 0; i < b.N; i++ {
		state.modifiers.display(state)
	}
}
