package middlewares

import (
	"fmt"
	"net/http"
)

func RequestLogger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("Received request", r.Method, r.URL.Path)
		next.ServeHTTP(w, r) // call the next middleware in the chain
	})
}
