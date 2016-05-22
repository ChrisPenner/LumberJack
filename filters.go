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

type toggleFilter struct {
	filter int
}

func (action toggleFilter) Apply(state AppState, actions chan<- Action) AppState {
	state.filters[action.filter].active = !state.filters[action.filter].active
	return state
}

func (f File) filter(filters []filter) File {
	var filteredLines File
	atLeastOneFilter := false
	for _, line := range f {
		matchFilter := false
		for _, filter := range filters {
			if !filter.active {
				continue
			}
			atLeastOneFilter = true
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
			filteredLines = append(filteredLines, line)
		}
	}
	if !atLeastOneFilter {
		return f
	}
	return filteredLines
}
