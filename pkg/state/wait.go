package state

import (
	"fmt"

	"github.com/prairir/imacry/pkg/config"
	"math/rand"
	"time"
)

// state.Wait: the wait state which waits for either an error or a signal
// from the heartbeat goroutine. if it gets a signal, move to next nextState.
// if gets an error, return it.
//
// params: the next state
// returns: error
func Wait(nextState config.State) error {
	// random seed
	// we wanna stay kinda sorta hidden so not the best idea
	// to read from /dev/urandom
	rand.Seed(time.Now().UnixNano())

	// loop til we get exit signal
	for {
		// a random sleep time to stop the heartbeat
		// from going every second
		// or being easily detectable
		//
		// it will randomly sleep between 0s and 1m
		n := rand.Intn(60)
		time.Sleep(time.Duration(n) * time.Second)

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
