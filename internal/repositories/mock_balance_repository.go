package repositories

import (
	"accounting-api-with-go/internal/models"
	"errors"
)

type MockBalanceRepository struct {
	balances map[int64]*models.Balance
}

func NewMockBalanceRepository() *MockBalanceRepository {
	return &MockBalanceRepository{
		balances: make(map[int64]*models.Balance),
	}
}

func (m *MockBalanceRepository) FindByUserID(userID int64) (*models.Balance, error) {
	balance, exists := m.balances[userID]
	if !exists {
		return nil, errors.New("balance not found")
	}
	return balance, nil
}

func (m *MockBalanceRepository) UpdateBalance(userID int64, amount float64) error {
	balance, exists := m.balances[userID]
	if !exists {
		balance = &models.Balance{UserID: userID, Amount: 0}
		m.balances[userID] = balance
	}

	if balance.Amount+amount < 0 {
		return errors.New("insufficient funds")
	}

	balance.Amount += amount
	m.balances[userID] = balance
	return nil
}