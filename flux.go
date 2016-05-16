package main

import "time"

const renderInterval = 50 * time.Millisecond

// Store for Flux
type Store struct {
	Actions chan Action
}

// NewStore is the Store constructor
func NewStore() Store {
	store := Store{}
	store.Actions = make(chan Action, 100)
	return store
}

// ReduceLoop will continually apply actions to state
func (store Store) ReduceLoop(state AppState) {
	// This debouncing logic allows us to keep applying state changes, but only render every renderInterval
	debouncer := time.After(renderInterval)
	rendered := true
	for {
		if rendered {
			action := <-store.Actions
			state = action.Apply(state, store.Actions)
			rendered = false
		}
		select {
		case <-debouncer:
			Render(state)
			debouncer = time.After(renderInterval)
			rendered = true
		case action := <-store.Actions:
			state = action.Apply(state, store.Actions)
			rendered = false
		}
	}
}

// Action represents a change to take place
type Action interface {
	Apply(AppState, chan<- Action) AppState
}

type render struct{}

func (action render) Apply(state AppState, actions chan<- Action) AppState {
	return state
}
