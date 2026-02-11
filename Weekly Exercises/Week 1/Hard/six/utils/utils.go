package utils

import (
	"fmt"
	"strconv"
	"strings"
)

type Item struct {
	ID    int
	Name  string
	Count int
}

func ParseItems(lines []string) (items []Item, parsed int, err error) {
	if len(lines) == 0 {
		return nil, 0, fmt.Errorf("no lines proivided")
	}

	parsed = 0

	for i, line := range lines {
		if line == "" {
			return nil, 0, fmt.Errorf("line %d is empty", i)
		}
		if strings.Count(line, ",") != 2 {
			return nil, 0, fmt.Errorf("line %d has too many commas", i)
		}

		parts := strings.Split(line, ",")
		if parts[0] == "" || parts[1] == "" || parts[2] == "" {
			return nil, 0, fmt.Errorf("line %d has a missing field", i)
		}

		id, err := strconv.Atoi(parts[0])
		if err != nil || id < 0 {
			return nil, 0, fmt.Errorf("line %d has an invalid id", i)
		}

		num, err := strconv.Atoi(parts[2])
		if err != nil {
			return nil, 0, fmt.Errorf("line %d has an invalid count", i)
		}
		if num < 0 {
			return nil, 0, fmt.Errorf("line %d has a negative count", i)
		}
		items = append(items, Item{ID: id, Name: parts[1], Count: num})
		parsed++
	}
	return items, parsed, nil
}

// AI answer and better written code:
// func ParseItems(lines []string) ([]Item, error) {
//     if len(lines) == 0 {
//         return nil, fmt.Errorf("no lines provided")
//     }

//     // Pre-allocate slice if you know the size to avoid multiple allocations
//     items := make([]Item, 0, len(lines))

//     for i, line := range lines {
//         parts := strings.Split(line, ",")
//         if len(parts) != 3 {
//             return nil, fmt.Errorf("line %d: expected 3 fields, got %d", i, len(parts))
//         }

//         id, err := strconv.Atoi(parts[0])
//         if err != nil || id < 0 {
//             return nil, fmt.Errorf("line %d: invalid id", i)
//         }

//         count, err := strconv.Atoi(parts[2])
//         if err != nil || count < 0 {
//             return nil, fmt.Errorf("line %d: invalid count", i)
//         }

//         items = append(items, Item{
//             ID:    id,
//             Name:  parts[1],
//             Count: count,
//         })
//     }

//     return items, nil
// }
