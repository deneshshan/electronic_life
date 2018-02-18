package main

import (
	"fmt"
	"os"

	"github.com/buger/goterm"
	"github.com/deneshshan/electronic_life/display"
	"github.com/deneshshan/electronic_life/engine"
)

func main() {

	engine_builder := engine.GetEngineBuilder()

	width := goterm.Width()
	height := goterm.Height()

	engine_builder.Width(width).Height(height).PercentageWalls(45).Fps(30)
	engine := engine_builder.Build()

	display.New(engine)

	fmt.Println("Bye!!")
	os.Exit(0)
}
