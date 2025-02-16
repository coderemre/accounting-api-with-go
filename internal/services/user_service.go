package services

import "accounting-api-with-go/internal/models"

type UserService interface {
	CreateUser(user *models.User) error
	GetUserByID(id int64) (*models.User, error)
	AuthenticateUser(email, password string) (*models.User, error)
}