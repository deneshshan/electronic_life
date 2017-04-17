package entities

import (
	"github.com/looplab/fsm"
)

type rabbitBuilder struct {
	entityBuilder
}

func (eb *rabbitBuilder) build() actor {
	entity := &Rabbit{}
	entity.FSM = getRabbitFSM(entity)
	return entity
}

func getRabbitFSM(sm stateMachine) *fsm.FSM {
	fsm := fsm.NewFSM(
		"searching",
		fsm.Events{},
		fsm.Callbacks{
			"enter_state": func(e *fsm.Event) { sm.enterState(e) },
		},
	)

	return fsm
}
