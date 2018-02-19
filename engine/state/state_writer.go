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
	Update          chan types.MapTile
	UpdateProcessed chan struct{}
	Done            chan struct{}
}

//
// Creates a new StateWriter object which creates the state.
//
func NewStateWriter(w, h int) *StateWriter {

	updateDone := make(chan struct{}, 1)
	update := make(chan types.MapTile, 20)
	done := make(chan struct{}, 1)

	stateWriter := StateWriter{state: newState(w, h), Update: update, UpdateProcessed: updateDone, Done: done}

	go stateWriter.processUpdates()

	return &stateWriter
}

//
// Processes the updating of the game state. The thing sending the updates should block on sw.UpdateProcessed
// Until continuing if another goroutine is proving the updates. Also handles closing of the UpdateProcesed
// channel.
//
func (sw *StateWriter) processUpdates() {
	tiles := *sw.state.tiles

	for {
		select {
		case update := <-sw.Update:
			tiles[update.X][update.Y].Tile = update.Tile
			sw.setUpdates(-1)

			if sw.updatesRemaining() == 0 {
				sw.UpdateProcessed <- struct{}{}
			}
		case <-sw.Done:
			close(sw.UpdateProcessed)

			return
		}
	}
}

//
// Called before any goroutes update the state. The amount of updates to be processed is passed in as an
// argument. Returns an error if attempting to call this method whilst other updates are in progress.
//
func (sw *StateWriter) StartBatchUpdate(amountTilesToUpdate int) (err error) {
	if sw.updatesRemaining() > 0 {
		return errors.New("Some updates are waiting to complete.")
	}

	sw.setUpdates(amountTilesToUpdate)

	return nil
}

//
// Returns how many updates remaining to be processed.
//
func (sw *StateWriter) updatesRemaining() int {
	defer sw.updateCountLock.Unlock()

	var remaining int

	sw.updateCountLock.Lock()
	remaining = sw.updateCount

	return remaining
}

//
// Addes the amount to the current amount of updates to be processed.
//
func (sw *StateWriter) setUpdates(amount int) {
	defer sw.updateCountLock.Unlock()

	sw.updateCountLock.Lock()
	sw.updateCount = sw.updateCount + amount
}
