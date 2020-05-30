package service

import (
	_ "fmt"
)

type ResultService struct {
	*BoardService
}

func NewResultService(bs *BoardService) *ResultService {
	return &ResultService{bs}
}

func (rs *ResultService) checkRow(mark string) bool {
	count := 0
	for i := uint8(0); i < rs.Size; i++ {
	}
	for i := 0; i < int(rs.Size*rs.Size); i++ {
		if rs.Board.Cells[i].GetMark() == mark {
			count++
		}
		if count == int(rs.Size) {
			return true
		}
		if (i+1)%int(rs.Size) == 0 {
			count = 0
		}
	}
	return false
}

func (rs *ResultService) checkColumn(mark string) bool {
	for i := uint8(0); i < rs.Size; i++ {
		count := 0
		for j := i; j <= (rs.Size*rs.Size)-(rs.Size-i); j += rs.Size {
			if rs.Board.Cells[j].GetMark() == mark {
				count++
			}
			//fmt.Println(j, count, count == int(rs.Size))
			if count == int(rs.Size) {
				return true
			}
		}
	}
	return false
}

func (rs *ResultService) checkFirstDiagonal(mark string) bool {
	size := rs.Size
	for i := uint8(0); i < size; i++ {
		if rs.Board.Cells[size*i+i].GetMark() != mark {
			return false
			// fmt.Println(flag)
		}
	}
	// fmt.Println(flag)
	return true
}

func (rs *ResultService) checkSecondDiagonal(mark string) bool {
	size := rs.Size
	for i := uint8(0); i < size; i++ {
		if rs.Board.Cells[(size*i)+(size-1-i)].GetMark() != mark {
			return false
		}
	}
	return true
}
