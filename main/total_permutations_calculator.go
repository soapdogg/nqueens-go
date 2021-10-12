package main

func totalPermutations(boardSize int) int64 {
	totalCells := boardSize * boardSize
	result := int64(totalCells)
	for i := 1; i < boardSize; i++ {
		mult := int64(totalCells - i)
		result *= mult
	}
	return result
}
