package components

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

func (cell *Cell) SetMark(mark string) bool {
	if cell.mark == NoMark {
		cell.mark = mark
		return true
	}
	return false
}

func (cell *Cell) GetMark() string {
	return cell.mark
}
