package server

import (
	"fmt"
)

var channl chan string
var sli []string

func ChanTest(msg []string) []string {
	channl = make(chan string)
	sli = make([]string, 0)

	go chanlOutPut()

	for _, ch := range msg {
		fmt.Println(ch)
		channl <- ch
	}

	return sli
}

func chanlOutPut() {
	for ch := range channl {
		fmt.Printf("add %s", ch)
		sli = append(sli, "Say "+ch)
	}
}
