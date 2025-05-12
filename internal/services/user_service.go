package services

import (
	"accounting-api-with-go/internal/auth"
	"accounting-api-with-go/internal/models"
	"accounting-api-with-go/internal/repositories"
	"context"
	"errors"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	UserRepo      *repositories.UserRepository
	BalanceService *BalanceService
}

func NewUserService(userRepo *repositories.UserRepository, balanceService *BalanceService) *UserService {
	return &UserService{
		UserRepo:      userRepo,
		BalanceService: balanceService,
	}
}

func (s *UserService) Register(ctx context.Context, user *models.User) (*models.User, string, error) {
	if validationError := user.Validate(false); validationError != "" {
		return nil, "", errors.New(validationError.String())
	}

	existingUser, _ := s.UserRepo.GetUserByEmail(ctx, user.Email)
	if existingUser != nil {
		return nil, "", errors.New("email already exists")
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, "", errors.New("failed to hash password")
	}
	user.Password = string(hashedPassword)

	userID, err := s.UserRepo.CreateUser(ctx, user)
	if err != nil {
		return nil, "", err
	}

	createdUser, err := s.UserRepo.GetUserByEmail(ctx, user.Email)
	if err != nil {
		return nil, "", err
	}

	_ = s.BalanceService.UpdateBalance(userID, 0)

	token, err := auth.GenerateJWT(*createdUser)
	if err != nil {
		return nil, "", errors.New("failed to generate token")
	}

	return createdUser, token, nil
}

func (s *UserService) Login(ctx context.Context, email, password string) (*models.User, string, error) {
	user := &models.User{Email: email, Password: password}
	if validationError := user.Validate(true); validationError != "" {
		return nil, "", errors.New(validationError.String())
	}

	existingUser, err := s.UserRepo.GetUserByEmail(ctx, email)
	if err != nil || existingUser == nil {
		return nil, "", errors.New("invalid email or password")
	}

	if err := bcrypt.CompareHashAndPassword([]byte(existingUser.Password), []byte(password)); err != nil {
		return nil, "", errors.New("invalid email or password")
	}

	token, err := auth.GenerateJWT(*existingUser)
	if err != nil {
		return nil, "", errors.New("failed to generate token")
	}

	return existingUser, token, nil
}

func (s *UserService) GetAllUsers(ctx context.Context) ([]models.User, error) {
	return s.UserRepo.GetAllUsers(ctx)
}

func (s *UserService) GetUserByID(ctx context.Context, userID int64) (*models.User, error) {
	return s.UserRepo.GetUserByID(ctx, userID)
}

func (s *UserService) UpdateUser(ctx context.Context, userID int64, user *models.User) error {
	return s.UserRepo.UpdateUser(ctx, userID, user)
}

func (s *UserService) DeleteUser(ctx context.Context, userID int64) error {
	return s.UserRepo.DeleteUser(ctx, userID)
}