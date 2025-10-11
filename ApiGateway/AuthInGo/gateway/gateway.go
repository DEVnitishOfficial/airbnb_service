package gateway

import (
	"AuthInGo/middlewares"
	"AuthInGo/utils"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func NewGatewayRouter() http.Handler {
	r := chi.NewRouter()

	// Forward requests to HotelService
	r.With(middlewares.JWTAuthMiddleware, middlewares.RequireAnyRole("user", "admin")).HandleFunc("/hotelService/*", utils.ProxyToService(
		"http://localhost:3001", // Target service
		"/hotelService",         // Prefix to strip
	))

	// Forward requests to BookingService
	r.With(middlewares.JWTAuthMiddleware, middlewares.RequireAnyRole("user", "admin")).HandleFunc("/bookingService/*", utils.ProxyToService(
		"http://localhost:3005",
		"/bookingService",
	))

	// Forward requests to ReviewService
	r.With(middlewares.JWTAuthMiddleware, middlewares.RequireAnyRole("user", "admin")).HandleFunc("/reviewService/*", utils.ProxyToService(
		"http://localhost:8081",
		"/reviewService",
	))

	return r
}
