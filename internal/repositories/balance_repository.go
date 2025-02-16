package repositories

import "accounting-api-with-go/internal/models"

type BalanceRepository interface {
	FindByUserID(userID int64) (*models.Balance, error)
	UpdateBalance(userID int64, amount float64) error
}