package router

import (
	"AuthInGo/controllers"
	"AuthInGo/middlewares"

	"github.com/go-chi/chi/v5"
)

type RolePermissionRouter struct {
	RolePermissionController controllers.RolePermissionController
}

func NewRolePermissionRouter(_rolePermissionController controllers.RolePermissionController) Router {
	return &RolePermissionRouter{
		RolePermissionController: _rolePermissionController,
	}
}

func (rpr *RolePermissionRouter) Register(r chi.Router) {
	r.Get("/rolepermission/getbyid/{id}", rpr.RolePermissionController.GetRolePermissionByIdController)
	r.Get("/rolepermission/getbyroleid/{id}", rpr.RolePermissionController.GetRolePermissionByRoleIdController)
	r.Get("/rolepermission/getall", rpr.RolePermissionController.GetAllRolePermissionsController)
	r.With(middlewares.AssignPermissionRequestValidator).Post("/rolepermission/addbyroleid/{id}", rpr.RolePermissionController.AddPermissionToRoleController)
	r.With(middlewares.RemovePermissionRequestValidator).Post("/rolepermission/removebyroleid/{id}", rpr.RolePermissionController.RemovePermissionFromRoleController)
}
