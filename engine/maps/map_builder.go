package engine_maps

import (
	"math/rand"
	"sync"
	"time"
)

var (
	map_builder_built sync.Once

	map_builder MapBuilder
)

type MapBuilder struct {
	width           int
	height          int
	percentageWalls int
	built           sync.Once
	update          chan MapTile
	mapp            *Map
}

func GetMapBuilder() *MapBuilder {
	map_builder_built.Do(func() {
		map_builder = MapBuilder{width: 10, height: 10, percentageWalls: 20}
	})

	return &map_builder
}

func (builder *MapBuilder) Width(width int) *MapBuilder {
	builder.width = width
	return builder
}

func (builder *MapBuilder) Height(height int) *MapBuilder {
	builder.height = height
	return builder
}

func (builder *MapBuilder) PercentageWalls(percentageWalls int) *MapBuilder {
	builder.percentageWalls = percentageWalls
	return builder
}

func (builder *MapBuilder) Build() *Map {
	mapp := &Map{}

	builder.built.Do(func() {

		tiles := builder.initilizeLayout()
		builder.update = make(chan MapTile, 20)

		mapp = newMap(tiles, builder.update)
		builder.mapp = mapp

		go mapp.UpdateTiles()

		done := mapp.StartUpdate(builder.area())
		builder.generateRandomWalls()
		<-done

		for count := 0; count <= 5; count++ {
			builder.generateCaverns(mapp.Tiles())
		}

	})

	return mapp
}

func (builder *MapBuilder) area() int {
	return builder.width * builder.height
}

func (builder *MapBuilder) initilizeLayout() *[][]TileChar {
	tiles := make([][]TileChar, builder.width)

	for row := range tiles {
		tiles[row] = make([]TileChar, builder.height)
	}

	return &tiles
}

func (builder *MapBuilder) generateRandomWalls() {
	var mapMiddle int = 0
	rand := rand.New(rand.NewSource(time.Now().UnixNano()))

	var wg sync.WaitGroup

	mapMiddle = (builder.height / 2)

	for column := 0; column < builder.width; column++ {
		for row := 0; row < builder.height; row++ {
			wg.Add(1)

			go func(column, row int, random int) {
				defer wg.Done()

				switch {
				case column == 0:
					builder.setTile(column, row, Wall)
				case row == 0:
					builder.setTile(column, row, Wall)
				case column == (builder.width - 1):
					builder.setTile(column, row, Wall)
				case row == (builder.height - 1):
					builder.setTile(column, row, Wall)
				default:

					if row == mapMiddle {
						builder.setTile(column, row, EmptySpace)
					} else {
						if builder.percentageWalls >= random {
							builder.setTile(column, row, Wall)
						} else {
							builder.setTile(column, row, EmptySpace)
						}
					}
				}

			}(column, row, rand.Intn(100))

		}
	}

	wg.Wait()
}

func (builder *MapBuilder) setTile(column int, row int, tile TileChar) {
	update := MapTile{X: column, Y: row, Tile: tile}
	builder.update <- update
}

func (builder *MapBuilder) generateCaverns(tiles [][]TileChar) {
	mapp := builder.mapp

	for column := 0; column < builder.width; column++ {
		for row := 0; row < builder.height; row++ {
			done := mapp.StartUpdate(1)

			tile := builder.surroundingTile(tiles, column, row)
			builder.setTile(column, row, tile)

			<-done
		}
	}
}

func (builder *MapBuilder) surroundingTile(tiles [][]TileChar, column, row int) TileChar {
	numberSurroundingWalls := builder.countAdjacentWalls(tiles, column, row)

	if tiles[column][row] == Wall {
		if numberSurroundingWalls >= 4 {
			return Wall
		} else if numberSurroundingWalls < 2 {
			return EmptySpace
		}
	} else {
		if numberSurroundingWalls >= 5 {
			return Wall
		}
	}

	return EmptySpace
}

func (builder *MapBuilder) countAdjacentWalls(tiles [][]TileChar, column int, row int) int {
	start_col := column - 1
	end_col := column + 1
	start_row := row - 1
	end_row := row + 1

	wall_count := 0

	for col := start_col; col <= end_col; col++ {
		for rw := start_row; rw <= end_row; rw++ {
			if (col == column && rw == row) == false {
				if builder.isConsideredWall(tiles, col, rw) {
					wall_count++
				}
			}
		}
	}

	return wall_count
}

func (builder *MapBuilder) isConsideredWall(tiles [][]TileChar, column, row int) bool {
	if column < 0 || row < 0 {
		return true
	} else if column > builder.width-1 || row > builder.height-1 {
		return true
	} else if tiles[column][row] == Wall {
		return true
	} else if tiles[column][row] == EmptySpace {
		return false
	}

	return false
}
