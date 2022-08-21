package main

import (
	"fmt"
)

func genMsg(ch1 chan<- string) {
	//send message on ch1
	ch1 <- "data from genMsg"
}

func relayMsg(ch1 <-chan string, ch2 chan<- string) {
	//receive message on ch1
	m := <-ch1
	//send it on ch2
	ch2 <- m
}

func main() {
	// create ch1 and ch2
	ch1 := make(chan string)
	ch2 := make(chan string)


	// spin goroutine genMsg and relayMsg
	go genMsg(ch1)

	go relayMsg(ch1,ch2)

	// receive msg on ch2
	v := <- ch2

	fmt.Println(v)
}