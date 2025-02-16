package services

import "accounting-api-with-go/internal/models"

type TransactionService interface {
	CreateTransaction(tx *models.Transaction) error
	GetTransactionByID(id int64) (*models.Transaction, error)
	CompleteTransaction(id int64) error
	FailTransaction(id int64) error
}