package display

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Events struct {
	quit bool
}

func NewEvents() (events *Events) {
	events = &Events{
		quit: false,
	}

	return events
}

func (events *Events) monitor() {
	for event := sdl.PollEvent(); event != nil; event = sdl.PollEvent() {
		switch event.(type) {
		case *sdl.QuitEvent:
			events.quit = true
			break
		}
	}
}
