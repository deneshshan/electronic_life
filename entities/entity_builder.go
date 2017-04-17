package entities

type EntityType int

const (
	carrot EntityType = iota
	rabbit
)

type EntityBuilder interface {
	setType(entityType EntityType) EntityBuilder
	build() actor
}

type entityBuilder struct {
}

func (eb entityBuilder) setType(entityType EntityType) EntityBuilder {
	switch entityType {
	case rabbit:
		return &rabbitBuilder{}
	default:
		panic("EntityBuilder typeOption not recognised")
	}
}
