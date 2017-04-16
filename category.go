package main

import (
	"fmt"
	"strings"

	ui "github.com/gizak/termui"
)

func (state *AppState) displayFileBar() *ui.Row {
	fileNames := ""
	for _, fileName := range state.orderedFileNames {
		file := state.Files[fileName]
		fileNames += fmt.Sprintf("%s [(%d)](fg-yellow) [|](fg-magenta) ", fileName, file.numUnread)
	}
	par := ui.NewPar(fileNames)
	par.Border = false
	par.Height = 1
	return ui.NewRow(ui.NewCol(12, 0, par))
}

func (state *AppState) getFilteredFileNames() []string {
	pattern := state.selectCategoryBuffer.text
	var results []string
	for _, fileName := range state.orderedFileNames {
		if strings.Contains(fileName, pattern) {
			results = append(results, fileName)
		}
	}
	return results
}

func (state *AppState) getBestMatch() (string, bool) {
	filtered := state.getFilteredFileNames()
	if len(state.selectCategoryBuffer.text) == 0 {
		return "", false
	}
	if len(filtered) > 0 {
		return filtered[0], true
	}
	return "", false
}

func (state *AppState) selectCategory(fileName string) *AppState {
	selectedView := state.LogViews[state.selected]
	selectedView.FileName = fileName
	state.LogViews[state.selected] = selectedView
	state = state.clearUnreadCounts()
	return state
}

func (state *AppState) clearUnreadCounts() *AppState {
	for i := 0; i < state.layout; i++ {
		f := state.Files[state.LogViews[i].FileName]
		f.numUnread = 0
		state.Files[state.LogViews[i].FileName] = f
	}
	return state
}
