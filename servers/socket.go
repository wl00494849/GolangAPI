package server

import (
	"fmt"
	"sync"
	"time"

	"golang.org/x/net/websocket"
)

type channel struct {
	msgChannl chan string
	wg        sync.WaitGroup
}

func Echo(ws *websocket.Conn) {

	var err error
	c := channel{
		msgChannl: make(chan string),
		wg:        sync.WaitGroup{},
	}

	go c.send(ws)

	for {
		var reply string

		if err = websocket.Message.Receive(ws, &reply); err != nil {
			fmt.Println("Nothing Recevie")
			break
		}

		c.msgChannl <- reply
	}
}

func (c *channel) send(ws *websocket.Conn) {
	var err error

	for reply := range c.msgChannl {

		fmt.Println("Success Message : " + reply)
		msg := "OK : " + reply + "(" + time.Now().Local().String() + ")"

		if err = websocket.Message.Send(ws, msg); err != nil {
			fmt.Println("Send error")
			err = nil
		}
	}
}
