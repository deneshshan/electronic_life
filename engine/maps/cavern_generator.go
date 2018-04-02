package map_builders

import (
	"math/rand"
	"sync"
	"time"

	"github.com/deneshshan/electronic_life/engine/types"
)

type CavernGenerator struct {
	builder *MapBuilder
}

func NewCavernGenerator(builder *MapBuilder) *CavernGenerator {
	cg := CavernGenerator{builder: builder}

	return &cg
}

// Generates walls and writes them via the State Writer. Can only be called once.
func (cg *CavernGenerator) Generate() {
	cg.generateRandomWalls()

	for i := 0; i < 5; i++ {
		cg.generateCaverns()
	}
}

func (cg *CavernGenerator) area() int {
	return cg.builder.area()
}

func (cg *CavernGenerator) width() int {
	return cg.builder.width
}

func (cg *CavernGenerator) height() int {
	return cg.builder.height
}

func (cg *CavernGenerator) percentageWalls() int {
	return cg.builder.percentageWalls
}

func (cg *CavernGenerator) generateRandomWalls() {
	var mapMiddle int = 0
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))

	mapMiddle = (cg.width() / 2)

	var wg sync.WaitGroup

	for column := 0; column < cg.width(); column++ {
		for row := 0; row < cg.height(); row++ {
			wg.Add(1)

			go func(column, row int, random int) {
				defer wg.Done()

				switch {
				case column == 0:
					cg.setTile(column, row, types.Wall)
				case row == 0:
					cg.setTile(column, row, types.Wall)
				case column == (cg.width() - 1):
					cg.setTile(column, row, types.Wall)
				case row == (cg.width() - 1):
					cg.setTile(column, row, types.Wall)
				default:

					if row == mapMiddle {
						cg.setTile(column, row, types.EmptySpace)
					} else {
						if cg.percentageWalls() >= random {
							cg.setTile(column, row, types.Wall)
						} else {
							cg.setTile(column, row, types.EmptySpace)
						}
					}
				}

			}(column, row, rand.Intn(100))

		}
	}

	wg.Wait()
}

func (cg *CavernGenerator) setTile(column int, row int, tile types.Tiles) {
	update := types.MapTile{X: column, Y: row, Tile: tile}
	cg.builder.tiless[column][row] = update
}

func (cg *CavernGenerator) tiles() [][]types.MapTile {
	return cg.builder.tiles()
}

func (cg *CavernGenerator) generateCaverns() {
	for column := 0; column < cg.width(); column++ {
		for row := 0; row < cg.height(); row++ {
			tile := cg.surroundingTile(column, row)
			cg.setTile(column, row, tile)
		}
	}
}

func (cg *CavernGenerator) surroundingTile(column, row int) types.Tiles {
	tiles := cg.tiles()
	numberSurroundingWalls := cg.countAdjacentWalls(tiles, column, row)

	if tiles[column][row].Tile == types.Wall {
		if numberSurroundingWalls >= 4 {
			return types.Wall
		} else if numberSurroundingWalls < 2 {
			return types.EmptySpace
		}
	} else {
		if numberSurroundingWalls >= 5 {
			return types.Wall
		}
	}

	return types.EmptySpace
}

func (cg *CavernGenerator) countAdjacentWalls(tiles [][]types.MapTile, column int, row int) int {
	start_col := column - 1
	end_col := column + 1
	start_row := row - 1
	end_row := row + 1

	wall_count := 0

	for col := start_col; col <= end_col; col++ {
		for rw := start_row; rw <= end_row; rw++ {
			if (col == column && rw == row) == false {
				if cg.isConsideredWall(tiles, col, rw) {
					wall_count++
				}
			}
		}
	}

	return wall_count
}

func (cg *CavernGenerator) isConsideredWall(tiles [][]types.MapTile, column, row int) bool {
	if column < 0 || row < 0 {
		return true
	} else if column > cg.width()-1 || row > cg.height()-1 {
		return true
	} else if tiles[column][row].Tile == types.Wall {
		return true
	} else if tiles[column][row].Tile == types.EmptySpace {
		return false
	}

	return false
}
