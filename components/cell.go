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

func Mark(cell *Cell, mark string) bool {
	if cell.mark == NoMark {
		cell.mark = mark
		return true
	}
	return false
}

func GetMark(cell *Cell) string {
	return cell.mark
}
