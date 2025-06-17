package services

import (
	"accounting-api-with-go/internal/config"
	"accounting-api-with-go/internal/repositories"
	"errors"
	"time"
)

type TransactionLimitService struct {
	TransactionRepo *repositories.TransactionRepository
	Limits          config.GlobalTransactionLimit
}

func NewTransactionLimitService(txRepo *repositories.TransactionRepository, limits config.GlobalTransactionLimit) *TransactionLimitService {
	return &TransactionLimitService{
		TransactionRepo: txRepo,
		Limits:          limits,
	}
}

func (s *TransactionLimitService) CheckLimit(userID int64, txType string, amount float64) error {
	limit := s.Limits

	if limit.MaxAmount > 0 && amount > limit.MaxAmount {
		return errors.New("max amount limit exceeded")
	}

	startTime := time.Now().Add(-24 * time.Hour)

	transactions, err := s.TransactionRepo.GetTransactionsByUserAndType(userID, txType, startTime)
	if err != nil {
		return err
	}

	var totalAmount float64
	for _, tx := range transactions {
		totalAmount += tx.Amount
	}

	if limit.MaxAmount > 0 && (totalAmount+amount) > limit.MaxAmount {
		return errors.New("transaction limit exceeded: total amount")
	}

	if limit.MaxCount > 0 && len(transactions) >= limit.MaxCount {
		return errors.New("transaction limit exceeded: max count")
	}

	return nil
}