package handlers

import (
	"fmt"
	"net/http"
	"time"

	"accounting-api-with-go/internal/middlewares"
	"accounting-api-with-go/internal/models"
	"accounting-api-with-go/internal/services"
	"accounting-api-with-go/internal/utils"
)

type BalanceHandler struct {
	BalanceService *services.BalanceService
}

func NewBalanceHandler(balanceService *services.BalanceService) *BalanceHandler {
	return &BalanceHandler{BalanceService: balanceService}
}

func (h *BalanceHandler) GetCurrentBalance(w http.ResponseWriter, r *http.Request) {
	fmt.Printf("DEBUG CONTEXT USER: %#v\n", r.Context().Value("user"))
	user, ok := r.Context().Value(middlewares.UserContextKey).(*models.User)
	if !ok || user == nil {
		utils.WriteErrorResponse(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	userID := user.ID

	balance, err := h.BalanceService.GetCurrentBalance(userID)
	if err != nil {
		utils.WriteErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.WriteSuccessResponse(w, map[string]float64{"current_balance": balance}, http.StatusOK)
}

func (h *BalanceHandler) GetHistoricalBalances(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value(middlewares.UserContextKey).(*models.User)
	if !ok || user == nil {
		utils.WriteErrorResponse(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	userID := user.ID

	history, err := h.BalanceService.GetBalanceHistory(userID)
	if err != nil {
		utils.WriteErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.WriteSuccessResponse(w, history, http.StatusOK)
}

func (h *BalanceHandler) GetBalanceAtTime(w http.ResponseWriter, r *http.Request) {
	user, ok := r.Context().Value(middlewares.UserContextKey).(*models.User)
	if !ok || user == nil {
		utils.WriteErrorResponse(w, "Unauthorized", http.StatusUnauthorized)
		return
	}
	userID := user.ID

	timestamp := r.URL.Query().Get("timestamp")
	if timestamp == "" {
		utils.WriteErrorResponse(w, "Timestamp is required", http.StatusBadRequest)
		return
	}

	parsedTime, err := time.Parse(time.RFC3339, timestamp)
	if err != nil {
		utils.WriteErrorResponse(w, "Invalid timestamp format, expected RFC3339", http.StatusBadRequest)
		return
	}

	balance, err := h.BalanceService.GetBalanceAtTime(userID, parsedTime)
	if err != nil {
		utils.WriteErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.WriteSuccessResponse(w, map[string]float64{"balance_at_time": balance}, http.StatusOK)
}