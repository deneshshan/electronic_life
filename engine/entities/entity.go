package engine_entities

type Positionable interface {
	SetPosition(Position)
}

type EntityType int

const (
	Carrot = iota
	Rabbit
)

type ActionType int

const (
	Spawn = iota
	Move
	Feed
)

type ActionDirection int

const (
	North = iota
	NortWest
	West
	SouthWest
	South
	SouthEast
	East
	NortEast
)

type Position struct {
	X int
	Y int
}

type Action struct {
	Type            ActionType
	Direction       ActionDirection
	CurrentPosition Position
	TargetPosition  Position
}

type Entity struct {
	Positionable
	position     Position
	health       int
	fsm          *fsm.FSM
	Tick         chan struct{}
	Done         chan struct{}
	NextAction   chan Action
	healthUpdate chan int
}

func (ent *Entity) SetPosition(position Position) {
	ent.position = position
}
