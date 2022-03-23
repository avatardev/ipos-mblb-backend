package middleware

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func CorsMiddleware() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
			if r.Method != http.MethodOptions {
				next.ServeHTTP(rw, r)
				return
			}

			log.Printf("[CorsMiddleware] received request %s -> %s\n", r.Host, r.URL)

			rw.Header().Set("Access-Control-Allow-Origin", "*")
			rw.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST, PUT, DELETE, PATCH")
			rw.Header().Set("Access-Control-Allow-Headers", "Content-Type, X-CSRF-Token")

			rw.Write([]byte("OKOK"))
		})
	}
}