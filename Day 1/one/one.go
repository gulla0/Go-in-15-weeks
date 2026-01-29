package main

import "fmt"

func main() {

	// x := 15
	y := 20

	// for x <= y {
	// 	fmt.Println(x)
	// 	if x < y {
	// 		fmt.Println("x is less than y")
	// 	} else {
	// 		fmt.Println("x is equal to y")
	// 	}
	// 	x++
	// }

	for x := 15; x <= y; x++ {
		fmt.Println(x)
		if x < y {

			fmt.Println("x is less than y")

		} else {

			fmt.Println("x is equal to y")

		}
	}
}
