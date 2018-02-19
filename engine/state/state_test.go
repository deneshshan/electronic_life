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

// Helper functions
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

func CheckStateEquivalent(t *testing.T, sw *StateWriter, input string, w, h int) {
	parsed := ParseTestRepresentation(t, input, w, h)
	tiles := *sw.state.tiles

	for i, row := range tiles {
		for j, tile := range row {
			parsedTile := parsed[i][j].Tile
			if parsedTile != tile.Tile {
				t.Fatalf("Parsed tile was %v whilst it is %v in state", parsedTile, tile.Tile)
			}
		}
	}

}
