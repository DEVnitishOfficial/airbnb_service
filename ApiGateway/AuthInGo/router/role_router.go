package router

import (
	"AuthInGo/controllers"

	"github.com/go-chi/chi/v5"
)

type RoleRouter struct {
	RoleController controllers.RoleController
}

func NewRoleRouter(_roleController controllers.RoleController) Router {
	return &RoleRouter{
		RoleController: _roleController,
	}
}

func (rr *RoleRouter) Register(r chi.Router) {
	r.Get("/role/{id}", rr.RoleController.GetRoleByIdController)
}
