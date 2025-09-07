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

type PermissionController struct {
	permissionService services.PermissionService
}

func NewPermissionController(_permissionService services.PermissionService) *PermissionController {
	return &PermissionController{
		permissionService: _permissionService,
	}
}

func (pc *PermissionController) GetPermissionByIdController(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Get permission by id called in controller layer")

	permissionIdStr := chi.URLParam(r, "id")

	fmt.Println("permissionId from the query param:", permissionIdStr)

	if permissionIdStr == "" {
		utils.WriteJSONErrorResponse(w, http.StatusBadRequest, "Permission ID is required", fmt.Errorf("permission ID is required"))
		return
	}

	permissionId, err := strconv.ParseInt(permissionIdStr, 10, 64)
	if err != nil {
		utils.WriteJSONErrorResponse(w, http.StatusBadRequest, "Invalid permission ID", err)
		return
	}

	permission, err := pc.permissionService.GetPermissionByIdService(permissionId)

	if err != nil {
		utils.WriteJSONErrorResponse(w, http.StatusInternalServerError, "Failed to fetch role", err)
		return
	}
	if permission == nil {
		utils.WriteJSONErrorResponse(w, http.StatusNotFound, "permission not found", fmt.Errorf("permission with ID %d not found", permissionId))
		return
	}
	utils.WriteJSONSuccessResponse(w, http.StatusOK, "permission fetched successfully", permission)
	fmt.Println("permission fetched successfully:", permission)

}

func (pc *PermissionController) GetPermissionByNameController(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetPermissionByName is called in the controller layer")
	permissionName := chi.URLParam(r, "name")
	permission, err := pc.permissionService.GetPermissionByNameService(permissionName)
	if err != nil {
		utils.WriteJSONErrorResponse(w, http.StatusInternalServerError, "Failed to fetch permission", err)
		return
	}
	if permission == nil {
		utils.WriteJSONErrorResponse(w, http.StatusNotFound, "Permission not found", fmt.Errorf("permission with name %s not found", permissionName))
		return
	}
	utils.WriteJSONSuccessResponse(w, http.StatusOK, "permission fetched successfully", permission)
	fmt.Println("Permission fetched successfully:", permission)
}

func (pc *PermissionController) GetAllPermissionController(w http.ResponseWriter, r *http.Request) {
	fmt.Println("GetAllPermission is called in the controller layer")
	permission, err := pc.permissionService.GetAllPermissionsService()
	if err != nil {
		utils.WriteJSONErrorResponse(w, http.StatusInternalServerError, "Failed to fetch permission", err)
		return
	}
	utils.WriteJSONSuccessResponse(w, http.StatusOK, "Permission fetched successfully", permission)
	fmt.Println("Roles fetched successfully:", permission)
}

func (pc *PermissionController) CreatePermissionController(w http.ResponseWriter, r *http.Request) {
	fmt.Println("CreatePermission is called in the controller layer")
	payload := r.Context().Value("payload").(dto.CreatePermissionRequestDto)
	fmt.Println("Payload received in the controller layer:", payload)
	permission, err := pc.permissionService.CreatePermissionService(payload.PermissionName, payload.Description, payload.Resource, payload.Action)
	if err != nil {
		utils.WriteJSONErrorResponse(w, http.StatusInternalServerError, "Failed to create Permission", err)
		return
	}
	utils.WriteJSONSuccessResponse(w, http.StatusOK, "Permission created successfully", permission)
	fmt.Println("Permission created successfully:", permission)
}

func (pc *PermissionController) DeletePermissionByIdController(w http.ResponseWriter, r *http.Request) {
	fmt.Println("DeletePermissionById is called in the controller layer")
	permissionId := chi.URLParam(r, "id")
	permissionInt, err := strconv.ParseInt(permissionId, 10, 64)
	if err != nil {
		utils.WriteJSONErrorResponse(w, http.StatusBadRequest, "Invalid permission ID", err)
		return
	}
	err = pc.permissionService.DeletePermissionByIdService(permissionInt)
	if err != nil {
		utils.WriteJSONErrorResponse(w, http.StatusInternalServerError, "Failed to delete permission", err)
		return
	}
	utils.WriteJSONSuccessResponse(w, http.StatusOK, "Permission deleted successfully", nil)
	fmt.Println("Permission deleted successfully")
}

func (pc *PermissionController) UpdatePermissionByIdController(w http.ResponseWriter, r *http.Request) {
	fmt.Println("UpdatePermissionById is called in the controller layer")
	permissionId := chi.URLParam(r, "id")
	permissionIdInt, err := strconv.ParseInt(permissionId, 10, 64)
	if err != nil {
		utils.WriteJSONErrorResponse(w, http.StatusBadRequest, "Invalid permission ID", err)
		return
	}
	payload := r.Context().Value("payload").(dto.UpdatePermissionRequestDto)
	permission, err := pc.permissionService.UpdatePermissionByIdService(permissionIdInt, payload.PermissionName, payload.Description, payload.Resource, payload.Action)

	if err != nil {
		utils.WriteJSONErrorResponse(w, http.StatusInternalServerError, "Failed to update permission", err)
		return
	}
	utils.WriteJSONSuccessResponse(w, http.StatusOK, "Permission updated successfully", permission)
	fmt.Println("Permission updated successfully:", permission)
}
