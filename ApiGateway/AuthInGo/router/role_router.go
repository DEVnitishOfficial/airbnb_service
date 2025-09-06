package router

import (
	"AuthInGo/controllers"
	"AuthInGo/middlewares"

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
	r.Get("/role/getbyid/{id}", rr.RoleController.GetRoleByIdController)
	r.Get("/role/getbyname/{name}", rr.RoleController.GetRoleByNameController)
	r.Get("/role/getall", rr.RoleController.GetAllRolesController)
	r.With(middlewares.RoleCreateRequestValidator).Post("/role/create", rr.RoleController.CreateRoleController)
	r.Delete("/role/deletebyid/{id}", rr.RoleController.DeleteRoleByIdController)
	r.With(middlewares.RoleUpdateRequestValidator).Put("/role/updatebyid/{id}", rr.RoleController.UpdateRoleByIdController)
}
