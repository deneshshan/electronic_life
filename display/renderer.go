package display

import (
	"github.com/veandco/go-sdl2/sdl"
)

const (
	windowTitle string = "Electronic Life"
	winWidth    int    = 800
	winHeight   int    = 600
)

type sdlContext struct {
	window   *sdl.Window
	renderer *sdl.Renderer
}

type RenderingResult struct {
	Success     bool
	ReturnValue int
	ErrorReason string
	Err         error
}

// Initialises SDL and starts rendering the scene.
//
// Returns a RenderingResult wrapping up whether initialising rendering was successful.
// Any error values are wrapped up in the RenderingResult
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

	context := &sdlContext{window: window, renderer: renderer}
	defer destroy(context)

	err = render(context)
	if err != nil {
		return newErrorRendingResult("Failed when interacting with renderer", 2, err)
	}

	return newSuccessRendingResult()
}

func newErrorRendingResult(errorReason string, returnValue int, err error) RenderingResult {
	return RenderingResult{Success: false, ErrorReason: errorReason, ReturnValue: returnValue, Err: err}
}

func newSuccessRendingResult() RenderingResult {
	return RenderingResult{Success: true}
}

func render(context *sdlContext) error {
	events := newEvents()

	for {
		events.monitor()
		if events.Quit == true {
			break
		}

		err := context.renderer.SetDrawColor(0, 92, 9, 100)
		err = context.renderer.Clear()
		context.renderer.Present()
		if err != nil {
			return err
		}
	}

	return nil
}

func destroy(context *sdlContext) {
	context.window.Destroy()
	context.renderer.Destroy()
}
