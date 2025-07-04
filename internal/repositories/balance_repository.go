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

func (r *BalanceRepository) GetBalance(userID int64, currency string) (*models.Balance, error) {
	var balance models.Balance
	err := r.DB.QueryRow(`
		SELECT user_id, amount, currency FROM balances 
		WHERE user_id = ? AND currency = ?`, userID, currency).
		Scan(&balance.UserID, &balance.Amount, &balance.Currency)

	if err == sql.ErrNoRows {
		return &models.Balance{UserID: userID, Amount: 0, Currency: currency}, nil
	}
	return &balance, err
}

func (r *BalanceRepository) UpdateBalance(userID int64, currency string, newAmount float64) error {
	query := `
		INSERT INTO balances (user_id, currency, amount, last_updated_at)
		VALUES (?, ?, ?, NOW())
		ON DUPLICATE KEY UPDATE amount = VALUES(amount), last_updated_at = NOW();
	`
	_, err := r.DB.Exec(query, userID, currency, newAmount)
	return err
}

func (r *BalanceRepository) GetBalanceHistory(userID int64, currency string) ([]models.BalanceHistory, error) {
	query := `
		SELECT user_id, amount, currency, created_at 
		FROM balance_history 
		WHERE user_id = ? AND currency = ? 
		ORDER BY created_at DESC
	`
	rows, err := r.DB.Query(query, userID, currency)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var history []models.BalanceHistory
	for rows.Next() {
		var record models.BalanceHistory
		if err := rows.Scan(&record.UserID, &record.Amount, &record.Currency, &record.CreatedAt); err != nil {
			return nil, err
		}
		history = append(history, record)
	}

	return history, nil
}

func (r *BalanceRepository) GetBalanceAtTime(userID int64, currency string, atTime time.Time) (float64, error) {
	var balance float64

	query := `
		SELECT amount FROM balance_history
		WHERE user_id = ? AND currency = ? AND created_at <= ?
		ORDER BY created_at DESC LIMIT 1
	`
	err := r.DB.QueryRow(query, userID, currency, atTime).Scan(&balance)

	if err == sql.ErrNoRows {
		return 0, errors.New("no balance record found for the given time")
	} else if err != nil {
		return 0, err
	}

	return balance, nil
}