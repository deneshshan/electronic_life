package state_test

import (
	"testing"

	state "github.com/deneshshan/electronic_life/engine/state"
	types "github.com/deneshshan/electronic_life/engine/types"
)

func setupStateWriter(t *testing.T, w int, h int) (*state.StateWriter, func()) {
	//if state != nil {
	//t.Errorf("Error on unit setup: %v", err)
	//return nil, func() {}
	//}
	sw := state.NewStateWriter(w, h)
	return sw, func() {
		//state.DestroyState()
	}
}

const (
	testEmptyMap = `
....
....
....
....
`
	testAllWallsMap = `
XXXXXXXX
XXXXXXXX
`
	testSomeWallsMap = `
XXXXXXX
XX....X
X....XX
XX..XXX
XXXXXXX
	`
)

func TestUpdate(t *testing.T) {
	tests := []struct {
		name, mapp string
		w, h       int
	}{
		{"Empy map", testEmptyMap, 4, 4},
		{"All walls map", testAllWallsMap, 8, 2},
		{"Some walls map", testSomeWallsMap, 7, 5},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			sw, teardown := setupStateWriter(t, test.w, test.h)
			defer teardown()

			updates, updateDone, err := sw.StartBatchUpdate(test.w * test.h)

			if err != nil {
				t.Fatalf("Error on starting update. Other updates in progress")
			}

			parsed := state.ParseTestRepresentation(t, test.mapp, test.w, test.h)

			for i := 0; i < test.h; i++ {
				for j := 0; j < test.w; j++ {
					go func(tile types.MapTile) {
						updates <- tile
					}(parsed[i][j])
				}
			}
			<-updateDone

			state.CheckStateEquivalent(t, sw, test.mapp, test.w, test.h)
		})
	}
}

func TestStartingUpdateInProgressReturnsError(t *testing.T) {
	t.Run("Update in progress", func(t *testing.T) {
		w, h := 2, 2
		sw, teardown := setupStateWriter(t, w, h)
		defer teardown()

		updates, _, err := sw.StartBatchUpdate(w * h)

		updates <- types.MapTile{X: 0, Y: 0, Tile: types.Wall}

		_, _, err = sw.StartBatchUpdate(w * h)

		if err == nil {
			t.Fatalf("Should return an error")
		}
	})
}
