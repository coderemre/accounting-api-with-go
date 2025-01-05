package models

import "time"

type AuditLog struct {
	ID         int64     `json:"id"`          
	EntityType string    `json:"entity_type"` 
	EntityID   int64     `json:"entity_id"`   
	Action     string    `json:"action"`      
	Details    string    `json:"details"`   
	CreatedAt  time.Time `json:"created_at"`  
}