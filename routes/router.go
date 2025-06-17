package routes

import (
	"context"
	"database/sql"
	"net/http"
	"os"

	"github.com/gorilla/mux"

	"accounting-api-with-go/handlers"
	"accounting-api-with-go/internal/cache"
	"accounting-api-with-go/internal/eventstore"
	"accounting-api-with-go/internal/middlewares"
	"accounting-api-with-go/internal/repositories"
	"accounting-api-with-go/internal/services"

	"go.opentelemetry.io/otel"
)

func SetupRoutes(db *sql.DB, es eventstore.EventStore) *mux.Router {
	tracer := otel.Tracer("router")

	_, span := tracer.Start(context.Background(), "SetupRoutes")
	defer span.End()

	redisAddr := os.Getenv("REDIS_PORT")
	if redisAddr == "" {
		redisAddr = "redis:6379"
	}
	redisCache := cache.NewRedisCache(redisAddr)

	router := mux.NewRouter()

	userRepo := repositories.NewUserRepository(db)
	balanceRepo := repositories.NewBalanceRepository(db)
	transactionRepo := repositories.NewTransactionRepository(db)

	balanceService := services.NewBalanceService(balanceRepo, redisCache)
	userService := services.NewUserService(userRepo, balanceService, es, redisCache)
	transactionService := services.NewTransactionService(transactionRepo, balanceService, redisCache)

	userHandler := handlers.NewUserHandler(userService)
	transactionHandler := handlers.NewTransactionHandler(transactionService)
	balanceHandler := handlers.NewBalanceHandler(balanceService)
	adminHandler := handlers.NewAdminHandler(services.NewReplayService(userRepo, transactionService),es)

	authRoutes := router.PathPrefix("/api/v1/auth").Subrouter()
	authRoutes.HandleFunc("/register", userHandler.Register).Methods("POST")
	authRoutes.HandleFunc("/login", userHandler.Login).Methods("POST")
	authRoutes.Handle("/refresh", middlewares.JWTAuthMiddleware(userRepo)(http.HandlerFunc(userHandler.RefreshToken))).Methods("POST")
	
	userRoutes := router.PathPrefix("/api/v1/users").Subrouter()
	userRoutes.Use(middlewares.JWTAuthMiddleware(userRepo))

	adminRoutes := router.PathPrefix("/api/v1/admin").Subrouter()
	adminRoutes.HandleFunc("/replay", adminHandler.Replay).Methods("POST")

	userRoutes.HandleFunc("", middlewares.RequireAdmin(userHandler.GetAllUsers)).Methods("GET")
	userRoutes.HandleFunc("/{id:[0-9]+}", middlewares.RequireAdmin(userHandler.GetUserByID)).Methods("GET")
	userRoutes.HandleFunc("/{id:[0-9]+}", middlewares.RequireAdmin(userHandler.DeleteUser)).Methods("DELETE")


	userRoutes.HandleFunc("/{id:[0-9]+}", userHandler.UpdateUser).Methods("PUT")

	transactionRoutes := router.PathPrefix("/api/v1/transactions").Subrouter()
	transactionRoutes.Use(middlewares.JWTAuthMiddleware(userRepo))
	transactionRoutes.HandleFunc("/credit", transactionHandler.Credit).Methods("POST")
	transactionRoutes.HandleFunc("/debit", transactionHandler.Debit).Methods("POST")
	transactionRoutes.HandleFunc("/transfer", transactionHandler.Transfer).Methods("POST")
	transactionRoutes.HandleFunc("/history", transactionHandler.GetTransactionHistory).Methods("GET")
	transactionRoutes.HandleFunc("/{id:[0-9]+}", transactionHandler.GetTransactionByID).Methods("GET")


	balanceRoutes := router.PathPrefix("/api/v1/balances").Subrouter()
	balanceRoutes.Use(middlewares.JWTAuthMiddleware(userRepo))
	balanceRoutes.HandleFunc("/current", balanceHandler.GetCurrentBalance).Methods("GET")
	balanceRoutes.HandleFunc("/historical", balanceHandler.GetHistoricalBalances).Methods("GET")
	balanceRoutes.HandleFunc("/at-time", balanceHandler.GetBalanceAtTime).Methods("GET")

	return router
}