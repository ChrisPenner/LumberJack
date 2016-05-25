package main

import "sort"

// AppState contains global state
type AppState struct {
	termHeight           int
	termWidth            int
	CurrentMode          mode
	LogViews             LogViews
	Files                Files
	Categories           Categories
	StatusBar            StatusBar
	HandleKeypress       func(string)
	selected             int
	searchIndex          int
	selectCategoryBuffer buffer
	searchBuffer         buffer
	wrap                 bool
	layout               int
	modifiers            modifiers
	showMods             bool
	selectedMod          int
}

// NewAppState constructs and appstate
func NewAppState(fileNames []string, height int, width int) AppState {
	sort.Strings(fileNames)
	files := make(map[string]File)
	state := AppState{
		Files:      files,
		termHeight: height,
		termWidth:  width,
	}

	for _, fileName := range fileNames {
		state.Files[fileName] = File{}
	}

	viewNames := fileNames[:]
	for i := 0; i < 4; i++ {
		viewNames = append(viewNames, fileNames[i%len(fileNames)])
	}

	var views []LogView
	for _, fileName := range viewNames {
		views = append(views, LogView{FileName: fileName})
	}
	state.LogViews = views

	state.Categories = fileNames

	state.modifiers = modifiers{
		// highlighters
		modifier{buffer: buffer{"INFO"}, kind: highlighter, color: "yellow", active: true},
		modifier{buffer: buffer{"WARNING"}, kind: highlighter, color: "green", active: true},
		modifier{buffer: buffer{"500"}, kind: highlighter, color: "magenta", active: true},
		modifier{buffer: buffer{"^ "}, kind: highlighter, color: "blue"},
		modifier{buffer: buffer{".*"}, kind: highlighter, color: "cyan"},
		modifier{buffer: buffer{"WARNING"}, kind: filter, color: "white"},
		modifier{buffer: buffer{"^ "}, kind: filter, color: "white"},
	}

	state.layout = 1
	return state
}
