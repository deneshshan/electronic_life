package engine

import (
	"sync"
)

type State struct {
	tiles           *[][]MapTile
	tilesToUpdate   int
	UpdateTile      chan MapTile
	UpdateProcessed chan bool
}

func newState(tiles *[][]MapTile, update chan MapTile) *State {
	updateProcessed := make(chan bool, 1)
	state := State{tiles: tiles, UpdateTile: update, UpdateProcessed: updateProcessed}

	return &state
}

func (state *State) Tiles() [][]MapTile {
	tiles := *state.tiles
	return tiles
}

func (state *State) StartUpdate(updateCount int) chan bool {
	state.setUpdateCount(updateCount)
	return state.UpdateProcessed
}

func (state *State) UpdateTiles() {
	tiles := *state.tiles

	for update := range state.UpdateTile {
		tiles[update.X][update.Y].Tile = update.Tile

		state.setUpdateCount(-1)

		if state.tilesToUpdate == 0 {
			state.UpdateProcessed <- true
		}
	}
}

func (state *State) setUpdateCount(value int) {
	var mu sync.Mutex

	mu.Lock()
	state.tilesToUpdate = state.tilesToUpdate + value
	mu.Unlock()
}
