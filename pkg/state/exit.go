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
	err := config.Config.Conn.WriteControl(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""), time.Now().Add(time.Second))
	if err != nil {
		return fmt.Errorf("state.Exit error: %w", err)
	}

	err = config.Config.Conn.Close()
	if err != nil {
		return fmt.Errorf("state.Exit error: %w", err)
	}
	return nil
}
