package main

import "strconv"

func generateHashStrings(mirroredAndRotatedBoards [][]bool) map[string] bool{
	result := map[string]bool{}
	for i := 0; i < len(mirroredAndRotatedBoards); i++ {
		hash := generateHashString(mirroredAndRotatedBoards[i])
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