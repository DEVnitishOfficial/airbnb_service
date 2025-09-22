package router

import (
	controllers "AuthInGo/controllers"
	"AuthInGo/middlewares"

	"github.com/go-chi/chi/v5"
)

type userRoleRouter struct {
	UserRoleController controllers.UserRoleController
}

func NewUserRoleRouter(_userRoleController controllers.UserRoleController) Router {
	return &userRoleRouter{
		UserRoleController: _userRoleController,
	}
}

func (urr *userRoleRouter) Register(r chi.Router) {
	r.With(middlewares.JWTAuthMiddleware, middlewares.RequireAnyRole("admin", "user")).Get("/userrole/getbyuserid/{id}", urr.UserRoleController.GetUserRolesController)
	r.With(middlewares.JWTAuthMiddleware, middlewares.RequireAllRoles("admin")).Get("/userrole/getallusersandtheirroles", urr.UserRoleController.GetAllUserAndTheirRolesController)
	r.With(middlewares.JWTAuthMiddleware, middlewares.RequireAllRoles("admin")).Post("/userrole/{userId}/assignrole/{roleId}", urr.UserRoleController.AssignRoleToUserController)
	r.With(middlewares.JWTAuthMiddleware, middlewares.RequireAllRoles("admin")).Post("/userrole/{userId}/removerole/{roleId}", urr.UserRoleController.RemoveRoleFromUserController)
	r.With(middlewares.JWTAuthMiddleware, middlewares.RequireAnyRole("admin", "user")).Get("/userrole/{userId}/hasrole/{roleName}", urr.UserRoleController.HasRoleController)
	r.With(middlewares.JWTAuthMiddleware, middlewares.RequireAnyRole("admin", "user")).Get("/userrole/{userId}/hasallroles", urr.UserRoleController.HasAllRolesController)
	r.With(middlewares.JWTAuthMiddleware, middlewares.RequireAnyRole("admin", "user")).Get("/userrole/{userId}/hasanyrole", urr.UserRoleController.HasAnyRoleController)

}
