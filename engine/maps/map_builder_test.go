package map_builders_test

import (
	"testing"

	maps "github.com/deneshshan/electronic_life/engine/maps"
	st "github.com/deneshshan/electronic_life/engine/state"
)

func setupMapBuilder(t *testing.T, width, height int) (*maps.MapBuilder, func()) {
	state := st.NewState(width, height)
	reader := st.NewGameStateReader(state)
	writer := st.NewGameStateWriter(state)

	builder := maps.NewMapBuilder(writer, reader)

	return builder, func() {
		writer.Done <- struct{}{}
	}
}

// Can only build an instance of mapbuilder once
func TestSingletonBuild(t *testing.T) {
}
