package server

import (
	"fmt"

	"golang.org/x/net/websocket"
)

func Echo(ws *websocket.Conn) {
	var err error

	for {
		var msg string

		if err = websocket.Message.Receive(ws, msg); err != nil {
			fmt.Println("Nothing Recevie")
			break
		}

		reply := "OK : " + msg
		fmt.Println("Success Message : " + msg)

		if err = websocket.Message.Send(ws, reply); err != nil {
			fmt.Println("Send error")
		}
	}
}
