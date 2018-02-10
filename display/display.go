package display

import (
	tl "github.com/JoelOtter/termloop"
	"github.com/deneshshan/electronic_life/engine"
	maps "github.com/deneshshan/electronic_life/engine/maps"
)

type Display struct {
	engineConnector chan *maps.Map
}

func NewDisplay(engine *engine.Engine) {
	game := tl.NewGame()

	game.Screen().SetFps(engine.Fps)

	buildLevel(game, engine)

	game.Start()
}
