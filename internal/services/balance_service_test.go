package services

import (
	"accounting-api-with-go/internal/repositories"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Mock Balance Service
type MockBalanceService struct {
	repo repositories.BalanceRepository
}

func (s *MockBalanceService) GetBalance(userID int64) (float64, error) {
	balance, err := s.repo.FindByUserID(userID)
	if err != nil {
		return 0, err
	}
	return balance.Amount, nil
}

func (s *MockBalanceService) UpdateBalance(userID int64, amount float64) error {
	return s.repo.UpdateBalance(userID, amount)
}

// Unit Test
func TestBalanceService(t *testing.T) {
	mockRepo := repositories.NewMockBalanceRepository() 
	service := &MockBalanceService{repo: mockRepo}

	balance, err := service.GetBalance(1)
	assert.NotNil(t, err)

	err = service.UpdateBalance(1, 500)
	assert.Nil(t, err)

	balance, err = service.GetBalance(1)
	assert.Nil(t, err)
	assert.Equal(t, float64(500), balance)

	err = service.UpdateBalance(1, 100)
	assert.Nil(t, err)

	balance, err = service.GetBalance(1)
	assert.Nil(t, err)
	assert.Equal(t, float64(600), balance)

	err = service.UpdateBalance(1, -700)
	assert.NotNil(t, err)
}