package services

import (
	"accounting-api-with-go/internal/cache"
	"accounting-api-with-go/internal/constants"
	"accounting-api-with-go/internal/models"
	"accounting-api-with-go/internal/repositories"
	"context"
	"encoding/json"
	"errors"
	"fmt"
)

type TransactionService struct {
	TransactionRepo *repositories.TransactionRepository
	BalanceService  *BalanceService
	Cache           cache.Cache
}

func NewTransactionService(txRepo *repositories.TransactionRepository, balanceService *BalanceService, cache cache.Cache) *TransactionService {
	return &TransactionService{TransactionRepo: txRepo, BalanceService: balanceService, Cache: cache}
}

func (s *TransactionService) ProcessTransaction(fromUserID int64, toUserID int64, amount float64, transactionType string) (*models.Transaction, error) {
	if amount == 0 {
		return nil, errors.New("transaction amount cannot be zero")
	}

	if transactionType == "credit" {
		err := s.BalanceService.UpdateBalance(toUserID, amount)
		fromUserID = constants.DEFAULT_SYSTEM_USER_ID
		if err != nil {
			return nil, err
		}
	} else if transactionType == "debit" {
		err := s.BalanceService.UpdateBalance(fromUserID, -amount)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, errors.New("invalid transaction type")
	}

	transaction := &models.Transaction{
		FromUserID: fromUserID,
		ToUserID:   toUserID,
		Amount:     amount,
		Type:       transactionType,
		Status:     "completed",
	}

	err := s.TransactionRepo.CreateTransaction(transaction)
	if err != nil {
		if transactionType == "credit" {
			_ = s.BalanceService.UpdateBalance(toUserID, -amount)
		} else if transactionType == "debit" {
			_ = s.BalanceService.UpdateBalance(fromUserID, amount)
		}
		return nil, err
	}

	return transaction, nil
}

func (s *TransactionService) Transfer(senderID, receiverID int64, amount float64) error {
	if amount <= 0 {
		return errors.New("transfer amount must be positive")
	}

	err := s.BalanceService.UpdateBalance(senderID, -amount)
	if err != nil {
		return errors.New("insufficient funds for transfer")
	}

	err = s.BalanceService.UpdateBalance(receiverID, amount)
	if err != nil {
		_ = s.BalanceService.UpdateBalance(senderID, amount)
		return err
	}

	return s.TransactionRepo.CreateTransfer(senderID, receiverID, amount)
}

func (s *TransactionService) GetTransactionByID(transactionID int64) (*models.Transaction, error) {
	var transaction models.Transaction
	cacheKey := fmt.Sprintf("transaction:%d", transactionID)
	ctx := context.Background()

	cached, err := s.Cache.Get(ctx, cacheKey)
	if err == nil {
		if err := json.Unmarshal([]byte(cached), &transaction); err == nil {
			return &transaction, nil
		}
	}

	tx, err := s.TransactionRepo.GetTransactionByID(transactionID)
	if err != nil {
		return nil, err
	}

	data, _ := json.Marshal(tx)
	_ = s.Cache.Set(ctx, cacheKey, string(data), 0)

	return tx, nil
}

func (s *TransactionService) GetTransactionHistory(userID int64) ([]models.Transaction, error) {
	return s.TransactionRepo.GetTransactionHistory(userID)
}