package middlewares

import (
	"net"
	"net/http"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

// Store IP -> limiter
var visitors = make(map[string]*rate.Limiter)
var mu sync.Mutex

// GetLimiter returns limiter for IP, or creates one
func getLimiter(ip string) *rate.Limiter {
	mu.Lock()
	defer mu.Unlock()

	limiter, exists := visitors[ip]
	if !exists {
		// 5 requests per minute per IP
		limiter = rate.NewLimiter(rate.Every(time.Minute), 5)
		visitors[ip] = limiter
	}
	return limiter
}

// Middleware
func RateLimitMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip, _, err := net.SplitHostPort(r.RemoteAddr)
		if err != nil {
			http.Error(w, "Unable to determine IP", http.StatusInternalServerError)
			return
		}

		// to check for differnet ip address uncomment the below code and comment the above code
		/** in postman to simulate different ip address we can use the X-Forwarded-For header
		in post man put X-Forwarded-For : 127.0.0.1

		ip := r.Header.Get("X-Forwarded-For")
		if ip == "" {
			ip, _, _ = net.SplitHostPort(r.RemoteAddr)
		}
		*/

		limiter := getLimiter(ip)
		if !limiter.Allow() {
			http.Error(w, "Too many requests", http.StatusTooManyRequests)
			return
		}
		next.ServeHTTP(w, r)
	})
}
