package main

func isQueenPositionValid(
	cell int,
	boardSize int,
	queens [] bool,
) bool {
	if queens[cell] {
		return false
	}

	cellX := cell / boardSize
	cellY := cell % boardSize

	for i := 0; i < boardSize; i++ {
		horizontalCell := i * boardSize + cellY
		verticalCell := cellX * boardSize + i
		if queens[horizontalCell] || queens[verticalCell] {
			return false
		}
	}

	if isQueenInDiagonal(cellX, cellY, -1, -1, queens, boardSize) ||
		isQueenInDiagonal(cellX, cellY, 1, -1, queens, boardSize) ||
		isQueenInDiagonal(cellX, cellY, -1, 1, queens, boardSize) ||
		isQueenInDiagonal(cellX, cellY, 1, 1, queens, boardSize) {
		return false
	}

	return true
}
