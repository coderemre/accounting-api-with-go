package services

import (
	"accounting-api-with-go/internal/eventstore"
	"accounting-api-with-go/internal/models"
	"accounting-api-with-go/internal/repositories"
	"context"
	"encoding/json"
)

type ReplayService struct {
    userRepo     *repositories.UserRepository
    transactionSvc *TransactionService
}

func NewReplayService(userRepo *repositories.UserRepository, ts *TransactionService) *ReplayService {
    return &ReplayService{userRepo: userRepo, transactionSvc: ts}
}

func (s *ReplayService) ApplyEvent(ctx context.Context, e eventstore.Event) error {
    switch e.Type {
    case "UserCreated":
        var u models.User
        json.Unmarshal(e.Payload, &u)
        return s.userRepo.CreateFromEvent(ctx, &u)
    }
    return nil
}