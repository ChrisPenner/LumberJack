package main

import ui "github.com/gizak/termui"

// StatusBar represents the status bar
type StatusBar struct {
	Text string
}

// Display returns a renderable status bar
func (s StatusBar) Display() *ui.Row {
	par := ui.NewPar(s.Text)
	par.Border = false
	par.Height = 1
	par.TextFgColor = ui.ColorCyan
	return ui.NewRow(ui.NewCol(12, 0, par))
}

// InitStatusBar sets up the status bar
type InitStatusBar struct {
}

// Apply the InitStatusBar
func (action InitStatusBar) Apply(state *AppState) {
	state.StatusBar = StatusBar{Text: "NewText!"}
}
