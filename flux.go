package main

// Store for Flux
type Store struct {
	Actions chan Action
}

// NewStore is the Store constructor
func NewStore() *Store {
	store := Store{}
	store.Actions = make(chan Action, 10)
	return &store
}

// ReduceLoop will continually apply actions to state
func (store Store) ReduceLoop(state *AppState) {
	for {
		action := <-store.Actions
		action.Apply(state)
		Render(state)
	}
}

// Action represents a change to take place
type Action interface {
	Apply(*AppState)
}

// type nullAction struct{}

// func (action nullAction) Apply(state *AppState) {}

type render struct{}

func (action render) Apply(state *AppState) {}
