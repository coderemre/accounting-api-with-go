package routes

import (
	"accounting-api-with-go/handlers"
	"accounting-api-with-go/internal/middlewares"
	"accounting-api-with-go/internal/repositories"
	"accounting-api-with-go/internal/services"
	"database/sql"

	"github.com/gorilla/mux"
)

func SetupRoutes(db *sql.DB) *mux.Router {
	router := mux.NewRouter()

	// Repositories
	userRepo := repositories.NewUserRepository(db)
	transactionRepo := repositories.NewTransactionRepository(db)
	balanceRepo := repositories.NewBalanceRepository(db)

	// Services
	userService := services.NewUserService(userRepo)
	balanceService := services.NewBalanceService(balanceRepo)
	transactionService := services.NewTransactionService(transactionRepo, balanceService)

	// Handlers
	userHandler := handlers.NewUserHandler(userService)
	transactionHandler := handlers.NewTransactionHandler(transactionService)
	balanceHandler := handlers.NewBalanceHandler(balanceService)

	// Public routes (JWT gerekmeyen)
	router.HandleFunc("/user-login", userHandler.Login).Methods("POST")
	router.HandleFunc("/user-register", userHandler.Register).Methods("POST")

	// Protected routes (JWT doğrulama gerektiren)
	protectedRoutes := router.PathPrefix("/api/v1").Subrouter()
	protectedRoutes.Use(middlewares.JWTAuthMiddleware(userRepo))

	// Transaction Endpoints
	protectedRoutes.HandleFunc("/transactions/credit", transactionHandler.Credit).Methods("POST")
	protectedRoutes.HandleFunc("/transactions/debit", transactionHandler.Debit).Methods("POST")
	protectedRoutes.HandleFunc("/transactions/transfer", transactionHandler.Transfer).Methods("POST")
	protectedRoutes.HandleFunc("/transactions/history", transactionHandler.GetTransactionHistory).Methods("GET")
	protectedRoutes.HandleFunc("/transactions/{id:[0-9]+}", transactionHandler.GetTransactionByID).Methods("GET")

	// Balance Endpoints
	protectedRoutes.HandleFunc("/balances/current", balanceHandler.GetCurrentBalance).Methods("GET")
	protectedRoutes.HandleFunc("/balances/historical", balanceHandler.GetHistoricalBalances).Methods("GET")
	protectedRoutes.HandleFunc("/balances/at-time", balanceHandler.GetBalanceAtTime).Methods("GET")

	return router
}