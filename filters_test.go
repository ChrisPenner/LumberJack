package main

import "testing"

func TestFiltering(t *testing.T) {
	file := File{"one", "twox", "three", "xfour", "fivex"}
	filters := []filter{
		filter{textBuffer: textBuffer{text: "x"}, active: true},
	}
	filtered := file.filter(filters, 2, 1)
	expected := File{"twox", "xfour", "fivex"}
	if len(filtered) != 3 || filtered[0] != expected[0] || filtered[1] != expected[1] || filtered[2] != expected[2] {
		t.Error(filtered)
	}
}
