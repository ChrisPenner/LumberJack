package main

import (
	"fmt"
	"regexp"

	ui "github.com/gizak/termui"
)

type filters []filter
type filter struct {
	active     bool
	textBuffer textBuffer
}

func (f filters) display(state AppState) *ui.Row {
	filterList := ui.NewList()
	var listItems []string
	for i, f := range state.filters {
		var attrs, title string
		if f.active {
			attrs = "fg-green"
		} else {
			attrs = "fg-red"
		}
		if i == state.selectedFilter && (state.CurrentMode == filterMode || state.CurrentMode == editFilter) {
			title = fmt.Sprintf("[[%d]](bg-cyan,fg-black) [%s](%s)", i+1, f.textBuffer.text, attrs)
		} else {
			title = fmt.Sprintf("[[%d] %s](%s)", i+1, f.textBuffer.text, attrs)
		}
		if state.CurrentMode == editFilter && i == state.selectedFilter {
			title += "_"
		}
		listItems = append(listItems, title)
	}
	filterList.Items = listItems
	filterList.Height = logViewHeight(state.termHeight)
	if state.CurrentMode == filterMode || state.CurrentMode == editFilter {
		filterList.BorderFg = ui.ColorGreen
	}
	return ui.NewCol(1, 0, filterList)
}

func (state AppState) toggleFilter(filter int) AppState {
	state.filters[filter].active = !state.filters[filter].active
	return state
}

func (f File) filter(filters []filter, height int, offSet int) File {
	var filteredLines File
	for i := range f {
		// Go through file in reverse
		line := f[len(f)-i-1]
		matchFilter := false
		for _, filter := range filters {
			if !filter.active {
				continue
			}
			matched, err := regexp.Match(filter.textBuffer.text, []byte(line))
			if err != nil {
				continue
			}
			if matched {
				matchFilter = true
				break
			}
		}
		if matchFilter {
			// Build up file in reverse
			filteredLines = append([]string{line}, filteredLines...)
			if len(filteredLines) == height+offSet {
				break
			}
		}
	}
	return filteredLines
}
