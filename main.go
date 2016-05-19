package main

import (
	"fmt"
	"github.com/veandco/go-sdl2/sdl"
	"os"
)

const windowTitle string = "Electronic Life"

var winWidth, winHeight int = 800, 600

func main() {
	run()
}

func run() {
	window, err := sdl.CreateWindow(windowTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		winWidth, winHeight, sdl.WINDOW_SHOWN)
	handle_error(err, "Failed to create window", 1)
	defer window.Destroy()

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	handle_error(err, "Failed to create renderer", 2)
	defer renderer.Destroy()

	events := newEvents()

	for {
		events.monitor()
		if events.quit == true {
			fmt.Println("bye !!!!")
			break
		}

		err = renderer.SetDrawColor(0, 92, 9, 100)
		err = renderer.Clear()
		renderer.Present()
		handle_error(err, "Failed when interacting with renderer", 2)
	}

	os.Exit(0)
}

func handle_error(err error, message string, returnv int) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "{}: %s\n", message, err)
		os.Exit(returnv)
	}
}
