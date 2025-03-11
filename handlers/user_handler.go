package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"

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

func (h *UserHandler) GetAllUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.UserService.GetAllUsers()
	if err != nil {
		utils.WriteErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}
	utils.WriteSuccessResponse(w, users, http.StatusOK)
}

func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	idParam := mux.Vars(r)["id"]
	userID, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		utils.WriteErrorResponse(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	user, err := h.UserService.GetUserByID(userID)
	if err != nil {
		utils.WriteErrorResponse(w, err.Error(), http.StatusNotFound)
		return
	}

	utils.WriteSuccessResponse(w, user, http.StatusOK)
}

func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	idParam := mux.Vars(r)["id"]
	userID, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		utils.WriteErrorResponse(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	var updated models.User
	if err := json.NewDecoder(r.Body).Decode(&updated); err != nil {
		utils.WriteErrorResponse(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	err = h.UserService.UpdateUser(userID, &updated)
	if err != nil {
		utils.WriteErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.WriteSuccessResponse(w, map[string]string{"message": "User updated successfully"}, http.StatusOK)
}

func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	idParam := mux.Vars(r)["id"]
	userID, err := strconv.ParseInt(idParam, 10, 64)
	if err != nil {
		utils.WriteErrorResponse(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	err = h.UserService.DeleteUser(userID)
	if err != nil {
		utils.WriteErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.WriteSuccessResponse(w, map[string]string{"message": "User deleted successfully"}, http.StatusOK)
}

func (h *UserHandler) RefreshToken(w http.ResponseWriter, r *http.Request) {
	utils.WriteSuccessResponse(w, map[string]string{
		"message": "Token refreshed successfully (not implemented)",
	}, http.StatusOK)
}