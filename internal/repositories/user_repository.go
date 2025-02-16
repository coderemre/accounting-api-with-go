package repositories

import "accounting-api-with-go/internal/models"

type UserRepository interface {
	Save(user *models.User) error
	FindByID(id int64) (*models.User, error)
	FindByEmail(email string) (*models.User, error)
}