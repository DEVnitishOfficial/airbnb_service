package router

import (
	"AuthInGo/controllers"

	"github.com/go-chi/chi/v5"
)

// here *chi.Mux is the return type of the chi router
func SetUpRouter() *chi.Mux {
	router := chi.NewRouter()

	router.Get("/ping", controllers.PingHandler)

	return router
}
