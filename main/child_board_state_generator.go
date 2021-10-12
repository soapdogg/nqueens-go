package main

func generateChildBoardStates(
	b * board,
	boardSize int,
) []*board {
	result := []*board{}
	for i := 0; i < len(b.queenPositions); i++ {
		isValidQueenPos := isQueenPositionValid(i, boardSize, b.queenPositions)

		if isValidQueenPos {
		 	c := make([]bool, boardSize*boardSize)
		    for j := 0; j < boardSize * boardSize; j++ {
				if b.queenPositions[j] || j == i {
					c[j] = true
				}
			}
			allQueens := [][]bool{}
			mirroredQueens := getAllMirrors(c, boardSize)
			//allQueens = append(allQueens, mirroredQueens...)
			for _, mirroredQueen := range mirroredQueens {
				rotation := getAllRotations(mirroredQueen, boardSize)
				allQueens = append(allQueens, rotation...)
			}

			hashes := generateHashStrings(allQueens)
			childBoard := NewBoard(
				c,
				b.size + 1,
				hashes,
			)
			result = append(result, childBoard)
		}
	}
	return result
}
