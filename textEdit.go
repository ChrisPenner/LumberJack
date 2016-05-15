package main

// Backspace action
type Backspace struct {
	Buffer *TextBuffer
}

// Apply the Backspace
func (action Backspace) Apply(state *AppState) {
	text := action.Buffer.Text
	if len(text) > 0 {
		text = text[:len(text)-1]
	}
	action.Buffer.Text = text
}

// TypeKey types a key
type TypeKey struct {
	Key    string
	Buffer *TextBuffer
}

// Apply the Keystroke
func (action TypeKey) Apply(state *AppState) {
	action.Buffer.Text = action.Buffer.Text + action.Key
}

// TextBuffer provides an abstraction over editing text
type TextBuffer struct {
	Cursor int
	Text   string
}
