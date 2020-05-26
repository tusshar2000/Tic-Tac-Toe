package service

import (
	"components"
	"fmt"
)

type BoardService struct {
	*components.Board
}

func NewBoardService(size uint8) *BoardService {
	return &BoardService{components.NewBoard(size)}
}

func (bs *BoardService) Mark(position int, mark string) {
	matrixString := ""
	bs.Board.Cells[position].SetMark(mark)
	for i := 0; i < (bs.Board.size * bs.Board.size); i++ {
		if i % (bs.Board.size) == 0 {
			matrixString += fmt.Sprint("\n\t")
		}
		matrixString  += fmt.Sprint(bs.Board.Cells[i].GetMark()," ")
	}
	return matrixString
}
