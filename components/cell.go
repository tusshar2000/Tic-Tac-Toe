package components

import (
	"errors"
)

type Cell struct {
	Mark string
}

const (
	NoMark = "-"
	XMark  = "X"
	OMark  = "O"
)

func NewCell() *Cell {
	return &Cell{Mark: NoMark}
}

func (cell *Cell) SetMark(mark string) (err error) {
	if cell.Mark == NoMark {
		cell.Mark = mark
		return nil
	}
	return errors.New("Cell is already marked.")
}

func (cell *Cell) GetMark() string {
	return cell.Mark
}
