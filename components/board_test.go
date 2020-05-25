package components

import (
	"testing"
)

func TestCreateBoard(t *testing.T) {
	want := make([]Cell, 9)
	got := CreateBoard()
	if len(want) != len(got.Cells) {
		t.Error(want,got.Cells)
	}
}
