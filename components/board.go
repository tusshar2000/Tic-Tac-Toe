package components

type Board struct {
	Cells []*Cell
	Size  uint8
}

func NewBoard(size uint8) *Board {
	BoardSize := size * size
	var cellArray = make([]*Cell, BoardSize)
	for i := range cellArray {
		cellArray[i] = NewCell()
	}
	return &Board{
		Cells: cellArray,
		Size:  size,
	}
}
