package main

import (
	"Hard/ten/internal/service"
	"Hard/ten/internal/store"
	"fmt"
)

func main() {
	// store1 := store.NewInMemoryStore() //Address of the store
	// fmt.Println(store1.Get(1))         // Calling Get on empty function will return false
	// // store1.Put(store.Item{ID: 1, Name: "Item 1", Count: 10})
	// service1 := service.NewService(store1) //Address of the service
	// store1.Put(store.Item{ID: 1, Name: "Item 1", Count: 10})
	// fmt.Println("service:", service1)
	// fmt.Println("store:", store1)
	// error := service1.AddStock(1, 10)
	// if error != nil {
	// 	fmt.Println("Error:", error)
	// }
	// fmt.Println(store1.Get(1))
	// fmt.Println("store:", store1)

	fakeStore1 := store.NewFakeStore()
	service1 := service.NewService(fakeStore1)
	fakeStore1.Put(store.Item{ID: 1, Name: "Item 1", Count: 10})
	fmt.Println("service:", service1)
	fmt.Println("store:", fakeStore1)
	error := service1.AddStock(1, 10)
	if error != nil {
		fmt.Println("Error:", error)
	}
	fmt.Println(fakeStore1.Get(1))
	fmt.Println("store:", fakeStore1)
}
