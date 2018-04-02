package map_builders

import "github.com/deneshshan/electronic_life/errors"

type GeneratorType int

const (
	Cavern = iota
	File
)

func CreateGenerator(generatorType GeneratorType, builder *MapBuilder) (MapGenerator, error) {
	switch generatorType {
	case Cavern:
		return NewCavernGenerator(builder), nil
	default:
		return nil, &errs.ArgumentError{Arg: generatorType, Issue: "No factory for this type"}
	}
}
