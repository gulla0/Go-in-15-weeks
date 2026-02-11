package utils

import "fmt"

type Item struct {
	ID    int
	Name  string
	Count int
}

type Op struct {
	ID    int
	Delta int
}

func ApplyOps(items []Item, ops []Op) (updated []Item, applied int, err error) {
	if len(items) == 0 {
		return nil, 0, fmt.Errorf("items is empty")
	}

	if len(ops) == 0 {
		return items, 0, nil
	}

	for i := range ops {
		for j := range items {
			if items[j].ID == ops[i].ID {
				items[j].Count += ops[i].Delta
				if items[j].Count < 0 {
					return nil, applied, fmt.Errorf("item with id %d has negative count after operation", items[j].ID)
				}
				applied++
				break
			}

			if j == len(items)-1 {
				missingId := ops[i].ID
				return nil, applied, fmt.Errorf("item with id %d not found", missingId)
			}
		}
	}
	return items, applied, nil
}
