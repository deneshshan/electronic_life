package map_builders

import (
	"sync"

	st "github.com/deneshshan/electronic_life/engine/state"
	"github.com/deneshshan/electronic_life/engine/types"
)

type MapBuilder struct {
	built           sync.Once
	width           int
	height          int
	percentageWalls int
	tiless          [][]types.MapTile
	generatorType   GeneratorType
	stateReader     st.IStateReader
	updater         st.StateBatchUpdater
	updates         chan<- types.MapTile
}

func NewMapBuilder(updater st.StateBatchUpdater, snapshotter st.IStateReader) *MapBuilder {
	map_builder := MapBuilder{width: 10, height: 10, percentageWalls: 20, updater: updater, stateReader: snapshotter, generatorType: Cavern}

	return &map_builder
}

func (builder *MapBuilder) SetWidth(width int) *MapBuilder {
	builder.width = width
	return builder
}

func (builder *MapBuilder) SetHeight(height int) *MapBuilder {
	builder.height = height
	return builder
}

func (builder *MapBuilder) SetPercentageWalls(percentageWalls int) *MapBuilder {
	builder.percentageWalls = percentageWalls
	return builder
}

func (builder *MapBuilder) SetGeneratorType(gtype GeneratorType) *MapBuilder {
	builder.generatorType = gtype
	return builder
}

func (builder *MapBuilder) Build() {
	builder.built.Do(func() {
		map_generator, err := CreateGenerator(builder.generatorType, builder)

		if err != nil {
			panic("Could not create map generator")
		}

		builder.tiless = builder.stateReader.Read()
		map_generator.Generate()

		builder.finalize(builder.area())
	})
}

func (builder *MapBuilder) finalize(amountToUpdate int) {
	updates, done, err := builder.updater.BatchUpdate(amountToUpdate)

	if err != nil {
		panic("Error whilst updating state in cavern generator")
	}

	for _, row := range builder.tiless {
		for _, tile := range row {
			updates <- tile
		}
	}

	<-done
}

func (builder *MapBuilder) tiles() [][]types.MapTile {
	return builder.tiless
}

func (builder *MapBuilder) area() int {
	return builder.width * builder.height
}
