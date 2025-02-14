package routes

import (
	"accounting-api-with-go/handlers"

	"github.com/gorilla/mux"
)

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/user-login", handlers.Login).Methods("POST")
	router.HandleFunc("/user-register", handlers.Register).Methods("POST")

	return router
}