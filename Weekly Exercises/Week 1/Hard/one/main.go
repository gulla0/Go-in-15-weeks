package main

import (
	"Week1/hard/utils"
	"fmt"
)

func main() {
	// =============================================================================
	// Uncomment exactly ONE block below to run that test. Expected outcomes in comments.
	// =============================================================================

	// -----------------------------------------------------------------------------
	// 1. SUCCESS: Basic Update — ID:1 Count:10 + Delta:-2 => Count:8, applied:1, err:nil
	// -----------------------------------------------------------------------------
	items := []utils.Item{{ID: 1, Name: "Item 1", Count: 10}}
	ops := []utils.Op{{ID: 1, Delta: -2}}
	items, applied, err := utils.ApplyOps(items, ops)
	fmt.Printf("Basic Update: items=%v applied=%d err=%v\n", items, applied, err)

	// -----------------------------------------------------------------------------
	// 2. SUCCESS: Multiple Ops — ID:1 +5, ID:2 -1 => ID:1:10, ID:2:4, applied:2
	// -----------------------------------------------------------------------------
	// items := []utils.Item{
	// 	{ID: 1, Name: "A", Count: 5},
	// 	{ID: 2, Name: "B", Count: 5},
	// }
	// ops := []utils.Op{{ID: 1, Delta: 5}, {ID: 2, Delta: -1}}
	// items, applied, err := utils.ApplyOps(items, ops)
	// fmt.Printf("Multiple Ops: items=%v applied=%d err=%v\n", items, applied, err)

	// -----------------------------------------------------------------------------
	// 3. SUCCESS: Zero Result — ID:1 Count:5 + Delta:-5 => Count:0, applied:1
	// -----------------------------------------------------------------------------
	// items := []utils.Item{{ID: 1, Name: "Item 1", Count: 5}}
	// ops := []utils.Op{{ID: 1, Delta: -5}}
	// items, applied, err := utils.ApplyOps(items, ops)
	// fmt.Printf("Zero Result: items=%v applied=%d err=%v\n", items, applied, err)

	// -----------------------------------------------------------------------------
	// 4. SUCCESS: Empty Ops — no ops => Count unchanged (5), applied:0, err:nil
	// -----------------------------------------------------------------------------
	// items := []utils.Item{{ID: 1, Name: "Item 1", Count: 5}}
	// ops := []utils.Op{}
	// items, applied, err := utils.ApplyOps(items, ops)
	// fmt.Printf("Empty Ops: items=%v applied=%d err=%v\n", items, applied, err)
	// fmt.Println(items, applied, err)

	// -----------------------------------------------------------------------------
	// 5. ERROR: Empty Items — any op => error "items is empty", applied:0
	// -----------------------------------------------------------------------------
	// items := []utils.Item{}
	// ops := []utils.Op{{ID: 1, Delta: 1}}
	// items, applied, err := utils.ApplyOps(items, ops)
	// fmt.Printf("Empty Items: items=%v applied=%d err=%v\n", items, applied, err)

	// -----------------------------------------------------------------------------
	// 6. ERROR: Missing ID — items have ID:1 only, op ID:99 => "not found", applied:0
	// -----------------------------------------------------------------------------
	// items := []utils.Item{{ID: 1, Name: "Item 1", Count: 5}}
	// ops := []utils.Op{{ID: 99, Delta: 1}}
	// items, applied, err := utils.ApplyOps(items, ops)
	// fmt.Printf("Missing ID: items=%v applied=%d err=%v\n", items, applied, err)

	// -----------------------------------------------------------------------------
	// 7. ERROR: Negative Stock — ID:1 Count:2 + Delta:-5 => "negative count", applied:0
	// -----------------------------------------------------------------------------
	// items := []utils.Item{{ID: 1, Name: "Item 1", Count: 2}}
	// ops := []utils.Op{{ID: 1, Delta: -5}}
	// items, applied, err := utils.ApplyOps(items, ops)
	// fmt.Printf("Negative Stock: items=%v applied=%d err=%v\n", items, applied, err)

	// -----------------------------------------------------------------------------
	// 8. ERROR: Partial Success — Op1: ID:1 -5 (ok), Op2: ID:2 -10 (would go negative) => error, applied:1
	// -----------------------------------------------------------------------------
	// items := []utils.Item{
	// 	{ID: 1, Name: "A", Count: 10},
	// 	{ID: 2, Name: "B", Count: 5},
	// }
	// ops := []utils.Op{{ID: 1, Delta: -5}, {ID: 2, Delta: -10}}
	// items, applied, err := utils.ApplyOps(items, ops)
	// fmt.Printf("Partial Success: items=%v applied=%d err=%v\n", items, applied, err)
}
