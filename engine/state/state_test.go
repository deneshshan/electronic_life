package state

import (
	"strings"
	"testing"

	types "github.com/deneshshan/electronic_life/engine/types"
)

var tileRepresentations = map[rune]types.Tiles{
	'.': types.EmptySpace,
	'X': types.Wall,
}

func TestInitializeLayout(t *testing.T) {
	tests := []struct {
		name                  string
		w, h, percentageWalls int
	}{
		{"9x20", 9, 20, 15},
		{"9x9", 9, 9, 15},
		{"40x30", 40, 30, 45},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			st, teardown := setupState(t, test.w, test.h)
			defer teardown()

			tiles := *st.tiles

			for row := range tiles {
				row_length := len(tiles[row])
				if row_length != test.w {
					t.Fatalf("Initialized row wrong size. Should be %v but is %v", test.w, row_length)
				}
			}
		})
	}
}

// Tests that the returned state snapshot is not the actual state of the simulation.
func TestStateSnapshotIsNotEqualToState(t *testing.T) {
	state, teardown := setupState(t, 5, 20)
	defer teardown()

	snapshot := state.Read()

	if &snapshot == state.tiles {
		t.Fatalf("Snapshot should not reference the actual state")
	}
}

// Helper functions
func setupState(t *testing.T, w, h int) (*State, func()) {
	st := NewState(w, h)
	return st, func() {
	}
}

func ParseTestRepresentation(t *testing.T, input string, w, h int) [][]types.MapTile {
	parsed := make([][]types.MapTile, h)
	for x, row := range strings.Split(strings.TrimSpace(input), "\n") {
		parsed[x] = make([]types.MapTile, w)
		for y, char := range row {
			representation, ok := tileRepresentations[char]
			if !ok {
				t.Fatalf("Error parsing test map. Tile %c not recognised", char)
			} else {
				parsed[x][y].Tile = representation
				parsed[x][y].X = x
				parsed[x][y].Y = y
			}
		}
	}

	return parsed
}

func CheckStateEquivalent(t *testing.T, state *State, tilesToCompare [][]types.MapTile, w, h int) {
	tiles := *state.tiles

	for i, row := range tiles {
		for j, tile := range row {
			comparison := tilesToCompare[i][j].Tile
			if comparison != tile.Tile {
				t.Fatalf("Parsed tile was %v whilst it is %v in state", comparison, tile.Tile)
			}
		}
	}
}
