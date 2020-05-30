package service

import (
	"Tic-Tac-Toe/components"
	"errors"
	"testing"
)

func TestPutMarkInPosition(t *testing.T) {
	tests := []struct {
		wantBoard    *BoardService
		wantPlayer   components.Player
		wantPosition uint8
		wantError    error
	}{
		{&BoardService{components.NewBoard(uint8(3))}, components.Player{Name: "Tusshar", Mark: components.XMark}, 1, nil},
		{&BoardService{&components.Board{
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
		if test.wantError != nil || got != nil {
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
		{&BoardService{components.NewBoard(uint8(4))}, "\n\t- - - - \n\t- - - - \n\t- - - - \n\t- - - - "},
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
		{&BoardService{components.NewBoard(uint8(3))}, false},
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
