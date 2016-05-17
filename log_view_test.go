package main

import "testing"

func TestGetFileSliceInRange(t *testing.T) {
	file := File{"1", "2", "3", "4", "5", "6"}
	view := LogView{offSet: 2}
	slice := file.getVisibleSlice(view, 3)
	if len(slice) != 3 || slice[0] != "2" {
		t.Fail()
	}
}

func TestGetFileSlicePastStart(t *testing.T) {
	file := File{"1", "2", "3", "4", "5", "6"}
	view := LogView{offSet: 2}
	slice := file.getVisibleSlice(view, 6)
	if len(slice) != 6 || slice[0] != "1" {
		t.Fail()
	}
}

func TestGetFileSliceMoreVisibleThanLines(t *testing.T) {
	file := File{"1", "2"}
	view := LogView{offSet: 0}
	slice := file.getVisibleSlice(view, 6)
	if len(slice) != 2 || slice[0] != "1" {
		t.Fail()
	}
}
