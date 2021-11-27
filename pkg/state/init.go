package state

import (
	"fmt"
	"net/url"

	"github.com/gorilla/websocket"
	"github.com/prairir/imacry/pkg/config"
	"github.com/prairir/imacry/pkg/heartbeat"
)

// state.Init: the initialization state
// It connects to the server, gets the password(if one wasnt provided),
// sets up channel for trigger event, sets up channel for hb error,
// start the heartbeat goroutine, and sets the next state
//
// params: the next state
// returns: error
func Init(nextState config.State) error {
	// the url to connect to the server
	url := url.URL{
		Scheme: "ws",
		Host:   config.Config.Address,
		Path:   "/",
	}

	conn, _, err := websocket.DefaultDialer.Dial(url.String(), nil)
	if err != nil {
		return fmt.Errorf("state.Init error: %w", err)
	}

	// if the password doesnt exist
	// get it from server
	if config.Config.Password == "" {
		err = conn.WriteMessage(websocket.TextMessage, []byte("init:"))
		if err != nil {
			return fmt.Errorf("state.Init error: %w", err)
		}

		mt, message, err := conn.ReadMessage()
		if err != nil {
			return fmt.Errorf("state.Init error: %w", err)
		}

		// if its a text message and starts with `pass:` then make the rest the password
		// else return an error
		fmt.Println(string(message))
		if mt == websocket.TextMessage && string(message[:5]) == "pass:" {
			config.Config.Password = string(message[6:])
		} else {
			return fmt.Errorf("state.Init error: Bad response from server")
		}

	}

	// set the global connection to this connection
	config.Config.Conn = conn

	// signal channel
	// when this is closed, they can read to it meaning that an event
	// happened
	config.Config.Signal = make(chan struct{})

	// error channel
	// when this is populated
	// print error and die
	config.Config.HBError = make(chan error, 1)

	// launch the heartbeat system in a goroutine
	go heartbeat.HeartBeat()

	// after were finished, set the state to the next one
	config.Config.State = nextState
	return nil
}
