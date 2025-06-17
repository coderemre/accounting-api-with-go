package models

import "time"

type Balance struct {
	UserID  int64   `json:"user_id"`
	Amount  float64 `json:"amount"`
	Currency string
}

type BalanceHistory struct {
	UserID    int64     `json:"user_id"`
	Amount    float64   `json:"amount"`
	CreatedAt time.Time `json:"created_at"`
	Currency string
}