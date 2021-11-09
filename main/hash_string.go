package main

import "strconv"

func generateHashStrings(mirroredAndRotatedBoards []map[int]bool, boardSize int) map[string]bool {
	result := map[string]bool{}
	for i := 0; i < len(mirroredAndRotatedBoards); i++ {
		board := make([]bool, boardSize*boardSize)
		b := mirroredAndRotatedBoards[i]
		for j := range b {
			board[j] = true
		}
		hash := generateHashString(board)
		result[hash] = true
	}
	return result
}

func generateHashString(queens []bool) string {
	result := ""
	for i := 0; i < len(queens); i++ {
		if queens[i] {
			result += strconv.Itoa(i)
			result += "-"
		}
	}
	return result
}
