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

type RolePermissionController struct {
	rolePermissionService services.RolePermissionService
}

func NewRolePermissionController(_rolePermissionService services.RolePermissionService) *RolePermissionController {
	return &RolePermissionController{
		rolePermissionService: _rolePermissionService,
	}
}

func (rpc *RolePermissionController) GetRolePermissionByIdController(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetRolePermissionById is called in the controller layer")

	rolePermissionId := chi.URLParam(r, "id")

	fmt.Println("rolePermissionId from URL param:", rolePermissionId)

	if rolePermissionId == "" {
		fmt.Println("Role Permission ID is required")
		utils.WriteJSONErrorResponse(w, http.StatusBadRequest, "Role Permission ID is required", fmt.Errorf("role permission ID is required"))
	}

	rolePermissionIdInt, err := strconv.ParseInt(rolePermissionId, 10, 64)
	if err != nil {
		fmt.Println("Invalid role permission ID:", err)
		utils.WriteJSONErrorResponse(w, http.StatusBadRequest, "Invalid role permission ID", err)
		return
	}

	rolePermission, err := rpc.rolePermissionService.GetRolePermissionByIdService(rolePermissionIdInt)
	if err != nil {
		fmt.Println("Failed to fetch role permission:", err)
		utils.WriteJSONErrorResponse(w, http.StatusInternalServerError, "Failed to fetch role permission", err)
		return
	}
	if rolePermission == nil {
		fmt.Println("Role permission not found")
		utils.WriteJSONErrorResponse(w, http.StatusNotFound, "Role permission not found", fmt.Errorf("role permission with ID %d not found", rolePermissionIdInt))
		return
	}
	fmt.Println("Role permission fetched successfully:", rolePermission)
	utils.WriteJSONSuccessResponse(w, http.StatusOK, "Role permission fetched successfully", rolePermission)
}

func (rpc *RolePermissionController) GetRolePermissionByRoleIdController(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetRolePermissionByRoleId called in controller layer")

	roleIdStr := chi.URLParam(r, "id")

	fmt.Println("rolePermissionId from URL param:", roleIdStr)

	if roleIdStr == "" {
		fmt.Println("Role Permission ID is required")
		utils.WriteJSONErrorResponse(w, http.StatusBadRequest, "Role Permission ID is required", fmt.Errorf("role permission ID is required"))
	}

	roleIdInt, err := strconv.ParseInt(roleIdStr, 10, 64)
	if err != nil {
		fmt.Println("Invalid role permission ID:", err)
		utils.WriteJSONErrorResponse(w, http.StatusBadRequest, "Invalid role permission ID", err)
		return
	}

	rolePermissions, err := rpc.rolePermissionService.GetRolePermissionByRoleIdService(roleIdInt)
	if err != nil {
		fmt.Println("Failed to fetch role permissions by role ID:", err)
		utils.WriteJSONErrorResponse(w, http.StatusInternalServerError, "Failed to fetch role permissions by role ID", err)
		return
	}
	if rolePermissions == nil {
		fmt.Println("No role permissions found for the given role ID")
		utils.WriteJSONErrorResponse(w, http.StatusNotFound, "No role permissions found for the given role ID", fmt.Errorf("no role permissions found for role ID %d", roleIdInt))
		return
	}
	utils.WriteJSONSuccessResponse(w, http.StatusOK, "Role permissions fetched successfully", rolePermissions)
	fmt.Println("Role permissions fetched successfully:", rolePermissions)
}

func (rpc *RolePermissionController) AddPermissionToRoleController(w http.ResponseWriter, r *http.Request) {
	fmt.Println("AddPermissionToRole called in controller layer")

	roleIdStr := chi.URLParam(r, "id")
	if roleIdStr == "" {
		utils.WriteJSONErrorResponse(w, http.StatusBadRequest, "Role ID is required", fmt.Errorf("missing role ID"))
		return
	}

	roleIdInt, err := strconv.ParseInt(roleIdStr, 10, 64)
	if err != nil {
		utils.WriteJSONErrorResponse(w, http.StatusBadRequest, "Invalid role ID", err)
		return
	}

	payload := r.Context().Value("payload").(dto.AssignPermissionRequestDTO)

	rolePermission, err := rpc.rolePermissionService.AddPermissionToRoleService(roleIdInt, payload.PermissionId)
	if err != nil {
		fmt.Println("Failed to add permission to role:", err)
		utils.WriteJSONErrorResponse(w, http.StatusInternalServerError, "Failed to add permission to role", err)
		return
	}
	if rolePermission == nil {
		fmt.Println("Failed to add permission to role: returned nil")
		utils.WriteJSONErrorResponse(w, http.StatusInternalServerError, "Failed to add permission to role", fmt.Errorf("failed to add permission to role"))
		return
	}
	utils.WriteJSONSuccessResponse(w, http.StatusOK, "Permission added to role successfully", rolePermission)
	fmt.Println("Permission added to role successfully:", rolePermission)
}

func (rpc *RolePermissionController) RemovePermissionFromRoleController(w http.ResponseWriter, r *http.Request) {
	fmt.Println("RemovePermissionFromRole called from controller layer")

	roleIdStr := chi.URLParam(r, "id")
	if roleIdStr == "" {
		utils.WriteJSONErrorResponse(w, http.StatusBadRequest, "Role ID is required", fmt.Errorf("missing role ID"))
		return
	}

	roleIdInt, err := strconv.ParseInt(roleIdStr, 10, 64)
	if err != nil {
		utils.WriteJSONErrorResponse(w, http.StatusBadRequest, "Invalid role ID", err)
		return
	}

	payload := r.Context().Value("payload").(dto.RemovePermissionRequestDTO)

	err = rpc.rolePermissionService.RemovePermissionFromRoleService(roleIdInt, payload.PermissionId)

	if err != nil {
		fmt.Println("Failed to remove permission from role:", err)
		utils.WriteJSONErrorResponse(w, http.StatusInternalServerError, "Failed to remove permission from role", err)
		return
	}
	utils.WriteJSONSuccessResponse(w, http.StatusOK, "Permission removed from role successfully", nil)
	fmt.Println("Permission removed from role successfully")
}

func (rpc *RolePermissionController) GetAllRolePermissionsController(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetAllRolePermissions called inside the controller layer")

	rolePermissions, err := rpc.rolePermissionService.GetAllRolePermissionsService()

	if err != nil {
		fmt.Println("Failed to fetch all role permissions:", err)
		utils.WriteJSONErrorResponse(w, http.StatusInternalServerError, "Failed to fetch all role permissions", err)
		return
	}
	utils.WriteJSONSuccessResponse(w, http.StatusOK, "Role permissions fetched successfully", rolePermissions)
	fmt.Println("Role permissions fetched successfully:", rolePermissions)
}
