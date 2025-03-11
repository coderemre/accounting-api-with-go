package services

import (
	"accounting-api-with-go/internal/auth"
	"errors"

	"accounting-api-with-go/internal/models"
	"accounting-api-with-go/internal/repositories"
	"accounting-api-with-go/internal/utils"

	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	UserRepo *repositories.UserRepository
}

func NewUserService(userRepo *repositories.UserRepository) *UserService {
	return &UserService{UserRepo: userRepo}
}

func (s *UserService) Register(user *models.User) (*models.User, string, error) {
	if validationError := user.Validate(false); validationError != "" {
		return nil, "", errors.New(string(validationError))
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, "", errors.New(utils.ErrPasswordHashFailed.String())
	}
	user.Password = string(hashedPassword)


	userID, err := s.UserRepo.CreateUser(user)
	if err != nil {
		return nil, "", err
	}

	createdUser, err := s.UserRepo.GetUserByEmail(user.Email)
	if err != nil {
		return nil, "", err
	}

	createdUser.ID = userID

	token, err := auth.GenerateJWT(*createdUser)
	if err != nil {
		return nil, "", errors.New(utils.ErrTokenGeneration.String())
	}

	return createdUser, token, nil
}

func (s *UserService) Login(email, password string) (*models.User, string, error) {
	user := &models.User{
		Email:    email,
		Password: password,
	}

	if validationError := user.Validate(true); validationError != "" {
		return nil, "", errors.New(string(validationError))
	}

	existingUser, err := s.UserRepo.GetUserByEmail(email)
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