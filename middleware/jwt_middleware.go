package middleware

import (
	"backend/helper"
	"context"
	"net/http"
	"strings"
)

// Define custom types for context keys to ensure type safety
type contextKey string

const (
	UserIDKey   contextKey = "user_id"
	UsernameKey contextKey = "username"
	IsAdminKey  contextKey = "is_admin"
)

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Get the Authorization header
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "Authorization header is required", http.StatusUnauthorized)
			return
		}

		// Check if the header has the Bearer prefix
		if !strings.HasPrefix(authHeader, "Bearer ") {
			http.Error(w, "Invalid token format. Use 'Bearer <token>'", http.StatusUnauthorized)
			return
		}

		// Extract the token
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == "" {
			http.Error(w, "Token is required", http.StatusUnauthorized)
			return
		}

		// Check if token has been invalidated (logged out)
		if helper.IsTokenInvalidated(tokenString) {
			http.Error(w, "Token has been invalidated", http.StatusUnauthorized)
			return
		}

		// Get the JWT secret
		jwtSecret := helper.GetJWTSecret()

		// Validate the token and get claims
		claims, err := helper.ValidateToken(tokenString, jwtSecret)
		if err != nil {
			http.Error(w, "Invalid or expired token", http.StatusUnauthorized)
			return
		}

		// Add claims to request context
		ctx := r.Context()
		ctx = context.WithValue(ctx, UserIDKey, claims.UserID)
		ctx = context.WithValue(ctx, UsernameKey, claims.Username)
		ctx = context.WithValue(ctx, IsAdminKey, claims.IsAdmin)

		// If all checks pass, call the next handler
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
