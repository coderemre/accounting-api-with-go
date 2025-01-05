package models

import "time"

type Balance struct {
	UserID       int64     `json:"user_id"`        
	Amount       float64   `json:"amount"`     
	LastUpdatedAt time.Time `json:"last_updated_at"`
}