package services

import (
	"context"
	"fmt"
	"strconv"
	"sync"
	"time"

	"accounting-api-with-go/internal/cache"
	"accounting-api-with-go/internal/models"
	"accounting-api-with-go/internal/repositories"
)

type BalanceService struct {
	BalanceRepo *repositories.BalanceRepository
	Cache       cache.Cache
	mu          sync.Mutex
}

func NewBalanceService(balanceRepo *repositories.BalanceRepository, cache cache.Cache) *BalanceService {
	return &BalanceService{BalanceRepo: balanceRepo, Cache: cache}
}

func (s *BalanceService) GetBalanceHistory(userID int64, currency string) ([]models.BalanceHistory, error) {
	return s.BalanceRepo.GetBalanceHistory(userID, currency)
}

func (s *BalanceService) GetCurrentBalance(userID int64, currency string) (float64, error) {
	cacheKey := fmt.Sprintf("balance:user:%d:currency:%s", userID, currency)
	ctx := context.Background()
	cachedBalance, err := s.Cache.Get(ctx, cacheKey)
	if err == nil {
		balance, err := strconv.ParseFloat(cachedBalance, 64)
		if err == nil {
			return balance, nil
		}
	}

	balance, err := s.BalanceRepo.GetBalance(userID, currency)
	if err != nil {
		return 0, err
	}

	err = s.Cache.Set(ctx, cacheKey, fmt.Sprintf("%f", balance.Amount), 10*time.Minute)
	if err != nil {
		return 0, err
	}

	return balance.Amount, nil
}

func (s *BalanceService) GetBalanceAtTime(userID int64, currency string, atTime time.Time) (float64, error) {
	balance, err := s.BalanceRepo.GetBalanceAtTime(userID, currency, atTime)
	if err != nil {
		return 0, err
	}

	return balance, nil
}

func (s *BalanceService) UpdateBalance(userID int64, currency string, newAmount float64) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	err := s.BalanceRepo.UpdateBalance(userID, currency, newAmount)
	if err != nil {
		return err
	}

	cacheKey := fmt.Sprintf("balance:user:%d:currency:%s", userID, currency)
	ctx := context.Background()
	err = s.Cache.Delete(ctx, cacheKey)
	if err != nil {
		return err
	}
	return nil
}