package entities

func Create(entityType EntityType) actor {
	instance := &entityBuilder{}
	entity := instance.setType(entityType).build()
	return entity
}
