package routes

import (
	"github.com/gorilla/mux"
	"accounting-api-with-go/controllers"
)

func SetupRoutes() *mux.Router {
	router := mux.NewRouter()

	router.HandleFunc("/user-login", controllers.Login).Methods("POST")
	router.HandleFunc("/user-register", controllers.Register).Methods("POST")
	// router.HandleFunc("/users/{id}", controllers.GetUserByID).Methods("GET")
	// router.HandleFunc("/users/{id}", controllers.UpdateUser).Methods("PUT")
	// router.HandleFunc("/users/{id}", controllers.DeleteUser).Methods("DELETE")

	// router.HandleFunc("/transactions", controllers.GetTransactions).Methods("GET")
	// router.HandleFunc("/transactions", controllers.CreateTransaction).Methods("POST")
	// router.HandleFunc("/transactions/{id}", controllers.GetTransactionByID).Methods("GET")
	// router.HandleFunc("/transactions/{id}", controllers.UpdateTransaction).Methods("PUT")
	// router.HandleFunc("/transactions/{id}", controllers.DeleteTransaction).Methods("DELETE")

	return router
}