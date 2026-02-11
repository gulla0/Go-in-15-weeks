package main

import "Hard/nine/utils"

func main() {
	items := []utils.Item{{Name: "Item 1", Count: 10}, {Name: "Item 2", Count: 20}, {Name: "Item 3", Count: 30}}
	utils.BadApplyDelta(&items, 10)
}
