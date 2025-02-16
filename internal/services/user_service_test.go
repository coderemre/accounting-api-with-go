package services

import (
	"accounting-api-with-go/internal/models"
	"accounting-api-with-go/internal/repositories"
	"testing"

	"github.com/stretchr/testify/assert"
)

type MockUserService struct {
	repo repositories.UserRepository
}

func (s *MockUserService) CreateUser(user *models.User) error {
	return s.repo.Save(user)
}

func (s *MockUserService) GetUserByID(id int64) (*models.User, error) {
	return s.repo.FindByID(id)
}

func (s *MockUserService) AuthenticateUser(email, password string) (*models.User, error) {
	user, err := s.repo.FindByEmail(email)
	if err != nil {
		return nil, err
	}
	if user.Password != password {
		return nil, nil
	}
	return user, nil
}

func TestUserService(t *testing.T) {
	mockRepo := repositories.NewMockUserRepository()
	service := &MockUserService{repo: mockRepo}

	err := service.CreateUser(&models.User{Username: "testuser", Email: "test@example.com", Password: "hashedpassword"})
	assert.Nil(t, err)

	user, err := service.GetUserByID(1)
	assert.Nil(t, err)
	assert.Equal(t, int64(1), user.ID)

	user, err = service.AuthenticateUser("test@example.com", "hashedpassword")
	assert.Nil(t, err)
	assert.NotNil(t, user)
}