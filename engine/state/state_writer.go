package state

import (
	"errors"
	"sync"

	"github.com/deneshshan/electronic_life/engine/types"
)

// StateWriter represents an entity that updates the state of the simulation.
// It is also responsible for creating the state.
type StateWriter struct {
	state           *State
	updateCount     int
	updateCountLock sync.Mutex
	update          chan types.MapTile
	updateProcessed chan bool
}

func NewStateWriter(w, h int) *StateWriter {

	updateDone := make(chan bool, 1)
	update := make(chan types.MapTile, 20)
	stateWriter := StateWriter{state: newState(w, h), update: update, updateProcessed: updateDone}

	go stateWriter.processUpdates()

	return &stateWriter
}

func (sr *StateWriter) processUpdates() {
	tiles := *sr.state.tiles

	for update := range sr.update {
		tiles[update.X][update.Y].Tile = update.Tile
		sr.setUpdates(-1)

		if sr.updatesRemaining() == 0 {
			sr.updateProcessed <- true
		}
	}
}

func (sr *StateWriter) StartBatchUpdate(amountTilesToUpdate int) (update chan types.MapTile, updateDone chan bool, err error) {
	if sr.updatesRemaining() > 0 {
		return nil, nil, errors.New("Some updates are waiting to complete.")
	}

	sr.setUpdates(amountTilesToUpdate)

	return sr.update, sr.updateProcessed, nil
}

func (sr *StateWriter) updatesRemaining() int {
	var remaining int

	sr.updateCountLock.Lock()
	remaining = sr.updateCount
	sr.updateCountLock.Unlock()

	return remaining
}

func (sr *StateWriter) setUpdates(amount int) {
	sr.updateCountLock.Lock()
	sr.updateCount = sr.updateCount + amount
	sr.updateCountLock.Unlock()
}
