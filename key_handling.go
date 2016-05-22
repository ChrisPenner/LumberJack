package main

import "strconv"

// KeyPress sends a keypress
type KeyPress struct {
	Key string
}

// Apply the KeyPress
func (action KeyPress) Apply(state AppState, actions chan<- Action) AppState {
	key := action.Key
	switch state.CurrentMode {
	case normal:
		switch key {
		case "<space>":
			actions <- ChangeMode{Mode: selectCategory}
		case "<tab>":
			actions <- ChangeMode{Mode: filterMode}
			state.showFilters = true
		case "?", "/", "<enter>":
			actions <- ChangeMode{Mode: search}
		case "w":
			state.wrap = !state.wrap
		case "<backspace>":
			// Actually c-h
			actions <- ChangeSelection{Direction: left}
		case "C-l":
			actions <- ChangeSelection{Direction: right}
		case "<up>", "k":
			actions <- Scroll{Direction: up, NumLines: 1}
		case "<down>", "j":
			actions <- Scroll{Direction: down, NumLines: 1}
		case "b", "C-u":
			actions <- Scroll{Direction: up, NumLines: state.termHeight / 2}
		case "C-d":
			actions <- Scroll{Direction: down, NumLines: state.termHeight / 2}
		case "G":
			actions <- Scroll{Direction: bottom}
		case "n":
			actions <- findNext{direction: up}
		case "N":
			actions <- findNext{direction: down}
		case "f":
			state.showFilters = !state.showFilters
		case "!", "@", "#", "$", "%", "^", "&", "(", ")":
			actions <- toggleFilter{filter: numFromSymbol(key)}
		case "1", "2", "3", "4":
			choice, _ := strconv.Atoi(key)
			actions <- changeLayout{choice: choice}
		default:
			state.StatusBar.Text = key
		}
	case selectCategory:
		switch key {
		case "<enter>":
			bestMatch, ok := state.Categories.getBestMatch(state)
			if ok {
				actions <- SelectCategory{FileName: bestMatch}
			}
			fallthrough
		case "<escape>":
			actions <- ChangeMode{Mode: normal}
			state.selectCategoryBuffer.text = ""

		default:
			actions <- typeKey{key: key}
		}
	case search:
		switch key {
		case "<enter>":
			// quit search here...
			fallthrough
		case "<escape>":
			actions <- ChangeMode{Mode: normal}
		default:
			actions <- typeKey{key: convertKey(key)}
		}
	case filterMode:
		switch key {
		case "<tab>":
			actions <- ChangeMode{Mode: normal}
		case "<enter>":
			actions <- ChangeMode{Mode: editFilter}
		case "f":
			state.showFilters = !state.showFilters
		case "!", "@", "#", "$", "%", "^", "&", "(", ")":
			actions <- toggleFilter{filter: numFromSymbol(key)}
		case "j":
			if state.selectedFilter < len(state.filters)-1 {
				state.selectedFilter++
			}
		case "k":
			if state.selectedFilter > 0 {
				state.selectedFilter--
			}
		}
	case editFilter:
		switch key {
		case "<enter>":
			actions <- ChangeMode{Mode: filterMode}
		default:
			actions <- typeKey{key: key}
		}
	default:
		panic("Didn't handle keypress!")
	}
	return state
}

func numFromSymbol(key string) int {
	switch key {
	case "!":
		return 0
	case "@":
		return 1
	case "#":
		return 2
	case "$":
		return 3
	case "%":
		return 4
	case "^":
		return 5
	case "&":
		return 6
	case "*":
		return 7
	case "(":
		return 8
	case ")":
		return 9
	}
	return 0
}
