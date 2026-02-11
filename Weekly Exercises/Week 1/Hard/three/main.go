package main

import (
	"Hard/three/utils"
	"fmt"
)

func main() {
	ch := make(chan string)
	go utils.ProduceLogs(ch, 10)
	count := utils.ConsumeLogs(ch)
	fmt.Println(count)
}
