package repositories

import (
	"accounting-api-with-go/internal/models"
	"database/sql"
)

type MySQLUserRepository struct {
	DB *sql.DB
}

func (r *MySQLUserRepository) Save(user *models.User) error {
	query := "INSERT INTO users (username, email, password_hash, role) VALUES (?, ?, ?, ?)"
	_, err := r.DB.Exec(query, user.Username, user.Email, user.Password, user.Role)
	return err
}

func (r *MySQLUserRepository) FindByID(id int64) (*models.User, error) {
	var user models.User
	query := "SELECT id, username, email, role FROM users WHERE id = ?"
	err := r.DB.QueryRow(query, id).Scan(&user.ID, &user.Username, &user.Email, &user.Role)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *MySQLUserRepository) FindByEmail(email string) (*models.User, error) {
	var user models.User
	query := "SELECT id, username, email, password_hash, role FROM users WHERE email = ?"
	err := r.DB.QueryRow(query, email).Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Role)
	if err != nil {
		return nil, err
	}
	return &user, nil
}