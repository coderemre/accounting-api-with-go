package models

import "database/sql"

type TransactionLimit struct {
	UserID    int64
	Type      string
	MaxAmount sql.NullFloat64
	MaxCount  sql.NullInt64
	Period    string
}