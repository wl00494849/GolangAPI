package server

import (
	"fmt"
	"sync"
)

type channels struct {
	channel  chan string
	channel1 chan string
	wg       sync.WaitGroup
	sli      []string
}

func ChanTest(msg []string) []string {
	c := channels{
		channel:  make(chan string),
		channel1: make(chan string),
		wg:       sync.WaitGroup{},
		sli:      make([]string, 0),
	}

	c.wg.Add(len(msg))
	go c.chanOutPut()
	go c.chanPrint()

	for _, ch := range msg {
		fmt.Println(ch)
		c.channel <- ch
	}

	c.wg.Wait()
	return c.sli
}

func (c *channels) chanOutPut() {
	for ch := range c.channel1 {
		c.sli = append(c.sli, "Say "+ch)
		c.wg.Done()
	}
}

func (c *channels) chanPrint() {
	for ch := range c.channel {
		fmt.Println("Add " + ch)
		c.channel1 <- ch
	}
}
