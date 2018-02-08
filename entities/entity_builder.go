//package entities

//import (
	//"github.com/pkg/errors"
//)

//type EntityType int

//const (
	//carrot EntityType = iota
	//rabbit
//)

//type buildable interface {
	//build() actor
//}

//func NewBuilder(entityType EntityType) (buildable, error) {
	//switch entityType {
	//case rabbit:
		//return &rabbitBuilder{}, nil
	//default:
		//err := errors.New("EntityBuilder typeOption not recognised")
		//return nil, errors.Wrap(err, "entityBuilder.SetType()")
	//}
//}
