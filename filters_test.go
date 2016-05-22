package main

import "testing"

func TestFiltering(t *testing.T) {
	file := File{"one", "twox", "three", "xfour"}
	filters := []filter{
		filter{textBuffer: textBuffer{text: "x"}, active: true},
	}
	filtered := file.filter(filters)
	expected := File{"twox", "xfour"}
	if len(filtered) != 2 || filtered[0] != expected[0] || filtered[1] != expected[1] {
		t.Error(filtered)
	}
}
