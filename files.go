package main

import tail "github.com/hpcloud/tail"

// Files list
type Files map[string]file

type lines []string

// file contains the lines of a given file
type file struct {
	lines
	*filteredFileSelector
}

func newFile(l lines) file {
	return file{
		lines:                l,
		filteredFileSelector: &filteredFileSelector{},
	}
}

type filteredFileSelector struct {
	lastLen                int
	lastHlAndFiltered      lines
	lastHLFilteredSearched lines
	lastModifiers          modifiers
	lastSearchTerm         string
	lastHeight             int
	lastOffSet             int
}

func (f file) hlAndFiltered(state AppState) lines {
	filteredLines := f.lastHlAndFiltered
	searchTerm := state.searchBuffer.text
	if f.lastLen != len(f.lines) || (!state.modifiers.isEqual(f.lastModifiers)) {
		// Copy actual modifier structs, not just the slice.
		newModifiers := []modifier{}
		for _, mod := range state.modifiers {
			newModifiers = append(newModifiers, mod)
		}
		f.lastModifiers = newModifiers
		filteredLines = f.lines.filter(state)
		filteredLines = filteredLines.highlight(state)
		f.lastHlAndFiltered = filteredLines
		f.lastHLFilteredSearched = filteredLines.highlightMatches(searchTerm)
	} else if f.lastSearchTerm != searchTerm || f.lastLen != len(f.lines) {
		f.lastSearchTerm = searchTerm
		f.lastHLFilteredSearched = f.lastHlAndFiltered.highlightMatches(searchTerm)
		f.lastSearchTerm = searchTerm
	}
	f.lastLen = len(f.lines)
	return f.lastHLFilteredSearched
}

func (state AppState) getSelectedFileName() string {
	return state.LogViews[state.selected].FileName
}

func (state AppState) getSelectedView() LogView {
	return state.LogViews[state.selected]
}

func (state AppState) getSelectedFile() file {
	return state.Files[state.getSelectedFileName()]
}

func watchFile(fileName string, actions chan<- Action) {
	addNewLine := func(fileName string, newLine string) {
		actions <- AppendLine{FileName: fileName, Line: newLine}
	}
	go tailFile(fileName, addNewLine)
}

func addWatchers(fileNames []string, actions chan<- Action) {
	for _, fileName := range fileNames {
		watchFile(fileName, actions)
	}
}

// AppendLine to file
type AppendLine struct {
	FileName string
	Line     string
}

// Apply the AppendLine
func (action AppendLine) Apply(state AppState) AppState {
	file := state.Files[action.FileName]
	file.lines = append(file.lines, action.Line)
	state.Files[action.FileName] = file
	return state
}

func tailFile(fileName string, callback func(string, string)) {
	t, err := tail.TailFile(fileName, tail.Config{
		Follow:    true,
		Logger:    tail.DiscardingLogger,
		MustExist: true,
	})
	if err != nil {
		panic(err)
	}
	for line := range t.Lines {
		callback(fileName, line.Text)
	}
}
