//package engine_entities

//type Collidable interface {
//IsSolid() bool
//Collide(Collidable)
//}

//type Actor interface {
//Act()
//}

//type stateMachine interface {
//enterState(e *fsm.Event)
//}

//type Entity struct {
//Collidable
//Actor
//Type EntityType
//FSM  *fsm.FSM
//}

//func (ent Entity) act() {
//}

//func (ent Entity) enterState(e *fsm.Event) {
//}

//type Destroyable interface {
//Destroy() error
//}

//type Action struct {
//}

//type ResultValue int

//const (
//OK = iota
//)

//type Result struct {
//Value ResultValue
//}

//type Entity {
//actions <-chan  Action
//results ->chan  Result
//}
