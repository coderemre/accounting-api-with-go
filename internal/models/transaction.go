package models

import (
	"accounting-api-with-go/internal/utils"
	"errors"
	"time"
)

type TransactionStatus string

const (
	StatusPending   TransactionStatus = "pending"
	StatusCompleted TransactionStatus = "completed"
	StatusFailed    TransactionStatus = "failed"
)

type Transaction struct {
	ID        int64             `json:"id"`
	UserID    int64             `json:"user_id"`
	Amount    float64           `json:"amount"`
	Status    TransactionStatus `json:"status"`
	CreatedAt time.Time         `json:"created_at"`
}

func (t *Transaction) Validate() utils.Message {
	if t.UserID == 0 {
		return utils.ErrUserIDRequired
	}
	if t.Amount <= 0 {
		return utils.ErrInvalidAmount
	}
	if t.Status != StatusPending && t.Status != StatusCompleted && t.Status != StatusFailed {
		return utils.ErrInvalidTransactionStatus
	}
	return ""
}

func (t *Transaction) MarkCompleted() error {
	if t.Status != StatusPending {
		return errors.New(string(utils.ErrInvalidTransactionStatus))
	}
	t.Status = StatusCompleted
	return nil
}

func (t *Transaction) MarkFailed() error {
	if t.Status != StatusPending {
		return errors.New(string(utils.ErrInvalidTransactionStatus))
	}
	t.Status = StatusFailed
	return nil
}