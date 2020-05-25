// package components

// type Board struct {
// 	Cells []Cell
// }

// func NewBoard() *Board {
// 	myBoard := CreateBoard()
// 	for i := range myBoard.Cells {
// 		myBoard.Cells[i] = myBoard.NewCell()
// 	}
// 	return &myBoard
// }

// func CreateBoard() *Board {
// 	rows := 3 //can take input for this
// 	new = make(Board, rows)
// 	return &new
// }
package components

type Board struct {
	Cells []Cell
}

func NewBoard() *Board {
	//matrixSize := 9
	new := CreateBoard()
	for i := range new.Cells {
		new.Cells[i] = *NewCell()
	}
	return new
}

func CreateBoard() *Board {
	new := Board{}
	// if matrixSize <= 0 {
	// 	nil,fmt.Println("Matrix needs a positive number.")
	// }
	new.Cells = make([]Cell, 9)
	return &new
}
