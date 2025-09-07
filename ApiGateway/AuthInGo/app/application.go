package app

import (
	dbConfig "AuthInGo/config/db"
	config "AuthInGo/config/env"
	"AuthInGo/controllers"
	repo "AuthInGo/db/repositories"
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

	db, err := dbConfig.SetupDB()

	if err != nil {
		fmt.Println("Error setting up database:", err)
		return err
	}

	ur := repo.NewUserRepository(db)
	rr := repo.NewRoleRepository(db)
	pr := repo.NewPermissionRepository(db)

	us := service.NewUserService(ur)
	rs := service.NewRoleService(rr)
	ps := service.NewPermissionService(pr)

	uc := controllers.NewUserController(us)
	rc := controllers.NewRoleController(rs)
	pc := controllers.NewPermissionController(ps)

	uRouter := router.NewUserRouter(*uc)
	rRouter := router.NewRoleRouter(*rc)
	pRouter := router.NewPermissionRouter(*pc)

	//returning reference of the created server
	server := &http.Server{
		// below is the configuration of the server
		Addr:         app.Config.Addr,
		Handler:      router.SetUpRouter(uRouter, rRouter, pRouter),
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}
	fmt.Println("Starting server on http://localhost", app.Config.Addr)

	return server.ListenAndServe()
}
