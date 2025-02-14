package models

import "time"


type Transaction struct {
	ID        int64     `json:"id"`
	UserID    int64     `json:"user_id"`
	Amount    float64   `json:"amount"`
	Status    string    `json:"status"`
	CreatedAt time.Time `json:"created_at"`
}

func (t *Transaction) MarkCompleted() {
	t.Status = "completed"
}

func (t *Transaction) MarkFailed() {
	t.Status = "failed"
}
