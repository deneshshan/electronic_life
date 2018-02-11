package engine

type Engine struct {
	state            *State
	DisplayConnector chan *State
	Fps              float64
}

func (engine *Engine) BaseLevel() [][]MapTile {
	state := *engine.state
	return state.Tiles()
}

func (engine *Engine) work() {
	// 1. entities act
	// 2. update map
	// 3. send map to renderer
	//go engine.updateDisplay()
}
