package services

type BalanceService interface {
	GetBalance(userID int64) (float64, error)
	UpdateBalance(userID int64, amount float64) error
}