package main

import ui "github.com/gizak/termui"
import "regexp"

// import "strconv"

// Mode interface
type Mode interface {
	Render(*AppState)
	KeyboardHandler(string)
}

// ChangeMode changes modes
type ChangeMode struct {
	Mode Mode
}

// Apply the ChangeMode
func (action ChangeMode) Apply(state *AppState) {
	state.CurrentMode = action.Mode
}

// NormalMode is the main mode
type NormalMode struct {
}

// NewNormalMode is the NormalMode constructor
func NewNormalMode() Mode {
	return &NormalMode{}
}

// Render the mode
func (m *NormalMode) Render(_ *AppState) {
}

// KeyboardHandler for Normal Mode
func (m *NormalMode) KeyboardHandler(key string) {
	switch key {
	case "<enter>":
		nextMode := NewSelectCategoryMode()
		store.Actions <- ChangeMode{Mode: nextMode}
	}
	// if n, err := strconv.Atoi(key); err == nil {
	// m.state.StatusBar.Text = key
	// m.state.LogViews.Select(n - 1)
	// }
}

// SelectCategoryMode struct
type SelectCategoryMode struct {
	Buffer     *TextBuffer
	categories []string
}

// NewSelectCategoryMode is the NormalMode constructor
func NewSelectCategoryMode() Mode {
	return &SelectCategoryMode{Buffer: new(TextBuffer)}
}

// Render for the mode
func (m *SelectCategoryMode) Render(state *AppState) {
	height := ui.TermHeight()
	width := ui.TermWidth()

	par := ui.NewPar("Select a File: " + m.Buffer.Text + "_")
	par.Height = 3

	list := ui.NewList()
	list.Items = getLiteralMatches(m.Buffer.Text, state.Categories.Items)
	list.Height = 10

	row := ui.NewRow(ui.NewCol(6, 3, par, list))
	grid := ui.NewGrid(
		row,
	)
	grid.Width = width
	grid.Y = height/2 - 10
	grid.Align()
	ui.Render(grid)
}

// KeyboardHandler for SelectCategoryMode
func (m *SelectCategoryMode) KeyboardHandler(key string) {
	switch key {
	case "C-8":
		store.Actions <- Backspace{Buffer: m.Buffer}
	case "<enter>":
		nextMode := NewNormalMode()
		store.Actions <- ChangeMode{Mode: nextMode}
		store.Actions <- Backspace{Buffer: m.Buffer}
	default:
		store.Actions <- TypeKey{Key: key, Buffer: m.Buffer}
	}
}

func getLiteralMatches(pattern string, items []string) []string {
	r, _ := regexp.Compile(pattern)
	return Filter(items, func(s string) bool {
		return r.Match([]byte(s))
	})
}

// KeyPress sends a keypress
type KeyPress struct {
	Key string
}

// Apply the KeyPress
func (action KeyPress) Apply(state *AppState) {
	state.CurrentMode.KeyboardHandler(action.Key)
}
