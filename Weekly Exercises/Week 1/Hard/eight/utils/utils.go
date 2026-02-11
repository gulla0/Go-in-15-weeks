package utils

import (
	"fmt"
	"time"
)

func FirstResult(a, b <-chan int, timeout time.Duration) (int, error) {
	select {
	case v := <-a:
		fmt.Println("a received", v)
		return v, nil
	case v := <-b:
		fmt.Println("b received", v)
		return v, nil
	case <-time.After(timeout):
		fmt.Println("timeout")
		return 0, fmt.Errorf("timeout")
	}
}
