package handlers

import (
	"encoding/json"
	"net/http"

	"accounting-api-with-go/internal/middlewares"
	"accounting-api-with-go/internal/models"
	"accounting-api-with-go/internal/services"
	"accounting-api-with-go/internal/utils"
)

type UserHandler struct {
	UserService *services.UserService
}

func NewUserHandler(userService *services.UserService) *UserHandler {
	return &UserHandler{UserService: userService}
}

func (h *UserHandler) Register(w http.ResponseWriter, r *http.Request) {
	var user models.User
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.WriteErrorResponse(w, "Invalid request", http.StatusBadRequest)
		return
	}

	createdUser, token, err := h.UserService.Register(&user)
	if err != nil {
		utils.WriteErrorResponse(w, err.Error(), http.StatusConflict)
		return
	}

	utils.WriteSuccessResponse(w, map[string]interface{}{
		"message": "User registered successfully",
		"user": map[string]interface{}{
			"id":       createdUser.ID,
			"username": createdUser.Username,
			"email":    createdUser.Email,
			"role":     createdUser.Role,
		},
		"token": token,
	}, http.StatusCreated)
}

func (h *UserHandler) Login(w http.ResponseWriter, r *http.Request) {
	var credentials models.User
	w.Header().Set("Content-Type", "application/json")

	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		utils.WriteErrorResponse(w, "Invalid request", http.StatusBadRequest)
		return
	}

	user, token, err := h.UserService.Login(credentials.Email, credentials.Password)
	if err != nil {
		utils.WriteErrorResponse(w, err.Error(), http.StatusUnauthorized)
		return
	}

	utils.WriteSuccessResponse(w, map[string]interface{}{
		"message": "Login successful",
		"user": map[string]interface{}{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
			"role":     user.Role,
		},
		"token": token,
	}, http.StatusOK)
}

func (h *UserHandler) Profile(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value(middlewares.UserContextKey).(*models.User)
	if !ok || user == nil {
		utils.WriteErrorResponse(w, "Unauthorized", http.StatusUnauthorized)
		return
	}

	utils.WriteSuccessResponse(w, map[string]interface{}{
		"message": "User profile",
		"user": map[string]interface{}{
			"id":       user.ID,
			"username": user.Username,
			"email":    user.Email,
			"role":     user.Role,
		},
	}, http.StatusOK)
}