package middlewares

import (
	"context"
	"net/http"
	"strings"

	"accounting-api-with-go/internal/auth"
	"accounting-api-with-go/internal/models"
	"accounting-api-with-go/internal/repositories"
	"accounting-api-with-go/internal/utils"
)

type ContextKey string

const UserContextKey ContextKey = "user"

func JWTAuthMiddleware(userRepo *repositories.UserRepository) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				utils.WriteErrorResponse(w, "Missing Authorization header", http.StatusUnauthorized)
				return
			}

			tokenString := strings.TrimPrefix(authHeader, "Bearer ")
			claims, err := auth.ValidateJWT(tokenString)
			if err != nil {
				utils.WriteErrorResponse(w, "Invalid or expired token", http.StatusUnauthorized)
				return
			}

			user, err := userRepo.GetUserByEmail(r.Context(), claims.Email)
			if err != nil || user == nil {
				utils.WriteErrorResponse(w, "User not found", http.StatusUnauthorized)
				return
			}

			ctx := context.WithValue(r.Context(), UserContextKey, user)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func RequireAdmin(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		user, ok := r.Context().Value(UserContextKey).(*models.User)
		if !ok || user == nil || user.Role != "admin" {
			utils.WriteErrorResponse(w, "Forbidden: Admins only", http.StatusForbidden)
			return
		}
		next.ServeHTTP(w, r)
	}
}