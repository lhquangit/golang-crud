package middleware

import (
	"context"
	"go-crud/auth"
	"go-crud/internal/handler"
	"net/http"
)

type contextKey string

const (
    userContextKey contextKey = "user"
    roleContextKey contextKey = "role"
)

func JWTAuthMiddleware(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request){
        tokenString := r.Header.Get("Authorization")
        if tokenString == "" {
            handler.ResponseWithJson(w, http.StatusUnauthorized, map[string]string{"message": "Missing token"})
            return
        }

        claims, err := auth.ValidateJWT(tokenString)
        if err != nil {
            handler.ResponseWithJson(w, http.StatusUnauthorized, map[string]string{"message": "Invalid token"})
            return
        }

        ctx := context.WithValue(r.Context(), userContextKey, claims.Username)
        ctx = context.WithValue(ctx, roleContextKey, claims.Role)
        next.ServeHTTP(w, r.WithContext(ctx))
    })
}

func GetUsernameFromContext(ctx context.Context) (string, bool){
    username, ok := ctx.Value(roleContextKey).(string)
    return username, ok
}

func GetRoleFromContext(ctx context.Context) (string, bool) {
    role, ok := ctx.Value(roleContextKey).(string)
    return role, ok
}

// Midleware to check if the user is admin
func AdminOnly(next http.Handler) http.Handler {
    return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        role, ok := GetRoleFromContext(r.Context())
        if !ok || role != "ADMIN" {
            handler.ResponseWithJson(w, http.StatusForbidden, map[string]string{"message": "Access forbidden"})
            return
        }
        next.ServeHTTP(w, r)
    })
}
