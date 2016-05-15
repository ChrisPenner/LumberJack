package main

import ui "github.com/gizak/termui"
import "regexp"
import "strconv"

// Mode interface
type Mode interface {
	Render()
	Next() func(*AppState) Mode
	KeyboardHandler(string)
}

// NormalMode is the main mode
type NormalMode struct {
	state *AppState
}

// NewNormalMode is the NormalMode constructor
func NewNormalMode(state *AppState) Mode {
	return &NormalMode{state: state}
}

// Next mode to use
func (m *NormalMode) Next() func(*AppState) Mode {
	return NewSelectCategoryMode
}

// Render the mode
func (m *NormalMode) Render() {
}

// KeyboardHandler for Normal Mode
func (m *NormalMode) KeyboardHandler(key string) {
	switch key {
	case "<enter>":
		m.state.CurrentMode = NewSelectCategoryMode(m.state)
	}
	if n, err := strconv.Atoi(key); err == nil {
		m.state.StatusBar.Text = key
		m.state.LogViews.Select(n - 1)
	}
	renderFlag = true
}

// SelectCategoryMode struct
type SelectCategoryMode struct {
	text       string
	categories []string
}

// NewSelectCategoryMode is the NormalMode constructor
func NewSelectCategoryMode(state *AppState) Mode {
	return &SelectCategoryMode{categories: state.Categories.Items}
}

// Next Mode
func (m *SelectCategoryMode) Next() func(*AppState) Mode {
	return NewNormalMode
}

// Render for the mode
func (m *SelectCategoryMode) Render() {
	height := ui.TermHeight()
	width := ui.TermWidth()

	par := ui.NewPar("Select a File: " + m.text + "_")
	par.Height = 3

	list := ui.NewList()
	list.Items = getLiteralMatches(m.text, m.categories)
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
		// Backspace
		if len(m.text) > 0 {
			m.text = m.text[:len(m.text)-1]
		}
	case "<space>":
		key = " "
		fallthrough
	default:
		m.text = m.text + key
	}
}

func getLiteralMatches(pattern string, items []string) []string {
	r, _ := regexp.Compile(pattern)
	return Filter(items, func(s string) bool {
		return r.Match([]byte(s))
	})
}
