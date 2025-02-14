package services

import (
	"accounting-api-with-go/internal/models"
	"sync"
)

type BalanceService struct {
	lock sync.RWMutex
}

func (bs *BalanceService) UpdateBalanceSafely(balance *models.Balance, amount float64) {
	bs.lock.Lock()
	defer bs.lock.Unlock()
	balance.Amount += amount
}

func (bs *BalanceService) TrackBalanceHistory(userID int64) []models.Balance {

	return []models.Balance{}
}

func (bs *BalanceService) OptimizeBalanceCalculation(userID int64) float64 {

	return 0.0
}