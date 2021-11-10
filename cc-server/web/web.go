package web

import (
	"fmt"
	"github.com/apex/log"
	"github.com/gorilla/websocket"
	"net/http"
)

// default options
var upgrader = websocket.Upgrader{}

func initRoute(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	fmt.Println("hello")
	if err != nil {
		log.Errorf("init error: %s", err)
	}
	defer c.Close()
}

func Run() {
	http.HandleFunc("/init", initRoute)

	log.Fatalf("%s", http.ListenAndServe(":80", nil))

}
