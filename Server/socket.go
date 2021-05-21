package server

import (
	"fmt"
	"sync"

	"golang.org/x/net/websocket"
)

type channel struct {
	MsgChannl chan string
	wg        sync.WaitGroup
}

func Echo(ws *websocket.Conn) {

	var err error
	c := channel{
		MsgChannl: make(chan string),
		wg:        sync.WaitGroup{},
	}

	go c.send(ws)

	for {

		if err = websocket.Message.Receive(ws, c.MsgChannl); err != nil {
			fmt.Println("Nothing Recevie")
			break
		}
	}
}

func (c *channel) send(ws *websocket.Conn) {
	var err error

	for reply := range c.MsgChannl {
		fmt.Println("Success Message : " + reply)
		msg := "OK : " + reply
		if err = websocket.Message.Send(ws, msg); err != nil {
			fmt.Println("Send error")
			err = nil
		}
	}
}
