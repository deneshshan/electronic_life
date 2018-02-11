package engine

import (
	"sync"
)

var (
	engine_builder_built sync.Once

	engine_builder EngineBuilder
)

type EngineBuilder struct {
	width           int
	height          int
	percentageWalls int
	fps             float64
	built           sync.Once
	engine          *Engine
}

func (builder *EngineBuilder) Width(width int) *EngineBuilder {
	builder.width = width
	return builder
}

func (builder *EngineBuilder) Height(height int) *EngineBuilder {
	builder.height = height
	return builder
}

func (builder *EngineBuilder) PercentageWalls(percentageWalls int) *EngineBuilder {
	builder.percentageWalls = percentageWalls
	return builder
}

func (builder *EngineBuilder) Fps(fps float64) {
	builder.fps = fps
}

//TODO error
func GetEngineBuilder() *EngineBuilder {
	engine_builder_built.Do(func() {
		engine_builder = EngineBuilder{width: 10, height: 10, percentageWalls: 20, fps: 30}
	})

	return &engine_builder
}

//TODO: error
func (builder *EngineBuilder) Build() *Engine {
	engine := Engine{}

	builder.built.Do(func() {
		map_builder := GetMapBuilder()
		map_builder.Width(builder.width).Height(builder.height).PercentageWalls(builder.percentageWalls)
		levelMap := map_builder.Build()

		displayConnector := make(chan *State, 2)

		engine.state = levelMap
		engine.DisplayConnector = displayConnector
		engine.Fps = builder.fps
	})

	return &engine
}