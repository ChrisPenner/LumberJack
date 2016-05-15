package main

import ui "github.com/gizak/termui"
import "regexp"

const normalMode = "normalMode"
const selectCategoryMode = "selectCategoryMode"

// ChangeMode changes modes
type ChangeMode struct {
	Mode string
}

// Apply the ChangeMode
func (action ChangeMode) Apply(state *AppState) {
	state.CurrentMode = action.Mode
	state.StatusBar.Text = action.Mode
}

func renderSelectCategoryModal(state *AppState) {
	height := ui.TermHeight()
	width := ui.TermWidth()
	text := state.selectCategoryBuffer.Text

	par := ui.NewPar("Select a File: " + text + "_")
	par.Height = 3

	list := ui.NewList()
	list.Items = getLiteralMatches(text, state.Categories.Items)
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
	key := action.Key
	switch state.CurrentMode {
	case normalMode:
		switch key {
		case "<enter>":
			store.Actions <- ChangeMode{Mode: selectCategoryMode}
		}
	case selectCategoryMode:
		switch key {
		case "C-8":
			store.Actions <- Backspace{}
		case "<enter>":
			store.Actions <- ChangeMode{Mode: normalMode}
			store.Actions <- Backspace{}
		default:
			store.Actions <- TypeKey{Key: convertKey(key)}
		}
	default:
		panic(state.CurrentMode)
	}
}
