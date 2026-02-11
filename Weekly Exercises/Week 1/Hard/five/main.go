package main

import (
	"Hard/five/utils"
	"fmt"
)

func main() {
	// Uncomment one block at a time to run.

	// --- Regular: deposit and check balance ---
	// l := utils.NewLedger()
	// _ = l.Deposit("alice", 100)
	// bal, ok := l.Balance("alice")
	// fmt.Printf("alice balance: %d, ok: %v\n", bal, ok) // 100, true

	// --- Regular: deposit, withdraw, then balance ---
	// l := utils.NewLedger()
	// _ = l.Deposit("bob", 200)
	// _ = l.Withdraw("bob", 50)
	// bal, ok := l.Balance("bob")
	// fmt.Printf("bob balance: %d, ok: %v\n", bal, ok) // 150, true

	// --- Regular: multiple accounts ---
	// l := utils.NewLedger()
	// _ = l.Deposit("alice", 100)
	// _ = l.Deposit("bob", 200)
	// _ = l.Withdraw("alice", 30)
	// a, oa := l.Balance("alice")
	// b, ob := l.Balance("bob")
	// fmt.Printf("alice: %d (%v), bob: %d (%v)\n", a, oa, b, ob) // alice: 70 (true), bob: 200 (true)

	// --- Edge: deposit amount <= 0 ---
	l := utils.NewLedger()
	err := l.Deposit("alice", 0)
	fmt.Printf("Deposit(0): %v\n", err) // amount must be greater than 0
	err = l.Deposit("alice", -10)
	fmt.Printf("Deposit(-10): %v\n", err) // amount must be greater than 0

	// --- Edge: withdraw from non-existent account ---
	// l := utils.NewLedger()
	// err := l.Withdraw("nobody", 50)
	// fmt.Printf("Withdraw(nobody): %v\n", err) // account does not exist

	// --- Edge: withdraw more than balance ---
	// l := utils.NewLedger()
	// _ = l.Deposit("alice", 50)
	// err := l.Withdraw("alice", 100)
	// fmt.Printf("Withdraw(100 from 50): %v\n", err) // amount must be less than the balance

	// --- Edge: withdraw amount <= 0 ---
	// l := utils.NewLedger()
	// _ = l.Deposit("alice", 100)
	// err := l.Withdraw("alice", 0)
	// fmt.Printf("Withdraw(0): %v\n", err) // amount must be greater than 0
	// err = l.Withdraw("alice", -5)
	// fmt.Printf("Withdraw(-5): %v\n", err) // amount must be greater than 0

	// --- Edge: Balance on non-existent account (comma-ok) ---
	// l := utils.NewLedger()
	// _ = l.Deposit("alice", 100)
	// bal, ok := l.Balance("alice")
	// fmt.Printf("alice: %d, ok: %v\n", bal, ok) // 100, true
	// bal, ok = l.Balance("nobody")
	// fmt.Printf("nobody: %d, ok: %v\n", bal, ok) // 0, false
}
