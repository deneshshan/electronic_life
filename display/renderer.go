package display

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	windowTitle string = "Electronic Life"
	winWidth    int    = 800
	winHeight   int    = 600
)

type Renderer struct {
	window   *sdl.Window
	renderer *sdl.Renderer
}

type RenderingResult struct {
	Success     bool
	ReturnValue int
	Reason      string
	Err         error
}

func Start() RenderingResult {
	window, err := sdl.CreateWindow(windowTitle, sdl.WINDOWPOS_UNDEFINED, sdl.WINDOWPOS_UNDEFINED,
		winWidth, winHeight, sdl.WINDOW_SHOWN)
	if err != nil {
		return newErrorRendingResult("Failed to create window", 1, err)
	}

	renderer, err := sdl.CreateRenderer(window, -1, sdl.RENDERER_ACCELERATED)
	if err != nil {
		return newErrorRendingResult("Failed to create renderer", 2, err)
	}

	displayRenderer := &Renderer{window: window, renderer: renderer}
	defer destroy(displayRenderer)

	err = render(displayRenderer)
	if err != nil {
		return newErrorRendingResult("Failed when interacting with renderer", 2, err)
	}

	return newSuccessRendingResult()
}

func newErrorRendingResult(reason string, returnValue int, err error) RenderingResult {
	return RenderingResult{Success: false, Reason: reason, ReturnValue: returnValue, Err: err}
}

func newSuccessRendingResult() RenderingResult {
	return RenderingResult{Success: true}
}

func render(displayRenderer *Renderer) error {
	events := NewEvents()

	for {
		events.Monitor()
		if events.Quit == true {
			break
		}

		err := displayRenderer.renderer.SetDrawColor(0, 92, 9, 100)
		err = displayRenderer.renderer.Clear()
		displayRenderer.renderer.Present()
		if err != nil {
			return err
		}
	}

	return nil
}

func destroy(displayRenderer *Renderer) {
	displayRenderer.window.Destroy()
	displayRenderer.renderer.Destroy()
}
