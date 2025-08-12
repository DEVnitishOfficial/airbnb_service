package app

import (
	config "AuthInGo/config/env"
	"AuthInGo/controllers"
	db "AuthInGo/db/repositories"
	"AuthInGo/router"
	service "AuthInGo/services"
	"fmt"
	"net/http"
	"time"
)

type Config struct {
	Addr string // port
}

type Application struct {
	Config Config
}

// constructor for config
func NewConfig(addr string) Config {
	port := config.GetString("PORT", "8080") // Load the PORT from .env or use default
	return Config{
		Addr: port,
	}
}

// constructor for Application
func NewApplication(config Config) *Application {
	return &Application{
		Config: config, // here we pass the config to Application struct
	}
}

func (app *Application) Run() error {

	ur := db.NewUserRepository()
	us := service.NewUserService(ur)
	uc := controllers.NewUserController(us)
	uRouter := router.NewUserRouter(*uc)

	//returning reference of the created server
	server := &http.Server{
		// below is the configuration of the server
		Addr:         app.Config.Addr,
		Handler:      router.SetUpRouter(uRouter),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	fmt.Println("Starting server on http://localhost", app.Config.Addr)

	return server.ListenAndServe()
}
