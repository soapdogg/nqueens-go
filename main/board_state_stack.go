package main

type boardStack struct {
	 elements []*board
}

func NewBoardStateStack() *boardStack {
	return &boardStack{}
}

func (b *boardStack) Push(state *board) {
	b.elements = append(b.elements, state)
}

func (b *boardStack) Pop() *board {
	lastIndex := len(b.elements) - 1
	res := b.elements[lastIndex]
	b.elements = b.elements[:lastIndex]
	return res
}

func (b *boardStack) IsNotEmpty() bool {
	return len(b.elements) != 0
}
