package web

import (
	"net/http"

	"github.com/apex/log"
	"github.com/gorilla/websocket"
	"github.com/prairir/imacry/cc-server/pkg/config"
	"github.com/prairir/imacry/cc-server/pkg/handler"
)

// default options
var upgrader = websocket.Upgrader{}

func wsHandler(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Errorf("init error: %s", err)
		return
	}
	log.Infof("New Connection from %s", r.RemoteAddr)

	// ws event loop dispatcher
	for {
		mt, message, err := conn.ReadMessage()
		if err != nil {
			log.Errorf("fatal read error: %s", err)
			break
		}

		// if connection closes
		if mt == websocket.CloseNormalClosure {
			break
		}

		// init handler
		if mt == websocket.TextMessage && string(message[:5]) == "init:" {
			err = handler.Init(message, conn)
			if err != nil {
				log.Errorf("fatal init error: %s", err)
				break
			}
		}

		// Heart Beat handler
		// Heart Beat doubles as a trigger event
		if mt == websocket.TextMessage && string(message[:3]) == "hb:" {
			err = handler.HeartBeat(message, conn)
			if err != nil {
				log.Errorf("fatal init error: %s", err)
				break
			}
		}

	}

	// closing the websocket
	err = conn.Close()
	if err != nil {
		log.Errorf("fatal closing error: %s", err)
	}
}

func Run() {
	http.HandleFunc("/", wsHandler)

	log.Infof("Listen and Serve on port %s", config.Config.Port)
	log.Fatalf("%s", http.ListenAndServe(":"+config.Config.Port, nil))

}
