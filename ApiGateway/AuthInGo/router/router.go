package router

import (
	"AuthInGo/controllers"
	"AuthInGo/middlewares"
	"AuthInGo/utils"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

type Router interface {
	Register(r chi.Router)
}

// here *chi.Mux is the return type of the chi router
func SetUpRouter(UserRouter Router) *chi.Mux {
	chiRouter := chi.NewRouter()

	chiRouter.Use(middleware.Logger)
	// chiRouter.Use(middlewares.UserLoginRequestValidator)

	chiRouter.Use(middlewares.RateLimitMiddleware)
	chiRouter.Use(middlewares.RequestLogger)
	chiRouter.HandleFunc("/fakestoreService/*", utils.ProxyToService("http://fakestoreapi.in/", "/fakestoreService"))
	chiRouter.Get("/ping", controllers.PingHandler)

	UserRouter.Register(chiRouter)

	return chiRouter
}

// http://localhost:3000/fakestoreService/products/categories
