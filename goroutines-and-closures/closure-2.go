package main

import (
	"fmt"
	"sync"
)

// Case 1: they all print 4, this is because by the time go routines got chance to run,
//         the for loop had completed and the value of i has already been increamented to value 4

// func main() {
// 	var wg sync.WaitGroup

// 	for i := 0; i <= 3; i++ {
// 		wg.Add(1)
// 		go func() {
// 			defer wg.Done()
// 			fmt.Println(i)
// 		}()
// 	}

// 	wg.Wait()

// }

// Case 2: to fix case 1, we need to pass the variable i as a parameter to goroutine,
//         which is create a new variable
func main() {
	var wg sync.WaitGroup

	for i := 0; i <= 3; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			fmt.Println(i)
		}(i)
	}

	wg.Wait()

}
