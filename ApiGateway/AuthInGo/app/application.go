package app

import (
	config "AuthInGo/config/env"
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
		Config: config,
	}
}

func (app *Application) Run() error {
	//returning reference of the created server
	server := &http.Server{
		// below is the configuration of the server
		Addr:         app.Config.Addr,
		Handler:      nil, // TODO : setup a chi router and put it here
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	fmt.Println("Starting server on http://localhost", app.Config.Addr)

	return server.ListenAndServe()
}
