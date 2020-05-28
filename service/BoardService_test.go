package service

import (
	"Tic-Tac-Toe/components"
	"errors"
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
		got := NewBoardService(want.size)
		gotMatrixSize := uint8(len(got.Cells))
		if want.matrixSize != gotMatrixSize {
			t.Error(want.matrixSize, gotMatrixSize)
		}
		for _, cell := range got.Cells {
			gotMark := cell.GetMark()
			if want.mark != gotMark {
				t.Error(want.mark, gotMark)
			}

		}
	}

}

func TestPutMarkInPosition(t *testing.T) {
	tests := []struct {
		wantBoard    BoardService
		wantPlayer   components.Player
		wantPosition uint8
		wantError    error
	}{
		{*NewBoardService(2), components.Player{Name: "Tusshar", Mark: components.XMark}, 1, errors.New("Cell is already marked.")},
		{BoardService{&components.Board{
			Cells: []*components.Cell{
				&components.Cell{Mark: components.XMark},
				&components.Cell{Mark: components.XMark},
				&components.Cell{Mark: components.OMark},
				&components.Cell{Mark: components.NoMark},
			},
			Size: 2,
		},
		}, components.Player{Name: "Tusshar", Mark: components.OMark}, 1, errors.New("Cell is already marked.")},
	}

	for _, test := range tests {
		got := test.wantBoard.PutMarkInPosition(test.wantPlayer, test.wantPosition)
		if got != nil {
			if got.Error() != test.wantError.Error() {
				t.Error(got, test.wantError)
			}
		}
	}
}

func TestPrintBoard(t *testing.T) {
	tests := []struct {
		wantBoard        *BoardService
		wantMatrixString string
	}{
		{NewBoardService(4), "\n\t- - - - \n\t- - - - \n\t- - - - \n\t- - - - "},
		{&BoardService{&components.Board{
			Cells: []*components.Cell{
				&components.Cell{Mark: components.NoMark},
				&components.Cell{Mark: components.NoMark},
				&components.Cell{Mark: components.NoMark},
				&components.Cell{Mark: components.XMark},
			},
			Size: 2,
		},
		}, "\n\t- - \n\t- X "},
	}
	for _, test := range tests {
		got := test.wantBoard.PrintBoard()
		if test.wantMatrixString != got {
			t.Error(test.wantMatrixString, got)
		}
	}
}

func TestCheckBoardIsFull(t *testing.T) {
	tests := []struct {
		wantBoard *BoardService
		want      bool
	}{
		{NewBoardService(3), false},
		{&BoardService{&components.Board{
			Cells: []*components.Cell{
				&components.Cell{Mark: components.NoMark},
				&components.Cell{Mark: components.XMark},
				&components.Cell{Mark: components.OMark},
				&components.Cell{Mark: components.XMark},
			},
		},
		}, false},
		{&BoardService{&components.Board{
			Cells: []*components.Cell{
				&components.Cell{Mark: components.OMark},
				&components.Cell{Mark: components.XMark},
				&components.Cell{Mark: components.OMark},
				&components.Cell{Mark: components.XMark},
			},
		},
		}, true},
	}
	for _, test := range tests {
		got := test.wantBoard.CheckBoardIsFull()
		if test.want != got {
			t.Error(test.want, got)
		}
	}
}
