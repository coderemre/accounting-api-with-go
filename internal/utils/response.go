package utils

import (
	"encoding/json"
	"net/http"
)

func WriteErrorResponse(w http.ResponseWriter, message string, statusCode int) {
	response := map[string]interface{}{
		"error":   true,
		"message": message,
	}
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}

func WriteSuccessResponse(w http.ResponseWriter, data interface{}, statusCode int) {
	response := map[string]interface{}{
		"error": false,
		"data":  data,
	}
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}