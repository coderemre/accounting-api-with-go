package repositories

import "accounting-api-with-go/internal/models"

type BalanceRepository interface {
	GetBalance(userID int64) (*models.Balance, error)
	UpdateBalance(b *models.Balance) error
}
