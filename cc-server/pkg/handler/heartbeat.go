package handler

import (
	//"fmt"

	"time"

	"github.com/gorilla/websocket"
	"github.com/prairir/imacry/cc-server/pkg/config"
)

// handler.HeartBeat: heart beat and event trigger logic. It reads from
// config.Config.TriggerTime and config.Config.TriggerLength to determine when to
// trigger the event. If the current time is within the trigger period then respond
// with a trigger bit else respond with non triggered bit
//
// response when triggered:     `hb: 1`
// response when not triggered: `hb: 0`
//
// params: takes message and connection
// returns: nothing
func HeartBeat(message []byte, conn *websocket.Conn) {
	// get now
	now := time.Now()

	// format the start time with todays year, month, and day
	// but with the proper hour and minute
	startTime := time.Date(
		now.Year(),
		now.Month(),
		now.Day(),
		config.Config.TriggerTime.Hour(),
		config.Config.TriggerTime.Minute(),
		0, 0, now.Location())

	// end time is start + duration
	endTime := startTime.Add(config.Config.TriggerLength)

	// if current time is within trigger period
	// trigger
	// else dont
	if now.After(startTime) && now.Before(endTime) {
		conn.WriteMessage(websocket.TextMessage, []byte("hb: 1"))
	} else {
		conn.WriteMessage(websocket.TextMessage, []byte("hb: 0"))
	}
}
