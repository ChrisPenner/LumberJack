package main

import (
	"fmt"
	"strconv"
	"strings"

	ui "github.com/gizak/termui"
)

// LogViews is a list of viewnames
type LogViews []LogView

// LogView represents view into logs
type LogView struct {
	FileName string
	offSet   int
}
type initLogViews struct{}

// Display returns a Row object representing all of the logViews
func (logViews LogViews) display(state AppState) []*ui.Row {
	listBlocks := []*ui.List{}
	for i, view := range logViews {
		if i >= state.layout {
			break
		}
		logView := view.display(state)
		logView.BorderFg = ui.ColorWhite
		if state.wrap {
			logView.Overflow = "wrap"
		}
		listBlocks = append(listBlocks, logView)
	}
	if len(listBlocks) > 0 && (state.CurrentMode == normal || state.CurrentMode == selectCategory) {
		listBlocks[state.selected].BorderFg = ui.ColorMagenta
	}
	logViewColumns := []*ui.Row{}

	filterSize := 0
	if state.showMods {
		filterSize = getModifierSpan(state.termWidth)
	}
	numColumnsEach := (12 - filterSize) / state.layout
	leftOver := (12 - filterSize) - (numColumnsEach * state.layout)
	for _, logViewBlock := range listBlocks {
		extra := 0
		if leftOver > 0 {
			extra = 1
			leftOver--
		}
		logViewColumns = append(logViewColumns, ui.NewCol(numColumnsEach+extra, 0, logViewBlock))
	}
	return logViewColumns
}

func (view LogView) display(state AppState) *ui.List {
	list := ui.NewList()
	list.Height = logViewHeight(state.termHeight)
	active := state.getSelectedFileName() == view.FileName
	if active {
		list.BorderFg = ui.ColorWhite
	} else {
		list.BorderFg = ui.ColorYellow
	}
	list.BorderLabel = view.FileName
	file := state.getFile(view.FileName)
	height := view.numVisibleLines(state)
	filteredView := file
	if anyActiveModifiers(state.modifiers) {
		filteredView = file.filter(state.modifiers, height, view.offSet)
	}
	searchTerm := state.searchBuffer.text
	filteredView = filteredView.highlightMatches(searchTerm)
	visibleLines := filteredView.getVisibleSlice(view, height)
	list.Items = visibleLines
	return list
}

func (view LogView) scrollToSearch(state AppState) LogView {
	file := state.Files[view.FileName]
	searchResultOffset := file.getSearchResultLine(state.searchBuffer.text, state.searchIndex)
	if searchResultOffset >= 0 {
		view.offSet = searchResultOffset - (logViewHeight(state.termHeight) / 2)
		if view.offSet < 0 {
			view.offSet = 0
		}
	}
	return view
}

func (view LogView) numVisibleLines(state AppState) int {
	return logViewHeight(state.termHeight) - 2
}

func (file File) getVisibleSlice(view LogView, height int) []string {
	start := (len(file) - height) - view.offSet
	if start < 0 {
		start = 0
	}
	end := start + height
	if end > len(file) {
		end = len(file)
	}
	return file[start:end]
}

func (file File) getSearchResultLine(term string, searchIndex int) int {
	for i := range file {
		line := file[len(file)-i-1]
		if strings.Contains(line, term) {
			if searchIndex <= 0 {
				return i
			}
			searchIndex--
		}
	}
	return -1
}

func (file File) highlightMatches(term string) File {
	if term == "" {
		return file
	}
	var highlightedLines = make(File, len(file))
	for i, line := range file {
		hlTerm := fmt.Sprintf("[%s](bg-yellow,fg-black)", term)
		highlightedLines[i] = strings.Replace(line, term, hlTerm, -1) //hlTerm, -1)
	}
	return highlightedLines
}

// Scroll
func (state AppState) scroll(direction direction, amount int) AppState {
	view := state.getSelectedView()
	file := state.getSelectedFile()
	switch direction {
	case up:
		view.offSet += amount
	case down:
		view.offSet -= amount
	case bottom:
		view.offSet = 0
	}
	if view.offSet > len(file)-view.numVisibleLines(state) {
		view.offSet = len(file) - view.numVisibleLines(state)
	}
	if view.offSet < 0 {
		view.offSet = 0
	}
	state.LogViews[state.selected] = view
	state.StatusBar.Text = strconv.Itoa(state.getSelectedView().offSet)
	return state
}

func anyActiveModifiers(modifiers modifiers) bool {
	for _, m := range modifiers {
		if m.active {
			return true
		}
	}
	return false
}
