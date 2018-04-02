package map_builders

import "github.com/deneshshan/electronic_life/engine/types"

type MapFromFileGenerator struct{}

func (mfg *MapFromFileGenerator) GetWalls() []types.MapTile {
	return nil
}
