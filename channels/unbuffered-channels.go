package main

import (
	"fmt"
	"time"
)

func main() {
	ch := make(chan string)

	go helper(ch)

	ch <- "raja"
	ch <- "rammohan"
	ch <- "roy"
	ch <- "from"
	ch <- "india"        // when the last value "india" is read by helper goroutine, main goroutine 
	                     // becomes non blocking and completes execution even before the read value is printed by helper goroutine
	// close(ch)
	time.Sleep(time.Millisecond)      // helps to block main goroutine so helper goroutine prints value before main terminates
}


func helper(ch chan string){
	for {
		fmt.Println("value read from channel is: "+ <-ch)
	}
}