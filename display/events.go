package display

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Events struct {
	Quit bool
}

func NewEvents() (events *Events) {
	events = &Events{
		Quit: false,
	}

	return events
}

func (events *Events) Monitor() {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		case *sdl.QuitEvent:
			events.Quit = true
			break
		}
	}
}
