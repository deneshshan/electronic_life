package display

import (
	tl "github.com/JoelOtter/termloop"
	eng "github.com/deneshshan/electronic_life/engine"
	maps "github.com/deneshshan/electronic_life/engine/maps"
)

func buildLevel(game *tl.Game, engine *eng.Engine) {
	level := tl.NewBaseLevel(tl.Cell{
		Bg: tl.ColorBlack,
		Fg: tl.ColorGreen,
		Ch: ' ',
	})

	levelMap := engine.BaseLevel()

	for i, column := range levelMap {
		processColumns(level, i, column)
	}

	game.Screen().SetLevel(level)
}

func processColumns(level *tl.BaseLevel, col_index int, column []maps.TileChar) {
	for row_index, cell := range column {
		if cell == maps.Wall {
			level.AddEntity(tl.NewRectangle(col_index, row_index, 1, 1, tl.ColorBlue))
		}
	}
}
