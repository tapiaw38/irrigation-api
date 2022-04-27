package middlewares

import (
	"log"
	"net/http"

	"github.com/tapiaw38/irrigation-api/claim"
)

// MiddlewareLog is a middleware that logs the request
func MiddlewareLog(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			log.Printf("New request. Method: %s. IP: %s. requested URL: %s\n",
				r.Method, r.RemoteAddr, r.URL)
			next.ServeHTTP(w, r)
		})
}

// MiddlewareAuth is a middleware that validates the token
func MiddlewareAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			token := r.Header.Get("Authorization")

			_, err := claim.ValidateJWT(token)

			if err != nil {
				log.Printf("Invalid token. Error: %s\n", err)
				w.WriteHeader(http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		})
}
