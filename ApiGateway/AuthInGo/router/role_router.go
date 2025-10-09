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
	r.With(middlewares.JWTAuthMiddleware, middlewares.RequireAnyRole("admin")).Get("/role/getbyid/{id}", rr.RoleController.GetRoleByIdController)
	r.With(middlewares.JWTAuthMiddleware, middlewares.RequireAnyRole("admin")).Get("/role/getbyname/{name}", rr.RoleController.GetRoleByNameController)
	r.With(middlewares.JWTAuthMiddleware, middlewares.RequireAnyRole("admin")).Get("/role/getall", rr.RoleController.GetAllRolesController)
	r.With(middlewares.JWTAuthMiddleware, middlewares.RequireAnyRole("admin"), middlewares.RoleCreateRequestValidator).Post("/role/create", rr.RoleController.CreateRoleController)
	r.With(middlewares.JWTAuthMiddleware, middlewares.RequireAnyRole("admin")).Delete("/role/deletebyid/{id}", rr.RoleController.DeleteRoleByIdController)
	r.With(middlewares.JWTAuthMiddleware, middlewares.RequireAnyRole("admin"), middlewares.RoleUpdateRequestValidator).Put("/role/updatebyid/{id}", rr.RoleController.UpdateRoleByIdController)
}
