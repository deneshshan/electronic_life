package display

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Events struct {
	Quit bool
}

func newEvents() (events *Events) {
	events = &Events{
		Quit: false,
	}

	return events
}

func (events *Events) monitor() {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		case *sdl.QuitEvent:
			events.Quit = true
			break
		}
	}
}
