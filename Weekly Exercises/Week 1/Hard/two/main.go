package main

import (
	"Hard/two/utils"
	"fmt"
)

func main() {
	map1 := utils.MapCounter{}
	fakeMap1 := utils.FakeCounter{}
	utils.RecordLogin(&map1, "user1")
	fmt.Println(map1.Get("login: user1"))
	utils.RecordLogin(&fakeMap1, "user2")
	fmt.Println(fakeMap1.Get("login: user2"))
}
