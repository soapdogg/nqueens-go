package main

type board struct {
	queenPositions map[int]bool
	size           int
	hashes         map[string]bool
}

func NewBoard(
	queenPositions map[int]bool,
	size int,
	hashes map[string]bool,
) *board {
	return &board{
		queenPositions,
		size,
		hashes,
	}
}

func (b *board) GetQueenPositions() map[int]bool {
	return b.queenPositions
}

func (b *board) GetSize() int {
	return b.size
}

func (b *board) GetHashes() map[string]bool {
	return b.hashes
}
