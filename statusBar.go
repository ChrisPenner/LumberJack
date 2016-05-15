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

type initStatusBar struct {
}

func (action initStatusBar) Apply(state *AppState) {
	state.StatusBar = StatusBar{Text: "NewText!"}
}
