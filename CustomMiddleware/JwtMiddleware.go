package CustomMiddleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/MicahParks/keyfunc/v3"
	"github.com/golang-jwt/jwt/v5"
	"github.com/nmcnew/tournament-winner/models"
)

func JwtMiddleware(next http.Handler, keyFunc keyfunc.Keyfunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		if tokenString == authHeader {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		claims := &models.Claims{}
		_, err := jwt.ParseWithClaims(tokenString, claims, keyFunc.Keyfunc)
		if err != nil {
			http.Error(w, "Failed to Parse jwt", http.StatusBadRequest)
		}
		ctx := context.WithValue(r.Context(), models.ClaimsKey, claims)
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
