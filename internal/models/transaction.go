package models

import "time"

type Transaction struct {
	ID        int64     `json:"id"`          
	FromUserID int64    `json:"from_user_id"`
	ToUserID  int64     `json:"to_user_id"`   
	Amount    float64   `json:"amount"`      
	Type      string    `json:"type"`       
	Status    string    `json:"status"`       
	CreatedAt time.Time `json:"created_at"`
}