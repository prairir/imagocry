package heartbeat

import (
	"fmt"

	"github.com/gorilla/websocket"
	"github.com/prairir/imacry/pkg/config"
)

// heartbeat.HeartBeat:
func HeartBeat() {
	for {
		err := config.Config.Conn.WriteMessage(websocket.TextMessage, []byte("hb:"))
		if err != nil {
			config.Config.HBError <- fmt.Errorf("heartbeat.HeartBeat error: %w", err)
			return
		}

		mt, message, err := config.Config.Conn.ReadMessage()
		if err != nil {
			config.Config.HBError <- fmt.Errorf("heartbeat.HeartBeat error: %w", err)
			return
		}

		fmt.Println(string(message))

		// if its a text message and the message is `hb: 1`
		// close the signal channel and exit
		// else if its a text message and the message is `hb: 0`
		// continue
		if mt == websocket.TextMessage && string(message[:5]) == "hb: 1" {
			close(config.Config.Signal)
			return
		} else if mt == websocket.TextMessage && string(message[:5]) == "hb: 0" {
			continue
		}

	}
}
