package models

import (
	"accounting-api-with-go/internal/utils"
	"sync"
)

type Balance struct {
	mu      sync.Mutex
	UserID  int64   `json:"user_id"`
	Amount  float64 `json:"amount"`
}


func (b *Balance) Validate() utils.Message {
	if b.UserID == 0 {
		return utils.ErrUserIDRequired
	}
	if b.Amount < 0 {
		return utils.ErrNegativeBalance
	}
	return ""
}

func (b *Balance) Update(amount float64) utils.Message {
	b.mu.Lock()
	defer b.mu.Unlock()

	if b.Amount+amount < 0 {
		return utils.ErrInsufficientFunds
	}

	b.Amount += amount
	return ""
}

func (b *Balance) GetBalance() float64 {
	b.mu.Lock()
	defer b.mu.Unlock()
	return b.Amount
}