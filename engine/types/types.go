package types

type Tiles int

const (
	EmptySpace = iota
	Wall
)

type MapTile struct {
	X    int
	Y    int
	Tile Tiles
}
