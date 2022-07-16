package middlewares

import (
	"log"
	"net/http"
	"strings"

	"github.com/tapiaw38/irrigation-api/claim"
)

var (
	NO_AUTH_NEDDED = []string{
		"/login",
		"/register",
	}
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

// shoudCheckAuth checks if the request should be authenticated
func shoudCheckToken(route string) bool {
	for _, r := range NO_AUTH_NEDDED {
		if strings.Contains(route, r) {
			return false
		}
	}
	return true
}

// MiddlewareAuth is a middleware that validates the token
func MiddlewareAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {

			if !shoudCheckToken(r.URL.Path) {
				next.ServeHTTP(w, r)
				return
			}

			token := r.Header.Get("Authorization")

			_, err := claim.ValidateJWT(token)

			if err != nil {
				log.Printf("Invalid token. Error: %s\n", err)
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		})
}
