package entities

// Factory method to create entities based on type
func Create(entityType EntityType) actor {
	instance := &entityBuilder{}
	entity := instance.setType(entityType).build()
	return entity
}
