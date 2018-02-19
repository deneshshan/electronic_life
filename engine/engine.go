package engine

import (
	st "github.com/deneshshan/electronic_life/engine/state"
)

type Engine struct {
	state            *st.State
	DisplayConnector chan *st.State
	Fps              float64
}

// TODO: needs to connect to state reader
//func (engine *Engine) BaseLevel() [][]types.MapTile {
//state := *engine.state
//return state.Tiles()
//}

func (engine *Engine) work() {
	// 1. entities act
	// 2. update map
	// 3. send map to renderer
	//go engine.updateDisplay()
}
