package map_builders_test

import (
	"testing"

	maps "github.com/deneshshan/electronic_life/engine/maps"
	st "github.com/deneshshan/electronic_life/engine/state"
	"github.com/deneshshan/electronic_life/engine/types"
)

type mockStateReader struct {
	st.GameStateReader
}

func (ms *mockStateReader) countWalls() int {
	tiles := ms.GameStateReader.Read()

	wallCount := 0

	for _, row := range tiles {
		for _, tile := range row {
			if tile.Tile == types.Wall {
				wallCount += 1
			}
		}
	}

	return wallCount
}

func setupCavernBuilder(b *testing.B, width, height, percentageWalls int) (*maps.CavernGenerator, *mockStateReader, func()) {
	state := st.NewState(width, height)
	writer := st.NewGameStateWriter(state)
	reader := &mockStateReader{*st.NewGameStateReader(state)}
	builder := maps.NewMapBuilder(writer, reader)
	builder = builder.SetWidth(width).SetHeight(height).SetPercentageWalls(percentageWalls)
	builder.Build()

	cg := maps.NewCavernGenerator(builder)

	return cg, reader, func() {
		writer.Done <- struct{}{}
	}
}

func BenchmarkLarge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_, _, teardown := setupCavernBuilder(b, 100, 3000, 45)
		defer teardown()
	}
}
