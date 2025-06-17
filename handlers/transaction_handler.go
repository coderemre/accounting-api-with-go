package handlers

import (
	"encoding/json"
	"net/http"
	"strconv"

	"accounting-api-with-go/internal/models"
	"accounting-api-with-go/internal/services"
	"accounting-api-with-go/internal/utils"

	"github.com/gorilla/mux"
)

type TransactionHandler struct {
	TransactionService *services.TransactionService
}

func NewTransactionHandler(transactionService *services.TransactionService) *TransactionHandler {
	return &TransactionHandler{TransactionService: transactionService}
}

func (h *TransactionHandler) Credit(w http.ResponseWriter, r *http.Request) {
	var request struct {
		ToUserID int64   `json:"to_user_id"`
		Amount   float64 `json:"amount"`
		Currency   string  `json:"currency"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		utils.WriteErrorResponse(w, "Invalid request", http.StatusBadRequest)
		return
	}

	transaction, err := h.TransactionService.ProcessTransaction(0, request.ToUserID, request.Amount, "credit", request.Currency)
	if err != nil {
		utils.WriteErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	utils.WriteSuccessResponse(w, transaction, http.StatusOK)
}

func (h *TransactionHandler) Debit(w http.ResponseWriter, r *http.Request) {
	var request struct {
		FromUserID int64   `json:"from_user_id"`
		Amount     float64 `json:"amount"`
		Currency   string  `json:"currency"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		utils.WriteErrorResponse(w, "Invalid request", http.StatusBadRequest)
		return
	}

	transaction, err := h.TransactionService.ProcessTransaction(request.FromUserID, 0, request.Amount, "debit", request.Currency)
	if err != nil {
		utils.WriteErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	utils.WriteSuccessResponse(w, transaction, http.StatusOK)
}

func (h *TransactionHandler) Transfer(w http.ResponseWriter, r *http.Request) {
	var request struct {
		SenderID   int64   `json:"sender_id"`
		ReceiverID int64   `json:"receiver_id"`
		Amount     float64 `json:"amount"`
		Currency   string  `json:"currency"`
	}

	if err := json.NewDecoder(r.Body).Decode(&request); err != nil {
		utils.WriteErrorResponse(w, "Invalid request", http.StatusBadRequest)
		return
	}

	err := h.TransactionService.Transfer(request.SenderID, request.ReceiverID, request.Amount, request.Currency)
	if err != nil {
		utils.WriteErrorResponse(w, err.Error(), http.StatusBadRequest)
		return
	}

	utils.WriteSuccessResponse(w, map[string]string{"message": "Transfer successful"}, http.StatusOK)
}

func (h *TransactionHandler) GetTransactionHistory(w http.ResponseWriter, r *http.Request) {
	userID, err := strconv.ParseInt(r.URL.Query().Get("user_id"), 10, 64)
	if err != nil {
		utils.WriteErrorResponse(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	transactions, err := h.TransactionService.GetTransactionHistory(userID)
	if err != nil {
		utils.WriteErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.WriteSuccessResponse(w, transactions, http.StatusOK)
}

func (h *TransactionHandler) GetTransactionByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	transactionID, err := strconv.ParseInt(vars["id"], 10, 64)
	if err != nil {
		utils.WriteErrorResponse(w, "Invalid transaction ID", http.StatusBadRequest)
		return
	}

	transaction, err := h.TransactionService.GetTransactionByID(transactionID)
	if err != nil {
		utils.WriteErrorResponse(w, err.Error(), http.StatusNotFound)
		return
	}

	utils.WriteSuccessResponse(w, transaction, http.StatusOK)
}

func (h *TransactionHandler) ProcessBatchTransactions(w http.ResponseWriter, r *http.Request) {
	var batch []models.Transaction

	if err := json.NewDecoder(r.Body).Decode(&batch); err != nil {
		utils.WriteErrorResponse(w, "Invalid batch payload", http.StatusBadRequest)
		return
	}

	transactions, err := h.TransactionService.ProcessBatchTransactions(r.Context(), batch)
	if err != nil {
		utils.WriteErrorResponse(w, err.Error(), http.StatusInternalServerError)
		return
	}

	utils.WriteSuccessResponse(w, transactions, http.StatusOK)
}