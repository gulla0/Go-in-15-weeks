package main

import (
	"Hard/seven/utils"
	"fmt"
)

func main() {
	jobs := []utils.Job{
		{ID: 1, Payload: 10},
		{ID: 2, Payload: 20},
		{ID: 3, Payload: 30},
	}
	results := utils.RunWorkers(jobs, 2)
	fmt.Println(results)
}
