package main

func printBoard(queenPositions []bool, boardSize int) {
	for i := 0; i < boardSize; i++ {
		for j := 0; j < boardSize; j++ {
			cell := (i *boardSize) + j
			if queenPositions[cell] {
				print('Q')
			} else {
				print('x')
			}
		}
		println()
	}
	println()
}
