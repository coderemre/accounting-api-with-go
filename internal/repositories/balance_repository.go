package repositories

import (
	"accounting-api-with-go/internal/models"
	"database/sql"
	"errors"
	"time"
)

type BalanceRepository struct {
	DB *sql.DB
}

func NewBalanceRepository(db *sql.DB) *BalanceRepository {
	return &BalanceRepository{DB: db}
}

func (r *BalanceRepository) GetBalance(userID int64) (*models.Balance, error) {
	var balance models.Balance
	err := r.DB.QueryRow(`SELECT user_id, amount FROM balances WHERE user_id = ?`, userID).Scan(&balance.UserID, &balance.Amount)
	if err == sql.ErrNoRows {
		return &models.Balance{UserID: userID, Amount: 0}, nil
	}
	return &balance, err
}

func (r *BalanceRepository) UpdateBalance(userID int64, newAmount float64) error {
	_, err := r.DB.Exec(`UPDATE balances SET amount = ? WHERE user_id = ?`, newAmount, userID)
	return err
}

func (r *BalanceRepository) GetBalanceHistory(userID int64) ([]models.BalanceHistory, error) {
	query := `SELECT user_id, amount, created_at FROM balance_history WHERE user_id = ? ORDER BY created_at DESC`
	rows, err := r.DB.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var history []models.BalanceHistory
	for rows.Next() {
		var record models.BalanceHistory
		if err := rows.Scan(&record.UserID, &record.Amount, &record.CreatedAt); err != nil {
			return nil, err
		}
		history = append(history, record)
	}

	return history, nil
}

func (r *BalanceRepository) GetBalanceAtTime(userID int64, atTime time.Time) (float64, error) {
	var balance float64

	query := `
		SELECT amount FROM balance_history
		WHERE user_id = ? AND created_at <= ?
		ORDER BY created_at DESC LIMIT 1
	`
	err := r.DB.QueryRow(query, userID, atTime).Scan(&balance)

	if err == sql.ErrNoRows {
		return 0, errors.New("no balance record found for the given time")
	} else if err != nil {
		return 0, err
	}

	return balance, nil
}