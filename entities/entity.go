package entities

import (
	"github.com/looplab/fsm"
)

type Mover interface {
	Move()
}

type Grower interface {
	Grow()
}

type Entity struct {
	Type EntityType
	FSM  *fsm.FSM
}
