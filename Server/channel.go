package server

import (
	"fmt"
	"sync"
)

var channel = make(chan string)
var channel1 = make(chan string)
var wg = sync.WaitGroup{}
var sli []string

func ChanTest(msg []string) []string {
	sli = make([]string, 0)

	wg.Add(len(msg))
	go chanOutPut()
	go chanPrint()

	for _, ch := range msg {
		fmt.Println(ch)
		channel <- ch
	}

	wg.Wait()
	return sli
}

func chanOutPut() {
	for ch := range channel1 {
		sli = append(sli, "Say "+ch)
		wg.Done()
	}
}

func chanPrint() {
	for ch := range channel {
		fmt.Println("Add " + ch)
		channel1 <- ch
	}
}
