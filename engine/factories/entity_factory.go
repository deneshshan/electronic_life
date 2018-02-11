package engine_factories

type EntityFactory interface {
	Create(int, int) *Entity
}

func GetFactory(factory EntityType) (*EntityFactory, error) {
	switch factory {
	case Carrot:
		return &CarrotFactory{}, nil
	case Rabbit:
		return &RabbitFactory{}, nil
	default:
		return nil, ArgumentError{Arg: factory, "Invalid factory type"}
	}
}
