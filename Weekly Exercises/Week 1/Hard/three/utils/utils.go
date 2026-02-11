package utils

import "fmt"

func ProduceLogs(out chan<- string, n int) {
	for i := range n {
		out <- fmt.Sprintf("log #%d", i)
	}
	close(out)
}

func ConsumeLogs(in <-chan string) int {
	count := 0
	for v := range in {
		fmt.Println(v)
		count++
	}
	return count
}
