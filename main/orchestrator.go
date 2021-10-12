package main

import "fmt"

func orchestrate(boardSize int) {

	totalPermutations := totalPermutations(boardSize)
	boardStateStack := NewBoardStateStack()

	initalBoard :=  make([]bool, boardSize * boardSize)

	var terminalBoardStates [][]bool

	allQueens := [][]bool{}
	mirroredQueens := getAllMirrors(initalBoard, boardSize)
	allQueens = append(allQueens, mirroredQueens...)
	for _, mirroredQueen := range mirroredQueens {
		rotation := getAllRotations(mirroredQueen, boardSize)
		allQueens = append(allQueens, rotation...)
	}

	hashes := generateHashStrings(allQueens)

	initialBoardState := NewBoard(initalBoard, 0, hashes)
	boardStateStack.Push(initialBoardState)

	totalTerminalBoards := 0
	totalPrunedBoards := 0
	totalBoardsProcesses := 0

	for boardStateStack.IsNotEmpty() {
		boardState := boardStateStack.Pop()
		totalBoardsProcesses++

		if totalBoardsProcesses % 100 * boardSize == 0 {
			fmt.Println("Processed ", totalBoardsProcesses, "/", totalPermutations, " boards; found ", totalTerminalBoards, " terminals, pruned ", totalPrunedBoards)
		}

		if boardState.size == boardSize {
			terminalBoardStates = append(terminalBoardStates, boardState.queenPositions)
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
	fmt.Println("Unique terminal boards identified: ", len(terminalBoardStates))
	fmt.Println("Total boards processed: ", totalBoardsProcesses)
	fmt.Println("Total boards pruned: ", totalPrunedBoards)


}
