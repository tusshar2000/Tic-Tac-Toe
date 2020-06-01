package service

import (
	_ "fmt"
)

type ResultService struct {
	*BoardService
}

type Result struct {
	Win  bool
	Draw bool
}

func NewResultService(bs *BoardService) *ResultService {
	return &ResultService{bs}
}

func (rs *ResultService) CheckRow(mark string) bool {
	size := rs.Size
	count := 0
	for i := uint8(0); i < size*size; i++ {
		if rs.Board.Cells[i].GetMark() == mark {
			count++
		}
		if count == int(size) {
			return true
		}
		if (i+1)%(size) == 0 {
			count = 0
		}
	}
	return false
}

func (rs *ResultService) CheckColumn(mark string) bool {
	size := rs.Size
	for i := uint8(0); i < size; i++ {
		count := 0
		for j := i; j <= (size*size)-(size-i); j += size {
			if rs.Board.Cells[j].GetMark() == mark {
				count++
			}
			//fmt.Println(j, count, count == int(size))
			if count == int(size) {
				return true
			}
		}
	}
	return false
}

func (rs *ResultService) CheckFirstDiagonal(mark string) bool {
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

func (rs *ResultService) CheckSecondDiagonal(mark string) bool {
	size := rs.Size
	for i := uint8(0); i < size; i++ {
		if rs.Board.Cells[(size*i)+(size-1-i)].GetMark() != mark {
			return false
		}
	}
	return true
}

func (rs *ResultService) GetResult(mark string) Result {
	if rs.CheckRow(mark) {
		return Result{true, false}
	} else if rs.CheckColumn(mark) {
		return Result{true, false}
	} else if rs.CheckFirstDiagonal(mark) {
		return Result{true, false}
	} else if rs.CheckSecondDiagonal(mark) {
		return Result{true, false}
	} else if rs.CheckBoardIsFull() {
		return Result{false, true}
	}
	return Result{false, false}

}
