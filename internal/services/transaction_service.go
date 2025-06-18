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
	"time"
)

type TransactionService struct {
	TransactionRepo      *repositories.TransactionRepository
	BalanceService       *BalanceService
	Cache                cache.Cache
	LimitService         *TransactionLimitService
}

func NewTransactionService(
	txRepo *repositories.TransactionRepository,
	balanceService *BalanceService,
	cache cache.Cache,
	limitService *TransactionLimitService,
) *TransactionService {
	return &TransactionService{
		TransactionRepo: txRepo,
		BalanceService:  balanceService,
		Cache:           cache,
		LimitService:    limitService,
	}
}

func (s *TransactionService) ProcessTransaction(fromUserID int64, toUserID int64, amount float64, transactionType string, currency string) (*models.Transaction, error) {
	if amount == 0 {
		return nil, errors.New("transaction amount cannot be zero")
	}

	if err := s.LimitService.CheckLimit(fromUserID, transactionType, amount); err != nil {
		return nil, err
	}

	if transactionType == "credit" {
		err := s.BalanceService.UpdateBalance(toUserID, currency, amount)
		fromUserID = constants.DEFAULT_SYSTEM_USER_ID
		if err != nil {
			return nil, err
		}
	} else if transactionType == "debit" {
		err := s.BalanceService.UpdateBalance(fromUserID, currency, -amount)
		if err != nil {
			return nil, err
		}
	} else if transactionType == "transfer" {
		err := s.BalanceService.UpdateBalance(fromUserID, currency, -amount)
		if err != nil {
			return nil, err
		}
		err = s.BalanceService.UpdateBalance(toUserID, currency, amount)
		if err != nil {
			_ = s.BalanceService.UpdateBalance(fromUserID, currency, amount)
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
		Currency:   currency,
		Status:     "completed",
		CreatedAt:  time.Now(),
	}

	if err := s.TransactionRepo.CreateTransaction(transaction); err != nil {
		switch transactionType {
		case "credit":
			_ = s.BalanceService.UpdateBalance(toUserID, currency, -amount)
		case "debit":
			_ = s.BalanceService.UpdateBalance(fromUserID, currency, amount)
		case "transfer":
			_ = s.BalanceService.UpdateBalance(toUserID, currency, -amount)
			_ = s.BalanceService.UpdateBalance(fromUserID, currency, amount)
		}
		return nil, err
	}

	return transaction, nil
}
func (s *TransactionService) Credit(ctx context.Context, userID int64, amount float64, currency string) error {
	_, err := s.ProcessTransaction(constants.DEFAULT_SYSTEM_USER_ID, userID, amount, "credit", currency)
	return err
}

func (s *TransactionService) Debit(ctx context.Context, userID int64, amount float64, currency string) error {
	_, err := s.ProcessTransaction(userID, constants.DEFAULT_SYSTEM_USER_ID, amount, "debit", currency)
	return err
}

func (s *TransactionService) Transfer(senderID, receiverID int64, amount float64, currency string) error {
	if amount <= 0 {
		return errors.New("transfer amount must be positive")
	}

	err := s.BalanceService.UpdateBalance(senderID, currency, -amount)
	if err != nil {
		return errors.New("insufficient funds for transfer")
	}

	err = s.BalanceService.UpdateBalance(receiverID, currency, amount)
	if err != nil {
		_ = s.BalanceService.UpdateBalance(senderID, currency, amount)
		return err
	}

	return s.TransactionRepo.CreateTransfer(senderID, receiverID, amount, currency)
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

func (s *TransactionService) ProcessBatchTransactions(ctx context.Context, batch []models.Transaction) ([]*models.Transaction, error) {
	var results []*models.Transaction

	for _, item := range batch {
		var tx *models.Transaction
		var err error

		switch item.Type {
		case "credit":
			tx, err = s.ProcessTransaction(0, item.ToUserID, item.Amount, "credit", item.Currency)
		case "debit":
			tx, err = s.ProcessTransaction(item.FromUserID, 0, item.Amount, "debit", item.Currency)
		case "transfer":
			tx, err = s.ProcessTransaction(item.FromUserID, item.ToUserID, item.Amount, "transfer", item.Currency)
		default:
			continue
		}

		if err == nil {
			results = append(results, tx)
		}
	}

	return results, nil
}