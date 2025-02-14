package services

import "accounting-api-with-go/internal/models"

type TransactionService interface {
	ProcessTransaction(t *models.Transaction) error
}