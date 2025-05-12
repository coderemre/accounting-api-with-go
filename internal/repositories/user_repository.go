package repositories

import (
	"accounting-api-with-go/internal/models"
	"context"
	"database/sql"
)

type UserRepository struct {
	DB *sql.DB
}

func NewUserRepository(db *sql.DB) *UserRepository {
	return &UserRepository{DB: db}
}

func (r *UserRepository) CreateUser(ctx context.Context, user *models.User) (int64, error) {
	query := `
		INSERT INTO users (username, email, password_hash, role)
		VALUES (?, ?, ?, ?)
	`
	result, err := r.DB.ExecContext(ctx, query, user.Username, user.Email, user.Password, user.Role)
	if err != nil {
		return 0, err
	}
	return result.LastInsertId()
}

func (r *UserRepository) GetUserByEmail(ctx context.Context, email string) (*models.User, error) {
	var user models.User
	query := `SELECT id, username, email, password_hash, role FROM users WHERE email = ?`
	err := r.DB.QueryRowContext(ctx, query, email).Scan(&user.ID, &user.Username, &user.Email, &user.Password, &user.Role)
	if err != nil {
		return nil, err
	}
	return &user, nil
}

func (r *UserRepository) GetAllUsers(ctx context.Context) ([]models.User, error) {
	rows, err := r.DB.QueryContext(ctx, `SELECT id, username, email, role FROM users`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var u models.User
		err := rows.Scan(&u.ID, &u.Username, &u.Email, &u.Role)
		if err != nil {
			return nil, err
		}
		users = append(users, u)
	}
	return users, nil
}

func (r *UserRepository) GetUserByID(ctx context.Context, userID int64) (*models.User, error) {
	var u models.User
	err := r.DB.QueryRowContext(ctx, `SELECT id, username, email, role FROM users WHERE id = ?`, userID).
		Scan(&u.ID, &u.Username, &u.Email, &u.Role)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (r *UserRepository) UpdateUser(ctx context.Context, userID int64, user *models.User) error {
	_, err := r.DB.ExecContext(ctx, `UPDATE users SET username = ?, email = ?, role = ? WHERE id = ?`,
		user.Username, user.Email, user.Role, userID)
	return err
}

func (r *UserRepository) DeleteUser(ctx context.Context, userID int64) error {
	_, err := r.DB.ExecContext(ctx, `DELETE FROM users WHERE id = ?`, userID)
	return err
}