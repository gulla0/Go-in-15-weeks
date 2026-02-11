package main

import (
	"Hard/eight/utils"
	"time"
)

func main() {
	a := make(chan int)
	b := make(chan int)
	go func() {
		time.Sleep(2 * time.Second)
		a <- 1
	}()
	go func() {
		time.Sleep(3 * time.Second)
		b <- 2
	}()
	utils.FirstResult(a, b, 2*time.Second)

}
