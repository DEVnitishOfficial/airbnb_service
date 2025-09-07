package router

import (
	"AuthInGo/controllers"
	"AuthInGo/middlewares"

	"github.com/go-chi/chi/v5"
)

type PermissionRouter struct {
	PermissionController controllers.PermissionController
}

func NewPermissionRouter(_permissionController controllers.PermissionController) Router {
	return &PermissionRouter{
		PermissionController: _permissionController,
	}
}

func (pr *PermissionRouter) Register(r chi.Router) {
	r.Get("/permission/getbyid/{id}", pr.PermissionController.GetPermissionByIdController)
	r.Get("/permission/getbyname/{name}", pr.PermissionController.GetPermissionByNameController)
	r.Get("/permission/getall", pr.PermissionController.GetAllPermissionController)
	r.With(middlewares.PermissionCreateRequestValidator).Post("/permission/create", pr.PermissionController.CreatePermissionController)
	r.Delete("/permission/deletebyid/{id}", pr.PermissionController.DeletePermissionByIdController)
	r.With(middlewares.PermissionUpdateRequestValidator).Put("/permission/updatebyid/{id}", pr.PermissionController.UpdatePermissionByIdController)
}
