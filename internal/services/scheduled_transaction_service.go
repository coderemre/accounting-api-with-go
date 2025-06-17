package services

import (
	"context"
	"log"
	"time"

	"accounting-api-with-go/internal/repositories"
)

type ScheduledTransactionService struct {
	repo              *repositories.ScheduledTransactionRepository
	transactionService *TransactionService
}

func NewScheduledTransactionService(
	repo *repositories.ScheduledTransactionRepository,
	txService *TransactionService,
) *ScheduledTransactionService {
	return &ScheduledTransactionService{
		repo:              repo,
		transactionService: txService,
	}
}

func (s *ScheduledTransactionService) ProcessScheduledTransactions(ctx context.Context) {
	ticker := time.NewTicker(1 * time.Minute)

	go func() {
		for {
			select {
			case <-ticker.C:
				log.Println("Checking for scheduled transactions...")
				s.process(ctx)
			case <-ctx.Done():
				ticker.Stop()
				return
			}
		}
	}()
}

func (s *ScheduledTransactionService) process(ctx context.Context) {
	now := time.Now()
	transactions, err := s.repo.GetPending(now)
	if err != nil {
		log.Printf("Failed to fetch scheduled transactions: %v\n", err)
		return
	}

	for _, st := range transactions {
		if st.Type == "credit" {
			err = s.transactionService.Credit(ctx, st.UserID, st.Amount, st.Currency)
		} else if st.Type == "debit" {
			err = s.transactionService.Debit(ctx, st.UserID, st.Amount, st.Currency)
		}

		if err != nil {
			log.Printf("Transaction failed for scheduled ID %d: %v\n", st.ID, err)
			continue
		}

		if err := s.repo.MarkAsExecuted(st.ID); err != nil {
			log.Printf("Failed to mark scheduled transaction %d as executed: %v\n", st.ID, err)
		}
	}
}