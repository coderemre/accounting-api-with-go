package models

import (
	"errors"
	"sync"
	"time"
)

type Balance struct {
	UserID    int64     `json:"user_id"`
	Amount    float64   `json:"amount"`
	lock      sync.RWMutex
	UpdatedAt time.Time `json:"updated_at"`
}

func (b *Balance) Deposit(amount float64) {
	b.lock.Lock()
	defer b.lock.Unlock()
	b.Amount += amount
	b.UpdatedAt = time.Now()
}

func (b *Balance) Withdraw(amount float64) error {
	b.lock.Lock()
	defer b.lock.Unlock()
	if b.Amount < amount {
		return errors.New("insufficient funds")
	}
	b.Amount -= amount
	b.UpdatedAt = time.Now()
	return nil
}