package models

import (
	"errors"
	"time"
)

type User struct {
	ID           int64     `json:"id"`
	Username     string    `json:"username"`
	Email        string    `json:"email"`
	PasswordHash string    `json:"password_hash"`
	Role         string    `json:"role"`
	CreatedAt    time.Time `json:"created_at"`
	UpdatedAt    time.Time `json:"updated_at"`
}

func (u *User) Validate() error {
	if u.Username == "" || u.Email == "" || u.PasswordHash == "" {
		return errors.New("username, email, and password cannot be empty")
	}
    
	return nil
}