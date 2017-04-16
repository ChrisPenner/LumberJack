package main

import "sort"

// AppState contains global state
type AppState struct {
	termHeight           int
	termWidth            int
	CurrentMode          mode
	LogViews             LogViews
	Files                Files
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
	orderedFileNames     []string
}

// NewAppState constructs and appstate
func NewAppState(fileNames []string, height int, width int) *AppState {
	sort.Strings(fileNames)
	files := make(map[string]file)
	state := AppState{
		Files:            files,
		orderedFileNames: fileNames,
		termHeight:       height,
		termWidth:        width,
	}

	for _, fileName := range fileNames {
		state.Files[fileName] = newFile(lines{})
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

	state.modifiers = modifiers{
		// highlighters
		modifier{buffer: buffer{"INFO"}, kind: highlighter, bgColor: "green", fgColor: "black", active: true},
		modifier{buffer: buffer{"WARNING"}, kind: highlighter, bgColor: "yellow", fgColor: "black", active: true},
		modifier{buffer: buffer{"500"}, kind: highlighter, bgColor: "magenta", fgColor: "black", active: true},
		modifier{buffer: buffer{"^ +"}, kind: highlighter, bgColor: "blue", fgColor: "black"},
		modifier{buffer: buffer{""}, kind: highlighter, bgColor: "cyan", fgColor: "black"},
		modifier{buffer: buffer{"WARNING"}, kind: filter, bgColor: "white", fgColor: "black"},
		modifier{buffer: buffer{"^ "}, kind: filter, bgColor: "white", fgColor: "black"},
	}

	state.layout = 1
	return &state
}
