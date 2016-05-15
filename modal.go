package main

import ui "github.com/gizak/termui"

// Modal interface
type Modal interface {
	Display(*appState) *ui.Grid
	Done(*appState)
}

// SelectCategoryModal struct
type SelectCategoryModal struct {
}

// Done completes the action
func (m SelectCategoryModal) Done(state *appState) {
	state.CurrentModal = nil
	state.HandleKeypress = state.CommandModeHandleKey
	state.textBuffer = ""
}

// Display the modal
func (m SelectCategoryModal) Display(state *appState) *ui.Grid {
	text := state.textBuffer
	height := ui.TermHeight()
	width := ui.TermWidth()
	par := ui.NewPar("Select a File: " + text + "_")
	par.Height = 3
	row := ui.NewRow(ui.NewCol(6, 3, par))
	grid := ui.NewGrid(
		row,
	)
	grid.Width = width
	grid.Y = height/2 - 1
	grid.Align()
	return grid
}
