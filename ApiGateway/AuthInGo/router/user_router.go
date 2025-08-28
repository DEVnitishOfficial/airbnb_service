package router

import (
	"AuthInGo/controllers"

	"github.com/go-chi/chi/v5"
)

type UserRouter struct {
	UserController controllers.UserController
}

func NewUserRouter(_userController controllers.UserController) Router {
	return &UserRouter{
		UserController: _userController,
	}
}

func (ur *UserRouter) Register(r chi.Router) {
	r.Get("/alluser", ur.UserController.GetAllUserController)
	r.Get("/profile", ur.UserController.GetUserById)
	r.Post("/signup", ur.UserController.CreateUser)
	r.Post("/signin", ur.UserController.LoginUser)
	r.Get("/getbyid/:id", ur.UserController.GetUserById)
}
