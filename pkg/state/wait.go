package state

import (
	"fmt"

	"github.com/prairir/imacry/pkg/config"
)

// state.Wait: the wait state which waits for either an error or a signal
// from the heartbeat goroutine. if it gets a signal, move to next nextState.
// if gets an error, return it.
//
// params: the next state
// returns: error
func Wait(nextState config.State) error {
	// loop over forever
	for {
		// if you can read from signal(like when its closed)
		// set the next state and exit
		//
		// if you can read an error
		// return the error
		//
		// if you can do either, try again
		select {
		case <-config.Config.Signal:
			config.Config.State = nextState
			return nil
		case err := <-config.Config.HBError:
			return fmt.Errorf("state.Wait error: %w", err)
		default:
			continue
		}
	}
}
