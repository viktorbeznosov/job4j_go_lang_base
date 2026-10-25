package base

import (
	"fmt"
	"math"
)

type Account struct {
    Balance int64
}

func (a *Account) Deposit(amount int64) error {
    if (amount <= 0) {
        return fmt.Errorf("the amount must be greater than 0")
    }

    if a.Balance > math.MaxInt64 - amount {
        return fmt.Errorf("result is outside int64 range")
    }

    a.Balance += amount

    return nil
}

func (a *Account) Withdraw(amount int64) error {
    if (amount <= 0) {
        return fmt.Errorf("the amount must be greater than 0")
    }
    if (a.Balance < amount) {
        return fmt.Errorf("the write‑off amount must not exceed the balance")
    }

    a.Balance -= amount

    return nil
}



