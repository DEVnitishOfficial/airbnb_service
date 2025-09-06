package controllers

import (
	"AuthInGo/dto"
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

	roleIdStr := chi.URLParam(r, "id")

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

	role, err := rc.roleService.GetRoleByIdService(roleId)

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

func (rc *RoleController) GetRoleByNameController(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetRoleByName is called in the controller layer")
	roleName := chi.URLParam(r, "name")
	role, err := rc.roleService.GetRoleByNameService(roleName)
	if err != nil {
		utils.WriteJSONErrorResponse(w, http.StatusInternalServerError, "Failed to fetch role", err)
		return
	}
	if role == nil {
		utils.WriteJSONErrorResponse(w, http.StatusNotFound, "Role not found", fmt.Errorf("role with name %s not found", roleName))
		return
	}
	utils.WriteJSONSuccessResponse(w, http.StatusOK, "Role fetched successfully", role)
	fmt.Println("Role fetched successfully:", role)
}

func (rc *RoleController) GetAllRolesController(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetAllRoles is called in the controller layer")
	roles, err := rc.roleService.GetAllRolesService()
	if err != nil {
		utils.WriteJSONErrorResponse(w, http.StatusInternalServerError, "Failed to fetch roles", err)
		return
	}
	utils.WriteJSONSuccessResponse(w, http.StatusOK, "Roles fetched successfully", roles)
	fmt.Println("Roles fetched successfully:", roles)
}

func (rc *RoleController) CreateRoleController(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CreateRole is called in the controller layer")
	payload := r.Context().Value("payload").(dto.CreateRoleRequestDto)
	fmt.Println("Payload received in the controller layer:", payload)
	role, err := rc.roleService.CreateRoleService(payload.RoleName, payload.Description)
	if err != nil {
		utils.WriteJSONErrorResponse(w, http.StatusInternalServerError, "Failed to create role", err)
		return
	}
	utils.WriteJSONSuccessResponse(w, http.StatusOK, "Role created successfully", role)
	fmt.Println("Role created successfully:", role)
}

func (rc *RoleController) DeleteRoleByIdController(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DeleteRoleById is called in the controller layer")
	roleId := chi.URLParam(r, "id")
	roleIdInt, err := strconv.ParseInt(roleId, 10, 64)
	if err != nil {
		utils.WriteJSONErrorResponse(w, http.StatusBadRequest, "Invalid role ID", err)
		return
	}
	err = rc.roleService.DeleteRoleByIdService(roleIdInt)
	if err != nil {
		utils.WriteJSONErrorResponse(w, http.StatusInternalServerError, "Failed to delete role", err)
		return
	}
	utils.WriteJSONSuccessResponse(w, http.StatusOK, "Role deleted successfully", nil)
	fmt.Println("Role deleted successfully")
}

func (rc *RoleController) UpdateRoleByIdController(w http.ResponseWriter, r *http.Request) {
	fmt.Println("UpdateRoleById is called in the controller layer")
	roleId := chi.URLParam(r, "id")
	roleIdInt, err := strconv.ParseInt(roleId, 10, 64)
	if err != nil {
		utils.WriteJSONErrorResponse(w, http.StatusBadRequest, "Invalid role ID", err)
		return
	}
	payload := r.Context().Value("payload").(dto.UpdateRoleRequestDto)
	role, err := rc.roleService.UpdateRoleByIdService(roleIdInt, payload.RoleName, payload.Description)
	if err != nil {
		utils.WriteJSONErrorResponse(w, http.StatusInternalServerError, "Failed to update role", err)
		return
	}
	utils.WriteJSONSuccessResponse(w, http.StatusOK, "Role updated successfully", role)
	fmt.Println("Role updated successfully:", role)
}
