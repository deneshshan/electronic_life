package engine

type MapBuilder struct {
	width           int
	height          int
	percentageWalls int
	built           bool
}

func New() *MapBuilder {
	builder := MapBuilder{width: 10, height: 10, percentageWalls: 20, built: false}
	return &builder
}

func (builder *MapBuilder) Width(width int) {
	if builder.built != true {
		builder.width = width
	}
}

func (builder *MapBuilder) Height(height int) {
	if builder.built != true {
		builder.height = height
	}
}

func (builder *MapBuilder) PercentageWalls(percentageWalls int) {
	if builder.built != true {
		builder.percentageWalls = percentageWalls
	}
}

func (builder *MapBuilder) Build() *Map {
	tiles := builder.initilizeLayout()
}

func (builder *MapBuilder) initilizeLayout() [][]int {
	tiles := make([][]int, builder.height)

	for row := range tiles {
		tiles[row] = make([]int, builder.width)
	}

	return tiles
}

func (builder *MapBuilder) generateWalls(tiles *[][]int) {
	//switch {
	//case
	//}
}
