package main

import "testing"

func TestBackSpace(t *testing.T) {
	tb := textBuffer{text: "text"}
	tb = tb.typeKey("<BS>")
	if tb.text != "tex" {
		t.Fail()
	}
}
func TestBackSpaceOnEmptyString(t *testing.T) {
	tb := textBuffer{text: ""}
	tb = tb.typeKey("<BS>")
	if tb.text != "" {
		t.Fail()
	}
}

func TestTypeKey(t *testing.T) {
	tb := textBuffer{text: ""}
	tb = tb.typeKey("a")
	if tb.text != "a" {
		t.Fail()
	}
	tb = tb.typeKey("b")
	if tb.text != "ab" {
		t.Fail()
	}
}

func TestConvertKey(t *testing.T) {
	if convertKey("<space>") != " " {
		t.Fail()
	}
	if convertKey("a") != "a" {
		t.Fail()
	}
	if convertKey("C-j") != "" {
		t.Fail()
	}
}
