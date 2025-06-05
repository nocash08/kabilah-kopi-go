package middleware

import (
	"backend/config"
	"backend/helper"
	"context"
	"net/http"
	"strings"
)

func AuthMiddleware() func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// Get token from cookie
			cookie, err := r.Cookie(helper.AccessTokenCookie)
			if err != nil {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			tokenString := cookie.Value
			if tokenString == "" {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			// Validate the token using secret from config
			claims, err := helper.ValidateToken(tokenString, config.AppConfig.JWTSecret)
			if err != nil {
				if strings.Contains(err.Error(), "expired") {
					http.Error(w, "Token expired", http.StatusUnauthorized)
					return
				}
				http.Error(w, "Invalid token", http.StatusUnauthorized)
				return
			}

			// Add claims to request context
			ctx := r.Context()
			ctx = context.WithValue(ctx, "user_id", claims.UserID)
			ctx = context.WithValue(ctx, "username", claims.Username)
			ctx = context.WithValue(ctx, "is_admin", claims.IsAdmin)

			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}
