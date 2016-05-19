package life

import (
	"display"
)

type entity struct {
	rect Rectangle
}

func NewEntity() (ent *entity) {
	rectangle = new(display.Rectangle)
	rectangle.x = 64.0
	rectangle.y = 64.0
	rectangle.w = 32.0
	rectangle.h = 32.0

	ent = &entity{}
	ent.rect = rectangle

	return ent
}

func (ent *entity) MoveInside(rect Rectangle, parent Rectangle) {
}
