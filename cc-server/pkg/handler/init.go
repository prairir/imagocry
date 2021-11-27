package handler

import (
	"fmt"

	"github.com/gorilla/websocket"
	"github.com/prairir/imacry/cc-server/pkg/config"
)

func Init(message []byte, conn *websocket.Conn) error {
	err := conn.WriteMessage(websocket.TextMessage, []byte("pass: "+config.Config.Password))
	if err != nil {
		return fmt.Errorf("state.Init error: %w", err)
	}
	return nil
}
