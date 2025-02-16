package repositories

import "accounting-api-with-go/internal/models"

type TransactionRepository interface {
	Save(transaction *models.Transaction) error
	FindByID(id int64) (*models.Transaction, error)
	UpdateStatus(id int64, status models.TransactionStatus) error
}