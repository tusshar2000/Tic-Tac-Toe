package components

import (
	"testing"
)

func TestCreateBoard(t *testing.T) {
	test := CreateBoard()
	expected := make([]Cell, 9)
	if len(test.Cells) != len(expected) {
		t.Error(test.Cells, expected)
	}
}
