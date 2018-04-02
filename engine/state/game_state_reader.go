package state

import (
	types "github.com/deneshshan/electronic_life/engine/types"
)

type GameStateReader struct {
	state *State
}

func NewGameStateReader(st *State) *GameStateReader {
	sr := GameStateReader{state: st}

	return &sr
}

func (sr *GameStateReader) Read() [][]types.MapTile {
	return sr.state.Read()
}
