package repositories

import (
	"accounting-api-with-go/internal/models"
	"errors"
)

type MockUserRepository struct {
	users map[int64]*models.User
}

func NewMockUserRepository() *MockUserRepository {
	return &MockUserRepository{
		users: make(map[int64]*models.User),
	}
}

func (m *MockUserRepository) Save(user *models.User) error {
	if user.ID == 0 {
		user.ID = int64(len(m.users) + 1)
	}
	m.users[user.ID] = user
	return nil
}

func (m *MockUserRepository) FindByID(id int64) (*models.User, error) {
	user, exists := m.users[id]
	if !exists {
		return nil, errors.New("user not found")
	}
	return user, nil
}

func (m *MockUserRepository) FindByEmail(email string) (*models.User, error) {
	for _, user := range m.users {
		if user.Email == email {
			return user, nil
		}
	}
	return nil, errors.New("user not found")
}