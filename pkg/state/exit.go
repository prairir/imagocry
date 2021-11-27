package state

import (
	"fmt"
	"time"

	"github.com/gorilla/websocket"
	"github.com/prairir/imacry/pkg/config"
)

// state.Exit: the exit state, closes the cc-server connection
// params: none
// returns: error
func Exit() error {
	// close the socket
	err := config.Config.Conn.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
	if err != nil {
		return fmt.Errorf("state.Exit error: %w", err)
	}

	// wait a second to make sure its finished
	select {
	case <-time.After(time.Second):
	}
	return nil
}
