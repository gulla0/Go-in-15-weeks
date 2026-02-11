package main

import (
	"fmt"
	"time"
)

func main() {
	// ch := make(chan int, 1)
	ch := make(chan int, 3)
	// ch := make(chan int)
	for i := range 3 {
		go func() {
			ch <- i
			time.Sleep(1 * time.Second)
		}()

		// fmt.Println(<-ch)

	}
	//***DOES NOT WORK***//
	// for range 3 {
	// 	go func() {
	// 		fmt.Println(<-ch)
	// 	}()
	// }
	// --------------------------------

	//**THIS WORKS BUT ORDER IS NOT GUARANTEED**//
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	//--------------------------------

}
