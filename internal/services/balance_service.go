package services

import (
	"errors"
	"sync"
	"time"

	"accounting-api-with-go/internal/models"
	"accounting-api-with-go/internal/repositories"
)

type BalanceService struct {
	BalanceRepo *repositories.BalanceRepository
	mu          sync.Mutex
}

func NewBalanceService(balanceRepo *repositories.BalanceRepository) *BalanceService {
	return &BalanceService{BalanceRepo: balanceRepo}
}

func (s *BalanceService) UpdateBalance(userID int64, amount float64) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	balance, err := s.BalanceRepo.GetBalance(userID)
	if err != nil {
		return err
	}

	if balance.Amount+amount < 0 {
		return errors.New("insufficient funds")
	}

	return s.BalanceRepo.UpdateBalance(userID, balance.Amount+amount)
}

func (s *BalanceService) GetBalanceHistory(userID int64) ([]models.BalanceHistory, error) {
	return s.BalanceRepo.GetBalanceHistory(userID)
}

func (s *BalanceService) GetCurrentBalance(userID int64) (float64, error) {
	balance, err := s.BalanceRepo.GetBalance(userID)
	if err != nil {
		return 0, err
	}
	return balance.Amount, nil
}

func (s *BalanceService) GetBalanceAtTime(userID int64, atTime time.Time) (float64, error) {
	balance, err := s.BalanceRepo.GetBalanceAtTime(userID, atTime)
	if err != nil {
		return 0, err
	}

	return balance, nil
}