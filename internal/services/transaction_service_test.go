package services

import (
	"accounting-api-with-go/internal/models"
	"accounting-api-with-go/internal/repositories"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockTransactionService struct {
	repo repositories.TransactionRepository
}

func (s *MockTransactionService) CreateTransaction(tx *models.Transaction) error {
	return s.repo.Save(tx)
}

func (s *MockTransactionService) GetTransactionByID(id int64) (*models.Transaction, error) {
	return s.repo.FindByID(id)
}

func (s *MockTransactionService) CompleteTransaction(id int64) error {
	return s.repo.UpdateStatus(id, models.StatusCompleted)
}

func (s *MockTransactionService) FailTransaction(id int64) error {
	return s.repo.UpdateStatus(id, models.StatusFailed)
}

func TestTransactionService(t *testing.T) {
	mockRepo := repositories.NewMockTransactionRepository()
	service := &MockTransactionService{repo: mockRepo}

	tx := &models.Transaction{UserID: 1, Amount: 100, Status: models.StatusPending}
	err := service.CreateTransaction(tx)
	assert.Nil(t, err)

	foundTx, err := service.GetTransactionByID(1)
	assert.Nil(t, err)
	assert.Equal(t, int64(1), foundTx.UserID)

	err = service.CompleteTransaction(1)
	assert.Nil(t, err)

	err = service.FailTransaction(2)
	assert.NotNil(t, err)
}