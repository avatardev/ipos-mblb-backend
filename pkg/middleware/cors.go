package middleware

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func CorsMiddleware() mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {

			log.Printf("[CorsMiddleware] received request %s (%s) -> %s %s\n", r.Header.Get("X-Forwarded-For"), r.Header.Get("X-Real-IP"), r.Method, r.URL)

			rw.Header().Set("Access-Control-Allow-Origin", "*")
			rw.Header().Set("Access-Control-Allow-Methods", "OPTIONS, GET, POST, PUT, DELETE, PATCH")
			rw.Header().Set("Access-Control-Allow-Headers", "*")

			if r.Method != http.MethodOptions {
				next.ServeHTTP(rw, r)
				return
			}

			rw.Write([]byte("okok"))
		})
	}
}
