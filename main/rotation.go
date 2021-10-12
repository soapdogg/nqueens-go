package main

var ROTATIONS = [][][]float32 {
	{{0.0, -1.0}, {1.0, 0.0}},
	{{-1.0, 0.0}, {0.0, -1.0}},
	{{0.0, 1.0}, {-1.0, 0.0}},
}

func getAllRotations(queens []bool, boardSize int) [][] bool {
	result := [][]bool{queens}
	for i := 0; i <= 2; i++ {
		rotated := rotateNinetyDegrees(queens, i, boardSize)
		result = append(result, rotated)
	}
	return result
}

func rotateNinetyDegrees(queens []bool, times int, boardSize int) []bool {
	center := float32(boardSize - 1) / 2
	rotation := ROTATIONS[times]
	x := rotation[0]
	y := rotation[1]
	result := make([]bool, boardSize * boardSize)

	for i := 0; i < len(queens); i++ {
		if queens[i] {
			cellX := i / boardSize
			cellY := i % boardSize
			baseX := float32(cellX) - center
			baseY := float32(cellY) - center
			rotatedX := baseX * x[0] + baseY * y[0] + center
			rotatedY := baseX * x[1] + baseY * y[1] + center
			rotatedCell := int(rotatedX) * boardSize + int(rotatedY)
			result[rotatedCell] = true
		}
	}

	return result
}
