package repositories

import (
	"accounting-api-with-go/internal/models"
	"accounting-api-with-go/internal/utils"
	"database/sql"
	"errors"
	"strings"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) CreateUser(user *models.User) (int64, error) {
	query := `
		INSERT INTO users (username, email, password_hash, role)
		VALUES (?, ?, ?, ?)
	`
	result, err := r.DB.Exec(query, user.Username, user.Email, user.Password, user.Role)
	if err != nil {
		if strings.Contains(err.Error(), "Duplicate entry") {
			return 0, errors.New(utils.ErrUserAlreadyExists.String())
		}
		return 0, err
	}
	return result.LastInsertId()
}

func (r *UserRepository) GetUserByEmail(email string) (*models.User, error) {
	var user models.User
	query := "SELECT id, username, email, password_hash, role FROM users WHERE email = ?"

	err := r.DB.QueryRow(query, email).Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Role)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}