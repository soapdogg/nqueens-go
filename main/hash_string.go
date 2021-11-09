package main

import (
	"sort"
	"strconv"
)

func generateHashStrings(mirroredAndRotatedBoards []map[int]bool) map[string]bool {
	result := map[string]bool{}
	for i := 0; i < len(mirroredAndRotatedBoards); i++ {
		hash := generateHashString(mirroredAndRotatedBoards[i])
		result[hash] = true
	}
	return result
}

func generateHashString(queens map[int]bool) string {
	q := make([]int, 0, len(queens))
	for i := range queens {
		q = append(q, i)
	}
	sort.Ints(q)
	result := ""
	for _, i := range q {
		result += strconv.Itoa(i)
		result += "-"
	}
	return result
}
