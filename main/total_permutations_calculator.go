package main

func totalPermutations(boardSize int) uint64 {
	totalCells := boardSize * boardSize
	result := uint64(totalCells)
	for i := 1; i < boardSize; i++ {
		mult := uint64(totalCells - i)
		result *= mult
	}
	return result
}
