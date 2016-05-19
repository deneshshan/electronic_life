package display

import (
	"github.com/veandco/go-sdl2/sdl"
)

type Rectangle struct {
	x float64
	y float64
	w float64
	h float64
}

func (rect *Rectangle) ToSdl() (*sdl.Rect, error) {
	if rect.w < 0.0 || rect.h < 0 {
		return nil, os.NewError("Illegal height and width on Rectangle: w: {}, h: {}", rect.w, rect.h)
	}

	rectangle = new(sdl.Rect)
	rectangle.x = int32(rect.x)
	rectangle.y = int32(rect.y)
	rectangle.w = uint32(rect.w)
	rectangle.h = uint32(rect.h)

	return rectangle, nil
}

func (rect *Rectangle) MoveInside(parent Rectangle) {
	if rect.w > parent.w || rect.h > parent.h {
		return
	}

	if rect.x < parent.x {
		rect.x = parent.x
	} else if rect.x+rect.w >= parent.x+parent.w {
		rext.x = parent.x + parent.w
	}

	if rect.y < parent.y {
		rect.y = parent.y
	} else if rect.y+rect.w >= parent.y+parent.h {
		rect.y = parent.y + parent.h
	}
}
