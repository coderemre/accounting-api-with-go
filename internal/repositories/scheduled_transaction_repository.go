package repositories

import (
	"database/sql"
	"time"
)

type ScheduledTransaction struct {
	ID            int64
	UserID        int64
	Type          string
	Currency       string
	Amount        float64
	ScheduledTime time.Time
	Executed      bool
}

type ScheduledTransactionRepository struct {
	db *sql.DB
}

func NewScheduledTransactionRepository(db *sql.DB) *ScheduledTransactionRepository {
	return &ScheduledTransactionRepository{db: db}
}

func (r *ScheduledTransactionRepository) Create(st ScheduledTransaction) error {
	_, err := r.db.Exec(`
		INSERT INTO scheduled_transactions (user_id, type, currency, amount, scheduled_time)
		VALUES (?, ?, ?, ?, ?)
	`, st.UserID, st.Type, st.Currency, st.Amount, st.ScheduledTime)
	return err
}

func (r *ScheduledTransactionRepository) GetPending(now time.Time) ([]ScheduledTransaction, error) {
	rows, err := r.db.Query(`
		SELECT id, user_id, type, currency, amount, scheduled_time, executed
		FROM scheduled_transactions
		WHERE executed = false AND scheduled_time <= ?
	`, now)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var result []ScheduledTransaction
	for rows.Next() {
		var st ScheduledTransaction
		if err := rows.Scan(&st.ID, &st.UserID, &st.Type, &st.Currency, &st.Amount, &st.ScheduledTime, &st.Executed); err != nil {
			return nil, err
		}
		result = append(result, st)
	}
	return result, nil
}

func (r *ScheduledTransactionRepository) MarkAsExecuted(id int64) error {
	_, err := r.db.Exec(`UPDATE scheduled_transactions SET executed = true WHERE id = ?`, id)
	return err
}