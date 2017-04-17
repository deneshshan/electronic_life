package entities

import (
	"github.com/looplab/fsm"
)

type EntityType int

const (
	carrot EntityType = iota
	rabbit
)

type entityBuilder struct {
}

type rabbitBuilder struct {
	entityBuilder
}

type EntityBuilder interface {
	setType(entityType EntityType) EntityBuilder
	build() actor
}

func Create(entityType EntityType) actor {
	instance := &entityBuilder{}
	entity := instance.setType(entityType).build()
	return entity
}

func (eb entityBuilder) setType(entityType EntityType) EntityBuilder {
	switch entityType {
	case rabbit:
		return &rabbitBuilder{}
	default:
		panic("EntityBuilder typeOption not recognised")
	}
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
