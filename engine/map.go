package engine

type Map struct {
	Tiles [][]int
}

const (
	Wall = iota
	EmptySpace
)
