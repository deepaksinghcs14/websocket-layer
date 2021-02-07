package main

import (
	"fmt"
	"github.com/gorilla/websocket"
	"net/http"
)

var wsUpgrade = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var connections = make(map[string]*websocket.Conn)

func wsHandler(w http.ResponseWriter, r *http.Request, clientId string) {
	conn, err := wsUpgrade.Upgrade(w, r, nil)
	connections[clientId] = conn
	if err != nil {
		fmt.Println("Failed to set websocket upgrade: %+v", err)
		return
	}

	for {
		t, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}
		conn.WriteMessage(t, msg)
	}
}

func sendMessage(message []byte, clientId string) {
	conn := connections[clientId]
	if conn != nil {
		conn.WriteMessage(1, message)
	}
}
