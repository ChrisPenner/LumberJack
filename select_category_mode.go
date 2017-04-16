package main

import ui "github.com/gizak/termui"

func renderSelectCategoryModal(state *AppState) {
	height := state.termHeight
	width := state.termWidth
	text := state.selectCategoryBuffer.text

	par := ui.NewPar("Select a File: " + text + "_")
	par.Height = 3
	par.BorderFg = ui.ColorYellow

	list := ui.NewList()
	list.Items = state.getFilteredFileNames()
	list.Height = 10
	list.BorderFg = ui.ColorYellow

	row := ui.NewRow(ui.NewCol(6, 3, par, list))
	grid := ui.NewGrid(
		row,
	)
	grid.Width = width
	grid.Y = height/2 - 10
	grid.Align()
	ui.Render(grid)
}
