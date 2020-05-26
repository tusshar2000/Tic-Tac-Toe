package components

import (
	"errors"
)

type Cell struct {
	mark string
}

const (
	NoMark = " "
	XMark  = "X"
	OMark  = "O"
)

func NewCell() *Cell {
	return &Cell{mark: NoMark}
}

func (cell *Cell) SetMark(mark string) (bool, error) {
	if cell.mark == NoMark {
		cell.mark = mark
		return true, nil
	}
	return false, errors.New("Cell is already marked.")
}

func (cell *Cell) GetMark() string {
	return cell.mark
}
