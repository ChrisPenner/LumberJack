package main

import ui "github.com/gizak/termui"

// StatusBar represents the status bar
type StatusBar struct {
	Text string
}

// Display returns a renderable status bar
func (s StatusBar) Display(mode string) *ui.Row {
	par := ui.NewPar(s.Text)
	par.Border = false
	par.Height = 1
	par.TextFgColor = ui.ColorCyan
	par.Text = mode
	return ui.NewRow(ui.NewCol(12, 0, par))
}
