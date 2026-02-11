package utils

import "fmt"

type Ledger struct {
	balances map[string]int
}

func NewLedger() *Ledger {
	return &Ledger{balances: make(map[string]int)}
}

func (l *Ledger) Deposit(acct string, amount int) error {
	if amount <= 0 {
		return fmt.Errorf("amount must be greater than 0")
	}
	l.balances[acct] += amount
	return nil
}

func (l *Ledger) Withdraw(acct string, amount int) error {
	_, ok := l.balances[acct]
	if !ok {
		return fmt.Errorf("account does not exist")
	}
	if amount <= 0 {
		return fmt.Errorf("amount must be greater than 0")
	}
	if amount > l.balances[acct] {
		return fmt.Errorf("amount must be less than the balance")
	}
	l.balances[acct] -= amount
	return nil
}

func (l *Ledger) Balance(acct string) (int, bool) {
	_, ok := l.balances[acct]
	balance, ok := l.balances[acct]
	return balance, ok
}
