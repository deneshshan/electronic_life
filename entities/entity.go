package entities

import (
	"github.com/looplab/fsm"
)

type actor interface {
	act()
}

type stateMachine interface {
	enterState(e *fsm.Event)
}

type Entity struct {
	Type EntityType
	FSM  *fsm.FSM
}

func (ent Entity) act() {
}

func (ent Entity) enterState(e *fsm.Event) {
}
