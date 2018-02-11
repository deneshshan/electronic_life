package display

import (
	tl "github.com/JoelOtter/termloop"
	"github.com/deneshshan/electronic_life/engine"
)

type Display struct {
	engineConnector chan *engine.State
}

func NewDisplay(engine *engine.Engine) {
	game := tl.NewGame()

	game.Screen().SetFps(engine.Fps)

	buildLevel(game, engine)

	game.Start()
}
