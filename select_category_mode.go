package main

import "regexp"
import ui "github.com/gizak/termui"

const selectCategoryMode = "selectCategoryMode"

func renderSelectCategoryModal(state AppState) {
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
