package middleware

import (
	"log"
	"net/http"
)

func CorsMiddlware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("Handle CORS")
		h.ServeHTTP(w, r)
	})
}
