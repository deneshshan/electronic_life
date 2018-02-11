package engine_factories

import "github.com/looplab/fsm"

type CarrotFactory struct {
	EntityFactory
}

func (cf *CarrotFactory) Create() *Entity {
	fsm := fsm.NewFSM(
		"grow",
		fsm.Events{
			{Name: "grow", Src: []string{"grow", "reproduce"}, Dst: []string{"grow", "reproduce"}},
			{Name: "reproduce", Src: "grow", Dst: "reproduced"},
			{Name: "reproduced", Src: "reproduce", Dst: "grow"},
		},
		fsm.Callbacks{
			"enter_state": func(e *fsm.Event) { d.enterState(e) },
		},
	)

	tick := make(chan struct{}, 1)
	done := make(chan struct{}, 1)
	healthUpdate := make(chan int, 2)
	nextAction := make(chan Action, 2)

	carrot := Carrot{
		Tick:         tick,
		fsm:          fsm,
		done:         done,
		health:       100,
		healthUpdate: healthUpdate,
		NextAction:   nextAction,
	}

	go carrot.Act()

	return &carrot
}
