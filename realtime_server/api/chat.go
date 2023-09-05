package api

import (
	"log"
	"net/http"
	socket "placio-realtime/pkg/websocket"
)

func ServeWs(hub *socket.Hub, w http.ResponseWriter, r *http.Request) {
	conn, err := socket.Upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	connection := socket.NewConnection(conn, hub)
	hub.Register <- connection
	go connection.Writer()
	go connection.Reader()
}
