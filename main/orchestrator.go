package main

import "fmt"

func orchestrate(boardSize int) {

	totalPermutations := totalPermutations(boardSize)
	boardStateStack := NewBoardStateStack()

	initialBoard := map[int]bool{}

	uniqueTerminalBoardCount := 0
	var hashes = map[string]bool{}
	initialBoardState := NewBoard(initialBoard, 0, hashes)
	boardStateStack.Push(initialBoardState)

	totalTerminalBoards := 0
	totalPrunedBoards := 0
	totalBoardsProcesses := 0

	for boardStateStack.IsNotEmpty() {
		boardState := boardStateStack.Pop()
		totalBoardsProcesses++

		if totalBoardsProcesses%100*boardSize == 0 {
			fmt.Println("Processed ", totalBoardsProcesses, "/", totalPermutations, " boards; found ", totalTerminalBoards, " terminals, pruned ", totalPrunedBoards)
		}

		if boardState.size == boardSize {
			uniqueTerminalBoardCount++
			totalTerminalBoards += len(boardState.GetHashes())
			continue
		}

		childBoardStates := generateChildBoardStates(boardState, boardSize)
		for _, childBoardState := range childBoardStates {

			shouldPrune := false
			for hash, _ := range childBoardState.hashes {
				if hashes[hash] {
					totalPrunedBoards++
					shouldPrune = true
					break
				}
			}
			if shouldPrune {
				continue
			}

			for hash, _ := range childBoardState.hashes {
				hashes[hash] = true
			}
			boardStateStack.Push(childBoardState)
		}
	}

	fmt.Println("Max board permutations: ", totalPermutations)
	fmt.Println("Total terminal boards: ", totalTerminalBoards)
	fmt.Println("Unique terminal boards identified: ", uniqueTerminalBoardCount)
	fmt.Println("Total boards processed: ", totalBoardsProcesses)
	fmt.Println("Total boards pruned: ", totalPrunedBoards)

}
