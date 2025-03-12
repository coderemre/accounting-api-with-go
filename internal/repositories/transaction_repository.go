package repositories

import (
	"accounting-api-with-go/internal/models"
	"database/sql"
	"time"
)

type TransactionRepository struct {
	DB *sql.DB
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{DB: db}
}

func (r *TransactionRepository) CreateTransaction(tx *models.Transaction) error {
	query := `
		INSERT INTO transactions (from_user_id, to_user_id, amount, type, status, created_at)
		VALUES (?, ?, ?, ?, ?, ?)
	`

	_, err := r.DB.Exec(
		query,
		tx.FromUserID,
		tx.ToUserID,
		tx.Amount,
		tx.Type,
		tx.Status,
		tx.CreatedAt,
	)

	return err
}

func (r *TransactionRepository) CreateTransfer(senderID, receiverID int64, amount float64) error {
	query := `
		INSERT INTO transactions (from_user_id, to_user_id, amount, type, status, created_at)
		VALUES (?, ?, ?, ?, ?, ?)
	`

	_, err := r.DB.Exec(
		query,
		senderID,
		receiverID,
		amount,
		"transfer",
		"completed",
		time.Now(),
	)

	return err
}

func (r *TransactionRepository) GetTransactionByID(transactionID int64) (*models.Transaction, error) {
	query := `
		SELECT id, from_user_id, to_user_id, amount, type, status, created_at
		FROM transactions
		WHERE id = ?
	`

	var tx models.Transaction
	err := r.DB.QueryRow(query).Scan(
		&tx.ID,
		&tx.FromUserID,
		&tx.ToUserID,
		&tx.Amount,
		&tx.Type,
		&tx.Status,
		&tx.CreatedAt,
	)

	if err != nil {
		return nil, err
	}

	return &tx, nil
}

func (r *TransactionRepository) GetTransactionHistory(userID int64) ([]models.Transaction, error) {
	query := `
		SELECT id, from_user_id, to_user_id, amount, type, status, created_at
		FROM transactions
		WHERE from_user_id = ? OR to_user_id = ?
		ORDER BY created_at DESC
	`

	rows, err := r.DB.Query(query, userID, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []models.Transaction
	for rows.Next() {
		var tx models.Transaction
		err := rows.Scan(
			&tx.ID,
			&tx.FromUserID,
			&tx.ToUserID,
			&tx.Amount,
			&tx.Type,
			&tx.Status,
			&tx.CreatedAt,
		)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, tx)
	}

	return transactions, nil
}