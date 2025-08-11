package app

import (
	config "AuthInGo/config/env"
	db "AuthInGo/db/repositories"
	"AuthInGo/router"
	"fmt"
	"net/http"
	"time"
)

type Config struct {
	Addr string // port
}

type Application struct {
	Config Config
	Store  db.Storage // In store we will have the access of all repositories
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
		Store:  *db.NewStorage(),
		// when we create a new Application, then  we pass it all the repositories which is inside the Storage struct
	}
}

func (app *Application) Run() error {
	//returning reference of the created server
	server := &http.Server{
		// below is the configuration of the server
		Addr:         app.Config.Addr,
		Handler:      router.SetUpRouter(),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	fmt.Println("Starting server on http://localhost", app.Config.Addr)

	return server.ListenAndServe()
}
