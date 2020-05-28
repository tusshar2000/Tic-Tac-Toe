package service

import (
	"Tic-Tac-Toe/components"
	"fmt"
)

type BoardService struct {
	*components.Board
}

func NewBoardService(size uint8) *BoardService {
	return &BoardService{components.NewBoard(size)}
}

func (bs *BoardService) PutMarkInPosition(p components.Player, position uint8) error {
	err := bs.Board.Cells[position].SetMark(p.Mark)
	return err
}

func (bs *BoardService) PrintBoard() string {
	matrixString := ""
	//bs.Board.Cells[position].SetMark(mark)
	boardSize := uint8(len(bs.Board.Cells))
	for i := uint8(0); i < (boardSize); i++ {
		if i%(bs.Board.Size) == 0 {
			matrixString += fmt.Sprint("\n\t")
		}
		matrixString += fmt.Sprint(bs.Board.Cells[i].GetMark(), " ")
	}
	return matrixString
}

func (bs *BoardService) CheckBoardIsFull() bool {
	boardSize := uint8(len(bs.Board.Cells))
	for i := uint8(0); i < (boardSize); i++ {
		if bs.Board.Cells[i].GetMark() == components.NoMark {
			return false
		}
	}
	return true
}
