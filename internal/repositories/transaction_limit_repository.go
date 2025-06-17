package repositories

import (
	"accounting-api-with-go/internal/models"
	"database/sql"
)

type TransactionLimitRepository struct {
	db *sql.DB
}

func NewTransactionLimitRepository(db *sql.DB) *TransactionLimitRepository {
	return &TransactionLimitRepository{db: db}
}

func (r *TransactionLimitRepository) GetLimit(userID int64, txType string) (models.TransactionLimit, error) {
	var limit models.TransactionLimit
	query := `
		SELECT max_amount, max_count, period
		FROM transaction_limits
		WHERE user_id = ? AND type = ?
	`
	err := r.db.QueryRow(query, userID, txType).Scan(
		&limit.MaxAmount, &limit.MaxCount, &limit.Period,
	)
	return limit, err
}