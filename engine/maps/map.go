package engine_maps

type TileChar int

const (
	EmptySpace = iota
	Wall
)

type MapTile struct {
	X    int
	Y    int
	Tile TileChar
}

type Map struct {
	tiles           *[][]TileChar
	tilesToUpdate   int
	UpdateTile      chan MapTile
	UpdateProcessed chan bool
}

func newMap(tiles *[][]TileChar, update chan MapTile) *Map {
	updateProcessed := make(chan bool, 1)
	mapp := Map{tiles: tiles, UpdateTile: update, UpdateProcessed: updateProcessed}

	return &mapp
}

func (mapp *Map) Tiles() [][]TileChar {
	tiles := *mapp.tiles
	return tiles
}

func (mapp *Map) StartUpdate(updateCount int) chan bool {
	mapp.tilesToUpdate = updateCount
	return mapp.UpdateProcessed
}

func (mapp *Map) UpdateTiles() {
	tiles := *mapp.tiles

	for update := range mapp.UpdateTile {
		tiles[update.X][update.Y] = update.Tile

		mapp.tilesToUpdate = mapp.tilesToUpdate - 1

		if mapp.tilesToUpdate == 0 {
			mapp.UpdateProcessed <- true
		}
	}
}
