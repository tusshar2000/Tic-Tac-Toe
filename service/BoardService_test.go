package service

import (
	"Tic-Tac-Toe/components"
	"testing"
)

func TestNewBoardService(t *testing.T) {
	tests := []struct {
		size       uint8
		matrixSize uint8
		mark       string
	}{
		{0, 0, components.NoMark},
		{3, 9, components.NoMark},
		//-1 uint8 overflow, handled by datatype.
	}
	for _, want := range tests {
		got := components.NewBoard(want.size).Cells
		gotMatrixSize := uint8(len(got))
		if want.matrixSize != gotMatrixSize {
			t.Error(want.matrixSize, gotMatrixSize)
		}
		for _, cell := range got {
			gotMark := cell.GetMark()
			if want.mark != gotMark {
				t.Error(want.mark, gotMark)
			}

		}
	}

}

// func TestPutMarkInPosition(t *testing.T){
// 	tests := []struct {
// 		player     components.Player
// 		position
// 		mark       string
// 	}
// }
