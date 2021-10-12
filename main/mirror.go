package main

func getAllMirrors(queens []bool, boardSize int) [][]bool {
	mirroredQueens0 := mirror(queens, false, false, boardSize)
	mirroredQueens1 := mirror(queens, true, false, boardSize)
	mirroredQueens2 := mirror(queens, false, true, boardSize)
	mirroredQueens3 := mirror(queens, true, true, boardSize)

	return [][]bool {mirroredQueens0, mirroredQueens1, mirroredQueens2, mirroredQueens3}
}

func mirror(queens []bool, horizontally bool, vertically bool, boardSize int) []bool {
	resultSize := boardSize *boardSize
	result := make([]bool, resultSize)

	for i := 0; i < resultSize; i++ {
		if queens[i] {
			queenX := i / boardSize
			queenY := i % boardSize

			var newQueenX int
			if horizontally {
				newQueenX = boardSize - 1 - queenX
			} else {
				newQueenX = queenX
			}

			var newQueenY int
			if vertically {
				newQueenY = boardSize - 1 - queenY
			} else {
				newQueenY = queenY
			}

			result[newQueenX * boardSize + newQueenY] = true
		}
	}

	return result
}