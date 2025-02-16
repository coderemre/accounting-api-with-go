package repositories

import (
	"accounting-api-with-go/internal/models"
	"errors"
)

type MockTransactionRepository struct {
	transactions map[int64]*models.Transaction
}

func NewMockTransactionRepository() *MockTransactionRepository {
	return &MockTransactionRepository{
		transactions: make(map[int64]*models.Transaction),
	}
}

func (m *MockTransactionRepository) Save(transaction *models.Transaction) error {
	if transaction.ID == 0 {
		transaction.ID = int64(len(m.transactions) + 1)
	}
	m.transactions[transaction.ID] = transaction
	return nil
}

func (m *MockTransactionRepository) FindByID(id int64) (*models.Transaction, error) {
	transaction, exists := m.transactions[id]
	if !exists {
		return nil, errors.New("transaction not found")
	}
	return transaction, nil
}

func (m *MockTransactionRepository) UpdateStatus(id int64, status models.TransactionStatus) error {
	transaction, exists := m.transactions[id]
	if !exists {
		return errors.New("transaction not found")
	}
	transaction.Status = status
	return nil
}