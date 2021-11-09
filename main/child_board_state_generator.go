package main

func generateChildBoardStates(
	b *board,
	boardSize int,
) []*board {
	result := []*board{}
	for i := 0; i < boardSize*boardSize; i++ {
		isValidQueenPos := isQueenPositionValid(i, boardSize, b.queenPositions)

		if isValidQueenPos {
			c := map[int]bool{}
			for key, value := range b.queenPositions {
				c[key] = value
			}
			c[i] = true

			var allQueens []map[int]bool
			mirroredQueens := getAllMirrors(c, boardSize)
			for _, mirroredQueen := range mirroredQueens {
				rotation := getAllRotations(mirroredQueen, boardSize)
				allQueens = append(allQueens, rotation...)
			}

			allQueens2 := [][]bool{}
			for j := 0; j < len(allQueens); j++ {
				queens := allQueens[j]
				queen2 := make([]bool, boardSize*boardSize)
				for i := range queens {
					queen2[i] = true
				}
				allQueens2 = append(allQueens2, queen2)
			}

			hashes := generateHashStrings(allQueens2)
			childBoard := NewBoard(
				c,
				b.size+1,
				hashes,
			)
			result = append(result, childBoard)
		}
	}
	return result
}
