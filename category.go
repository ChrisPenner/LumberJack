package main

import ui "github.com/gizak/termui"
import "strings"

// Categories contains info about a particular file
type Categories struct {
	Items []string
}

// Display returns a par for the categories
func (c Categories) Display() *ui.Row {
	par := ui.NewPar(strings.Join(c.Items, ", "))
	par.Border = false
	par.Height = 1
	return ui.NewRow(ui.NewCol(12, 0, par))
}

func (c Categories) getFiltered(state AppState) []string {
	return getLiteralMatches(state.selectCategoryBuffer.Text, c.Items)
}

func (c Categories) getBestMatch(state AppState) (string, bool) {
	filtered := c.getFiltered(state)
	if len(state.selectCategoryBuffer.Text) == 0 {
		return "", false
	}
	if len(filtered) > 0 {
		return c.getFiltered(state)[0], true
	}
	return "", false
}

// SelectCategory selects a category
type SelectCategory struct {
	FileName string
}

// Apply SelectCategory
func (action SelectCategory) Apply(state AppState, actions chan<- Action) AppState {
	state.LogViews[state.selected] = action.FileName
	return state
}
