package main

import ui "github.com/gizak/termui"
import "strings"

// Categories list
type Categories []string

// Display returns a par for the categories
func (c Categories) Display() *ui.Row {
	par := ui.NewPar(strings.Join(c, " [|](fg-magenta) "))
	par.Border = false
	par.Height = 1
	return ui.NewRow(ui.NewCol(12, 0, par))
}

func (c Categories) getFiltered(state AppState) []string {
	return getLiteralMatches(state.selectCategoryBuffer.text, c)
}

func (c Categories) getBestMatch(state AppState) (string, bool) {
	filtered := c.getFiltered(state)
	if len(state.selectCategoryBuffer.text) == 0 {
		return "", false
	}
	if len(filtered) > 0 {
		return c.getFiltered(state)[0], true
	}
	return "", false
}

func (state AppState) selectCategory(fileName string) AppState {
	selectedView := state.LogViews[state.selected]
	selectedView.FileName = fileName
	state.LogViews[state.selected] = selectedView
	return state
}
