package utils

import "fmt"

type Item struct {
	Name  string
	Count int
}

func BadApplyDelta(items *[]Item, delta int) {
	for i, item := range *items {
		// item.Count += delta
		(*items)[i].Count += delta
		fmt.Println(i, item, item.Count)
		fmt.Println(i, (*items)[i], (*items)[i].Count)
	}
	fmt.Println(items, "items")
}
