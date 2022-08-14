package main

import (
	"fmt"
	"sync"
)

// every program in Golang executes until main function is not terminated.

// since goroutines are forked from main function, it is possible that before finishing of the go routine ,
// main function can complete execution

// there are two solutions to the problem:
// 1. By adding the time.Sleep(time.Second) method at the end of the main function.
//    The program will pause execution for a set duration while blocking the main() function,
//    allowing other goroutines to execute. This solution is not scalable in a live project
//    serving thousands of requests.

// 2. You can use WaitGroup to solve the problem of empty output caused by goroutines in this
//    case. To use this package, import sync.WaitGroup from the standard Golang primitive.
func main() {
	var wg sync.WaitGroup
	wg.Add(2)

	go func() {
		count("sheep")
		defer wg.Done()
	}()

	go func() {
		count("dog")
		defer wg.Done()
	}()

	wg.Wait()
}

func count(thing string) {
	for i := 0; i < 5; i++ {
		fmt.Println(thing)
	}
}

// WaitGroup ships with 3 methods:

// 1) wg.Add(int) method: The code in line 23 (above) indicates the number of available goroutines to wait for.
//                        The integer in the function parameter acts like a counter.

// 2) wg.Wait() method: This method blocks the execution of code in line 35 until the internal counter
//                      reduces to a 0 value.

// 3) wg.Done() method: This method decreases the count parameter in Add(int) by 1.
