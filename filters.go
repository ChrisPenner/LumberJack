package main

import (
	"fmt"
	"regexp"

	ui "github.com/gizak/termui"
)

type filters []filter
type filter struct {
	active  bool
	pattern string
}

func (f filters) display(state AppState) *ui.Row {
	filterList := ui.NewList()
	var listItems []string
	for i, f := range state.filters {
		var title string
		if f.active {
			title = fmt.Sprintf("[[%d] %s](fg-green)", i+1, f.pattern)
		} else {
			title = fmt.Sprintf("[[%d] %s](fg-red)", i+1, f.pattern)
		}
		listItems = append(listItems, title)
	}
	filterList.Items = listItems
	filterList.Height = logViewHeight(state.termHeight)
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
			matched, err := regexp.Match(filter.pattern, []byte(line))
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
