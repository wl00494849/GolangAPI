package server

import (
	"fmt"
)

var channel chan string
var channel1 chan string
var sli []string

func ChanTest(msg []string) []string {
	channel = make(chan string)
	channel1 = make(chan string)
	sli = make([]string, 0)

	go chanOutPut()
	go chanPrint()

	for _, ch := range msg {
		fmt.Println(ch)
		channel <- ch
	}

	return sli
}

func chanOutPut() {
	for ch := range channel {
		sli = append(sli, "Say "+ch)
		channel1 <- ch
	}
}

func chanPrint() {
	for ch := range channel1 {
		fmt.Println("Add " + ch)
	}
}
