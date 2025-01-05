package controllers

import (
	"encoding/json"
	"net/http"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"golang.org/x/crypto/bcrypt"

	"accounting-api-with-go/internal/database"
	"accounting-api-with-go/internal/models"
	"accounting-api-with-go/internal/utils"
)

var jwtSecret = []byte("your_secret_key")

func Register(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var user models.User
	if err := json.NewDecoder(r.Body).Decode(&user); err != nil {
		utils.Log.Error().Err(err).Msg(utils.ErrInvalidRequest.String())
		writeErrorResponse(w, utils.ErrInvalidRequest, http.StatusBadRequest)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.PasswordHash), bcrypt.DefaultCost)
	if err != nil {
		utils.Log.Error().Err(err).Msg(utils.ErrPasswordHashFailed.String())
		writeErrorResponse(w, utils.ErrPasswordHashFailed, http.StatusInternalServerError)
		return
	}

	query := `
		INSERT INTO users (username, email, password_hash, role, created_at, updated_at)
		VALUES (?, ?, ?, ?, ?, ?)
	`
	now := time.Now()
	_, err = database.DB.Exec(query, user.Username, user.Email, hashedPassword, user.Role, now, now)
	if err != nil {
		utils.Log.Error().Err(err).Msg(utils.ErrDuplicateEntry.String())
		if strings.Contains(err.Error(), "users.username") {
			writeErrorResponse(w, utils.ErrUsernameExists, http.StatusConflict)
		} else if strings.Contains(err.Error(), "users.email") {
			writeErrorResponse(w, utils.ErrUserAlreadyExists, http.StatusConflict)
		} else {
			writeErrorResponse(w, utils.ErrDuplicateEntry, http.StatusConflict)
		}
		return
	}

	token, err := generateJWT(user)
	if err != nil {
		utils.Log.Error().Err(err).Msg(utils.ErrTokenGeneration.String())
		writeErrorResponse(w, utils.ErrTokenGeneration, http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"error":   false,
		"message": utils.SuccessUserRegistered.String(),
		"token":   token,
	}
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		utils.Log.Error().Err(err).Msg(utils.ErrResponseEncodingFailed.String())
		writeErrorResponse(w, utils.ErrResponseEncodingFailed, http.StatusInternalServerError)
	}
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	authHeader := r.Header.Get("Authorization")
	if authHeader != "" {
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		claims, err := verifyJWT(tokenString)
		if err == nil {
			var user models.User
			query := "SELECT id, username, email, role, created_at, updated_at FROM users WHERE id = ?"
			err := database.DB.QueryRow(query, claims["user_id"]).Scan(
				&user.ID, &user.Username, &user.Email, &user.Role, &user.CreatedAt, &user.UpdatedAt,
			)
			if err != nil {
				utils.Log.Error().Err(err).Msg(utils.ErrUserRetrievalFailed.String())
				writeErrorResponse(w, utils.ErrUserRetrievalFailed, http.StatusInternalServerError)
				return
			}

			newToken, err := generateJWT(user)
			if err != nil {
				utils.Log.Error().Err(err).Msg(utils.ErrTokenGeneration.String())
				writeErrorResponse(w, utils.ErrTokenGeneration, http.StatusInternalServerError)
				return
			}

			response := map[string]interface{}{
				"error":   false,
				"message": utils.SuccessTokenValidated.String(),
				"user":    user,
				"token":   newToken,
			}
			w.WriteHeader(http.StatusOK)
			if err := json.NewEncoder(w).Encode(response); err != nil {
				utils.Log.Error().Err(err).Msg(utils.ErrResponseEncodingFailed.String())
				writeErrorResponse(w, utils.ErrResponseEncodingFailed, http.StatusInternalServerError)
			}
			return
		}
	}

	var credentials struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	if err := json.NewDecoder(r.Body).Decode(&credentials); err != nil {
		utils.Log.Error().Err(err).Msg(utils.ErrInvalidRequest.String())
		writeErrorResponse(w, utils.ErrInvalidRequest, http.StatusBadRequest)
		return
	}

	var user models.User
	query := "SELECT id, username, email, password_hash, role, created_at, updated_at FROM users WHERE email = ?"
	err := database.DB.QueryRow(query, credentials.Email).Scan(
		&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.Role, &user.CreatedAt, &user.UpdatedAt,
	)
	if err != nil {
		utils.Log.Error().Err(err).Msg(utils.ErrInvalidCredentials.String())
		writeErrorResponse(w, utils.ErrInvalidCredentials, http.StatusUnauthorized)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.PasswordHash), []byte(credentials.Password)); err != nil {
		utils.Log.Error().Err(err).Msg(utils.ErrInvalidCredentials.String())
		writeErrorResponse(w, utils.ErrInvalidCredentials, http.StatusUnauthorized)
		return
	}

	token, err := generateJWT(user)
	if err != nil {
		utils.Log.Error().Err(err).Msg(utils.ErrTokenGeneration.String())
		writeErrorResponse(w, utils.ErrTokenGeneration, http.StatusInternalServerError)
		return
	}

	response := map[string]interface{}{
		"error":   false,
		"message": utils.SuccessLogin.String(),
		"user":    user,
		"token":   token,
	}
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(response); err != nil {
		utils.Log.Error().Err(err).Msg(utils.ErrResponseEncodingFailed.String())
		writeErrorResponse(w, utils.ErrResponseEncodingFailed, http.StatusInternalServerError)
	}
}

func writeErrorResponse(w http.ResponseWriter, err utils.Message, statusCode int) {
	response := map[string]interface{}{
		"error":   true,
		"message": err.String(),
	}
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(response)
}

func generateJWT(user models.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"email":   user.Email,
		"role":    user.Role,
		"exp":     time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(jwtSecret)
}

func verifyJWT(tokenString string) (jwt.MapClaims, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, http.ErrAbortHandler
		}
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, http.ErrAbortHandler
}