package entities

type EntityType int

const (
	Carrot EntityType = iota
	Rabbit
)

type entityBuilder struct {
	typeOption int
}

type EntityBuilder interface {
	setType()
	build()
}

func (eb *entityBuilder) Type(entityType EntityType) EntityBuilder {
	eb.typeOption = entityType
	return eb
}

func (eb *entityBuilder) Build() (EntityBuilder, error) {
	switch eb.typeOption {
	case EntityType.Rabbit:
		entity := &Rabbit{Type: eb.typeOption}
		entity.FSM = getRabbitFSM()
		return entity
	default:
		panic("EntityBuilder typeOption not recognised")
	}
}

func getRabbitFSM() *fsm.FSM {
}

func Create(entityType EntityType) *Entity {
	instance := &entityBuilder{}
	entity := instance.setType(entityType).build()
	return entity
}
