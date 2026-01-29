package main

import "fmt"

func main() {
	for x := 1; x < 10; x++ {

		fmt.Println(x)
		if x%2 == 0 {
			fmt.Println("even")
		} else {
			fmt.Println("odd")
		}
	}
}
