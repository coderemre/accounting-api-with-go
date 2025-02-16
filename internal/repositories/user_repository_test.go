package repositories

import (
	"accounting-api-with-go/internal/models"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUserRepository(t *testing.T) {
	mockRepo := NewMockUserRepository()
	
	err := mockRepo.Save(&models.User{Username: "testuser", Email: "test@example.com", Password: "hashedpassword"})
	assert.Nil(t, err)

	user, err := mockRepo.FindByID(1)
	assert.Nil(t, err)
	assert.Equal(t, int64(1), user.ID)

	user, err = mockRepo.FindByEmail("test@example.com")
	assert.Nil(t, err)
	assert.Equal(t, "test@example.com", user.Email)
}