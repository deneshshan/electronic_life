package state

import (
	types "github.com/deneshshan/electronic_life/engine/types"
)

type IStateReader interface {
	Read() [][]types.MapTile
}

// State represents the game state. Currently it is only a tile map of all obstacles
// and entity positions. Only one State can exist per Engine instance.
type State struct {
	tiles  *[][]types.MapTile
	width  int
	height int
}

func NewState(width, height int) *State {

	state := State{width: width, height: height}
	tiles := state.blankTileMap()
	state.tiles = &tiles

	return &state
}

func (st *State) Read() [][]types.MapTile {
	snapshot := st.blankTileMap()

	copy(snapshot, *st.tiles)

	return snapshot
}

func (st *State) blankTileMap() [][]types.MapTile {
	tiles := make([][]types.MapTile, st.height)

	for row := range tiles {
		tiles[row] = make([]types.MapTile, st.width)
	}

	return tiles
}

//}

// TODO: implement in state reader
//func (state *State) Tiles() [][]MapTile {
//tiles := *state.tiles
//return tiles
//}

//func (state *State) GetTiles(tileType TileChar) []MapTile {
//var tilesOfType []MapTile
//updateSlice := make(chan MapTile, 200)

//tiles := *state.tiles

//var wg sync.WaitGroup

//for _, column := range tiles {
//for _, tile := range column {
//wg.Add(1)

//go func(tile MapTile) {
//if tile.Tile == tileType {
//updateSlice <- tile
//}
//}(tile)
//}
//}

//go func() {
//for tile := range updateSlice {
//tilesOfType = append(tilesOfType, tile)
//wg.Done()
//}
//}()

//wg.Wait()

//close(updateSlice)

//return tilesOfType
//}

// TODO: implement in state writer
//func (state *State) StartUpdate(updateCount int) chan bool {
//state.setUpdateCount(updateCount)
//return state.UpdateProcessed
//}

// TODO: implement in state writer
//func (state *State) UpdateTiles() {
//tiles := *state.tiles

//for update := range state.UpdateTile {
//tiles[update.X][update.Y].Tile = update.Tile

//state.setUpdateCount(-1)

//if state.tilesToUpdate == 0 {
//state.UpdateProcessed <- true
//}
//}
//}

// TODO: implement in state writer
//func (state *State) setUpdateCount(value int) {
//var mu sync.Mutex

//mu.Lock()
//state.tilesToUpdate = state.tilesToUpdate + value
//mu.Unlock()
//}
