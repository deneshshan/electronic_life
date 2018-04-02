package state_test

import (
	"testing"

	st "github.com/deneshshan/electronic_life/engine/state"
)

func setupStateReader(t *testing.T, width, height int) (*st.GameStateReader, *st.State, func()) {
	state := st.NewState(width, height)

	reader := st.NewGameStateReader(state)
	return reader, state, func() {
	}
}

func TestSnapshotCanGetSnapshot(t *testing.T) {
}

// Integration
func TestGetsWholeStateSnapShot(t *testing.T) {
}

// Integration
func TestGetsAllEmptySpaceSnapshot(t *testing.T) {
}

// Integration
func TestGetsLocalisedSnapshot(t *testing.T) {
}
