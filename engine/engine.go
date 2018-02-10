package engine

import (
	maps "github.com/deneshshan/electronic_life/engine/maps"
)

type Engine struct {
	levelMap         *maps.Map
	DisplayConnector chan *maps.Map
	Fps              float64
}

func (engine *Engine) BaseLevel() [][]maps.TileChar {
	mapp := *engine.levelMap
	return mapp.Tiles()
}

func (engine *Engine) work() {
	// 1. entities act
	// 2. update map
	// 3. send map to renderer
	//go engine.updateDisplay()
}
