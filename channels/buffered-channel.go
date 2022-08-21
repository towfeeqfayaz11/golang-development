package main

import (
	"fmt"
)

// channel ch1 has buffer capacity of 3. That means it can hold 3 values, which is does at line no. 16 but since 
// the buffer is not overflowing (as we didn’t push any new value), the main goroutine will not block and program exists.

// func main() {
// 	fmt.Println("main() started")
// 	ch1 := make(chan string, 3)
// 	go helper(ch1)
// 	ch1 <- "sa"
// 	ch1 <- "re"
// 	ch1 <- "ga"
// 	fmt.Println("main() finished")
// }




// as now a filled buffer gets the push by ch <- "ma" send operation, main goroutine blocks and helper goroutine 
// drain out all the values.
func main(){
	fmt.Println("main() started")
	ch := make(chan string, 3)
	go helper(ch)
	ch <- "sa"
	ch <- "re"
	ch <- "ga"
	ch <- "ma"
	fmt.Println("main() finished")

}

func helper(ch chan string){

	for {
		fmt.Println("The value read is: "+ <-ch)
	}
}