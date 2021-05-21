package server

import "golang.org/x/net/websocket"

func Echo(ws *websocket.Conn) {
	var err error

	for {
		var msg string

		if err = websocket.Message.Receive(ws, msg); err != nil {

		}
	}
}
