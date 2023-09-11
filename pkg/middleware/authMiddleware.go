package middleware

import (
	"example/go-nurse-mgmt/pkg/helper"
	"net/http"
	"strings"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Extract the token from the request

		tokenString := extractTokenFromRequest(r)
		// Parse and validate the token
		_, err := helper.ParseToken(tokenString)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}
		// Call the next handler in the chain
		next.ServeHTTP(w, r)
	})
}

// Extract the token from the request
func extractTokenFromRequest(r *http.Request) string {
	token := r.Header.Get("Authorization")
	splitToken := strings.Split(token, "Bearer ")
	token = splitToken[1]
	return token
}
