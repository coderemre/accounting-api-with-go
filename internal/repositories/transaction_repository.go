package repositories

import (
	"accounting-api-with-go/internal/models"
	"database/sql"
	"errors"
)

type TransactionRepository struct {
	DB *sql.DB
}

func NewTransactionRepository(db *sql.DB) *TransactionRepository {
	return &TransactionRepository{DB: db}
}

func (r *TransactionRepository) CreateTransaction(tx *models.Transaction) error {
	query := `INSERT INTO transactions (user_id, amount, type, created_at) VALUES (?, ?, ?, NOW())`
	_, err := r.DB.Exec(query, tx.UserID, tx.Amount, tx.Type)
	return err
}

func (r *TransactionRepository) CreateTransfer(senderID, receiverID int64, amount float64) error {
	query := `INSERT INTO transactions (user_id, amount, type, created_at) VALUES (?, ?, ?, NOW()), (?, ?, ?, NOW())`
	_, err := r.DB.Exec(query, senderID, -amount, "debit", receiverID, amount, "credit")
	return err
}

func (r *TransactionRepository) GetTransactionByID(transactionID int64) (*models.Transaction, error) {
	var transaction models.Transaction

	query := `SELECT id, user_id, amount, type, created_at FROM transactions WHERE id = ?`
	err := r.DB.QueryRow(query, transactionID).Scan(
		&transaction.ID, &transaction.UserID, &transaction.Amount, &transaction.Type, &transaction.CreatedAt,
	)

	if err == sql.ErrNoRows {
		return nil, errors.New("transaction not found")
	} else if err != nil {
		return nil, err
	}

	return &transaction, nil
}

func (r *TransactionRepository) GetTransactionHistory(userID int64) ([]models.Transaction, error) {
	query := `SELECT id, user_id, amount, type, created_at FROM transactions WHERE user_id = ? ORDER BY created_at DESC`

	rows, err := r.DB.Query(query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []models.Transaction

	for rows.Next() {
		var transaction models.Transaction
		err := rows.Scan(&transaction.ID, &transaction.UserID, &transaction.Amount, &transaction.Type, &transaction.CreatedAt)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}

	return transactions, nil
}