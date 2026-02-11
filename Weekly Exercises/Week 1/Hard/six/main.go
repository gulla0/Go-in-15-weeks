package main

import (
	"Hard/six/utils"
	"fmt"
)

func main() {
	// -------------------------------------------------------------------------
	// Uncomment exactly one block below to run that edge case.
	// -------------------------------------------------------------------------

	// --- 1. Happy path: valid lines ---
	lines := []string{"12,Widget,30", "3,Gadget,0"}
	items, parsed, err := utils.ParseItems(lines)
	if err != nil {
		fmt.Println("ParseItems error:", err)
		return
	}
	fmt.Printf("ParseItems ok: items=%+v parsed=%d\n", items, parsed)

	// --- 2. Empty line (error includes line index) ---
	// lines := []string{"12,Widget,30", "", "3,Gadget,0"}
	// items, parsed, err := utils.ParseItems(lines)
	// if err != nil {
	// 	fmt.Println("ParseItems error:", err)
	// 	return
	// }
	// fmt.Printf("ParseItems ok: items=%+v parsed=%d\n", items, parsed)

	// --- 3. Invalid id (non-int) ---
	// lines := []string{"12,Widget,30", "abc,Gadget,5", "3,Gadget,0"}
	// items, parsed, err := utils.ParseItems(lines)
	// if err != nil {
	// 	fmt.Println("ParseItems error:", err)
	// 	return
	// }
	// fmt.Printf("ParseItems ok: items=%+v parsed=%d\n", items, parsed)

	// --- 4. Negative count ---
	// lines := []string{"12,Widget,30", "3,Gadget,-1"}
	// items, parsed, err := utils.ParseItems(lines)
	// if err != nil {
	// 	fmt.Println("ParseItems error:", err)
	// 	return
	// }
	// fmt.Printf("ParseItems ok: items=%+v parsed=%d\n", items, parsed)

	// --- 5. Empty name ---
	// lines := []string{"12,Widget,30", "3,,5"}
	// items, parsed, err := utils.ParseItems(lines)
	// if err != nil {
	// 	fmt.Println("ParseItems error:", err)
	// 	return
	// }
	// fmt.Printf("ParseItems ok: items=%+v parsed=%d\n", items, parsed)

	// --- 6. ParseItemsPartial: first valid, second invalid, third valid ---
	// lines := []string{"12,Widget,30", "bad,Gadget,5", "3,Gadget,0"}
	// items, parsed, err := utils.ParseItemsPartial(lines)
	// fmt.Printf("ParseItemsPartial: items=%+v parsed=%d err=%v\n", items, parsed, err)
}
