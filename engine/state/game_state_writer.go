package state

import (
	"errors"
	"sync"

	"github.com/deneshshan/electronic_life/engine/types"
)

type StateBatchUpdater interface {
	BatchUpdate(int) (chan<- types.MapTile, <-chan struct{}, error)
}

// GameStateWriter represents an entity that updates the state of the simulation.
// It is also responsible for creating the state.
type GameStateWriter struct {
	state           *State
	updateCount     int
	updateCountLock sync.Mutex
	Update          chan types.MapTile
	UpdateProcessed chan struct{}
	Done            chan struct{}
}

// Creates a new StateWriter object which creates the state.
func NewGameStateWriter(state *State) *GameStateWriter {

	update := make(chan types.MapTile, 20)
	updateProcessed := make(chan struct{}, 1)
	done := make(chan struct{}, 1)

	stateWriter := GameStateWriter{state: state, Update: update, UpdateProcessed: updateProcessed, Done: done}

	go stateWriter.processUpdates()

	return &stateWriter
}

// Called before any goroutes update the state. The amount of updates to be processed is passed in as an
// argument. Returns an error if attempting to call this method whilst other updates are in progress.
func (sw *GameStateWriter) BatchUpdate(amountTilesToUpdate int) (chan<- types.MapTile, <-chan struct{}, error) {
	if sw.updatesRemaining() > 0 {
		return nil, nil, errors.New("Some updates are waiting to complete.")
	}

	sw.setUpdates(amountTilesToUpdate)

	return sw.Update, sw.UpdateProcessed, nil
}

// Processes the updating of the game state. The thing sending the updates should block on sw.UpdateProcessed
// Until continuing if another goroutine is proving the updates. Also handles closing of the UpdateProcesed
// channel.
func (sw *GameStateWriter) processUpdates() {
	tiles := *sw.state.tiles

	for update := range sw.Update {
		tiles[update.X][update.Y].Tile = update.Tile
		sw.setUpdates(-1)

		if sw.updatesRemaining() == 0 {
			sw.UpdateProcessed <- struct{}{}
		}
	}
}

// Returns how many updates remaining to be processed.
func (sw *GameStateWriter) updatesRemaining() int {

	var remaining int

	sw.updateCountLock.Lock()
	remaining = sw.updateCount
	sw.updateCountLock.Unlock()

	return remaining
}

// Adds the amount to the current amount of updates to be processed.
func (sw *GameStateWriter) setUpdates(amount int) {

	sw.updateCountLock.Lock()
	sw.updateCount = sw.updateCount + amount
	sw.updateCountLock.Unlock()
}
