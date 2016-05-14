package main

import ui "github.com/gizak/termui"

func main() {
	// keyMap := map[string]string{
	// 	"<space>": " ",
	// }

	err := ui.Init()
	if err != nil {
		panic(err)
	}
	defer ui.Close()

	logView1 := ui.NewPar("String 1")
	logView2 := ui.NewPar("String 2")
	logView1.Height = ui.TermHeight() - 3
	logView2.Height = ui.TermHeight() - 3
	logView1.BorderLabel = "Logs 1"
	logView2.BorderLabel = "Logs 2"
	logView1.BorderRight = false
	logViews := ui.NewRow(
		ui.NewCol(6, 0, logView1),
		ui.NewCol(6, 0, logView2),
	)

	categories := ui.NewPar("Category 1, Category 2")
	categories.Border = false
	categories.Height = 1
	categoriesRow := ui.NewRow(
		ui.NewCol(12, 0, categories),
	)

	status := ui.NewPar("Status bar...")
	status.Border = false
	status.Height = 1
	statusBar := ui.NewRow(
		ui.NewCol(12, 0, status),
	)

	// build layout
	ui.Body.AddRows(
		categoriesRow,
		logViews,
		statusBar,
	)

	// calculate layout
	ui.Body.Align()

	ui.Render(ui.Body)
	// ui.Render(buf)

	ui.Handle("/sys/kbd/C-c", func(ui.Event) {
		ui.StopLoop()
	})
	ui.Handle("/sys/kbd/", func(e ui.Event) {
		keyPress := e.Data.(ui.EvtKbd).KeyStr
		status.Text = keyPress
		ui.Render(ui.Body)
	})
	ui.Handle("/sys/wnd/resize", func(e ui.Event) {
		// newHeight := e.Data.(ui.EvtWnd).Height
		newWidth := e.Data.(ui.EvtWnd).Width
		newHeight := e.Data.(ui.EvtWnd).Height
		logView1.Height = newHeight - 3
		logView2.Height = newHeight - 3
		status.Text = "Resized"
		ui.Body.Width = newWidth
		ui.Body.Align()
		ui.Render(ui.Body)
	})
	ui.Loop()
}
