package controllers

import (
	services "AuthInGo/services"
	"AuthInGo/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type RoleController struct {
	roleService services.RoleService
}

func NewRoleController(_roleService services.RoleService) *RoleController {
	return &RoleController{
		roleService: _roleService,
	}
}

func (rc *RoleController) GetRoleByIdController(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetRoleById is called in the controller layer")

	// roleIdStr := r.URL.Query().Get("id")
	roleIdStr := chi.URLParam(r, "id")

	// if roleId == "" {
	// 	roleId = r.Context().Value("roleId").(string)
	// }

	fmt.Println("userId from the context or query param:", roleIdStr)

	if roleIdStr == "" {
		utils.WriteJSONErrorResponse(w, http.StatusBadRequest, "Role ID is required", fmt.Errorf("role ID is required"))
		return
	}

	roleId, err := strconv.ParseInt(roleIdStr, 10, 64)
	if err != nil {
		utils.WriteJSONErrorResponse(w, http.StatusBadRequest, "Invalid role ID", err)
		return
	}

	role, err := rc.roleService.GetRoleById(roleId)

	if err != nil {
		utils.WriteJSONErrorResponse(w, http.StatusInternalServerError, "Failed to fetch role", err)
		return
	}
	if role == nil {
		utils.WriteJSONErrorResponse(w, http.StatusNotFound, "Role not found", fmt.Errorf("role with ID %d not found", roleId))
		return
	}
	utils.WriteJSONSuccessResponse(w, http.StatusOK, "Role fetched successfully", role)
	fmt.Println("Role fetched successfully:", role)

}
