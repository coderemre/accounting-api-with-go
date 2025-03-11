package services

import (
	"errors"

	"accounting-api-with-go/internal/models"
	"accounting-api-with-go/internal/repositories"
)

type TransactionService struct {
	TransactionRepo *repositories.TransactionRepository
	BalanceService  *BalanceService
}

func NewTransactionService(txRepo *repositories.TransactionRepository, balanceService *BalanceService) *TransactionService {
	return &TransactionService{TransactionRepo: txRepo, BalanceService: balanceService}
}

func (s *TransactionService) ProcessTransaction(userID int64, amount float64) (*models.Transaction, error) {
	if amount == 0 {
		return nil, errors.New("transaction amount cannot be zero")
	}

	err := s.BalanceService.UpdateBalance(userID, amount)
	if err != nil {
		return nil, err
	}

	transaction := &models.Transaction{
		UserID: userID,
		Amount: amount,
		Type:   "credit",
	}
	if amount < 0 {
		transaction.Type = "debit"
	}

	err = s.TransactionRepo.CreateTransaction(transaction)
	if err != nil {
		// Hata olursa rollback yap
		_ = s.BalanceService.UpdateBalance(userID, -amount)
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
	transaction, err := s.TransactionRepo.GetTransactionByID(transactionID)
	if err != nil {
		return nil, err
	}

	return transaction, nil
}

func (s *TransactionService) GetTransactionHistory(userID int64) ([]models.Transaction, error) {
	return s.TransactionRepo.GetTransactionHistory(userID)
}