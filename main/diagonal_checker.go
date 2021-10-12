package main

func isQueenInDiagonal(
	cellX int,
	cellY int,
	xDelta int,
	yDelta int,
	queens [] bool,
	boardSize int,
) bool {

	x := cellX + xDelta
	y := cellY + yDelta

	for x >= 0 && x < boardSize && y >= 0 && y < boardSize {
		diagCell := x * boardSize + y
		if queens[diagCell] {
			return true
		}
		x += xDelta
		y += yDelta
	}
	return false
}
