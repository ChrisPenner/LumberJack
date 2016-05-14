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
	return ui.NewRow(ui.NewCol(12, 0, par))
}
