package main

import "strconv"

// KeyPress sends a keypress
type KeyPress struct {
	Key string
}

// Apply the KeyPress
func (action KeyPress) Apply(state AppState) AppState {
	key := action.Key
	switch state.CurrentMode {
	case normal:
		switch key {
		case "<enter>":
			state = state.changeMode(selectCategory)
		case "<tab>":
			state.showFilters = !state.showFilters
		case "?", "/":
			state = state.changeMode(search)
		case "w":
			state.wrap = !state.wrap
		case "<backspace>":
			// Actually c-h
			state = state.changeSelection(left)
		case "C-l":
			state = state.changeSelection(right)
		case "<up>", "k":
			state = state.scroll(up, 1)
		case "<down>", "j":
			state = state.scroll(down, 1)
		case "b", "C-u":
			state = state.scroll(up, state.termHeight/2)
		case "C-d":
			state = state.scroll(down, state.termHeight/2)
		case "G":
			state = state.scroll(bottom, 0)
		case "n":
			state = state.findNext(up)
		case "N":
			state = state.findNext(down)
		case "!", "@", "#", "$", "%", "^", "&", "(", ")":
			state = state.toggleFilter(numFromSymbol(key))
		case "1", "2", "3", "4":
			choice, _ := strconv.Atoi(key)
			state = state.changeLayout(choice)
		default:
			state.StatusBar.Text = key
		}
	case selectCategory:
		switch key {
		case "<enter>":
			bestMatch, ok := state.Categories.getBestMatch(state)
			if ok {
				state = state.selectCategory(bestMatch)
			}
			fallthrough
		case "<escape>":
			state = state.changeMode(normal)
			state.selectCategoryBuffer.text = ""

		default:
			state = state.typeKey(key)
		}
	case search:
		switch key {
		case "<enter>":
			// quit search here...
			fallthrough
		case "<escape>":
			state = state.changeMode(normal)
		default:
			state = state.typeKey(key)
		}
	case filterMode:
		switch key {
		case "<tab>":
			state.showFilters = false
			state = state.changeMode(normal)
		case "<enter>":
			state = state.changeMode(editFilter)
		case "<backspace>":
			state.selected = state.layout - 1
			state = state.changeMode(normal)
		case "!", "@", "#", "$", "%", "^", "&", "(", ")":
			state = state.toggleFilter(numFromSymbol(key))
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
			state = state.changeMode(filterMode)
		default:
			state = state.typeKey(key)
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
