package main

import ui "github.com/gizak/termui"

// StatusBar represents the status bar
type StatusBar struct {
	Text string
}

// Display returns a renderable status bar
func (s StatusBar) display(state AppState) *ui.Row {
	var text string
	if state.CurrentMode == search {
		text = "?: " + state.searchBuffer.text + "_"
	} else {
		text = s.Text
	}
	par := ui.NewPar(text)
	par.Border = false
	par.Height = 1
	par.TextFgColor = ui.ColorCyan
	return ui.NewRow(ui.NewCol(12, 0, par))
}
