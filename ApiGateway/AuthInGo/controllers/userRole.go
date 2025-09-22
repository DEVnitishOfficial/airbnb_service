package controllers

import (
	services "AuthInGo/services"
	"AuthInGo/utils"
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type UserRoleController struct {
	UserRoleService services.UserRolesService
}

func NewUserRoleController(_userRoleService services.UserRolesService) *UserRoleController {
	return &UserRoleController{
		UserRoleService: _userRoleService,
	}
}

func (urc *UserRoleController) GetUserRolesController(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Fetching user roles from controller layer")

	userRoleId := chi.URLParam(r, "id")

	fmt.Println("userRoleId from URL param:", userRoleId)

	if userRoleId == "" {
		fmt.Println("User Role ID is required")
		http.Error(w, "User Role ID is required", http.StatusBadRequest)
		return
	}

	userRoleIdInt, err := strconv.ParseInt(userRoleId, 10, 64)
	if err != nil {
		fmt.Println("Invalid user role ID:", err)
		http.Error(w, "Invalid user role ID", http.StatusBadRequest)
		return
	}

	userRoles, err := urc.UserRoleService.GetUserRoleService(userRoleIdInt)
	if err != nil {
		fmt.Println("Got error while fetching user roles from controller layer", err)
		http.Error(w, "Failed to fetch user roles", http.StatusInternalServerError)
		return
	}

	if userRoles == nil {
		fmt.Println("User roles not found")
		http.Error(w, "User roles not found", http.StatusNotFound)
		return
	}

	fmt.Println("User roles fetched successfully:", userRoles)
	utils.WriteJSONSuccessResponse(w, http.StatusOK, "User roles fetched successfully", userRoles)
}

func (urc *UserRoleController) GetAllUserAndTheirRolesController(w http.ResponseWriter, r *http.Request) {
	fmt.Println("Fetching all users and their roles from controller layer")

	userRoles, err := urc.UserRoleService.GetAllUserAndTheirRolesService()
	if err != nil {
		fmt.Println("Got error while fetching all users and their roles from controller layer", err)
		http.Error(w, "Failed to fetch all users and their roles", http.StatusInternalServerError)
		return
	}

	if userRoles == nil {
		fmt.Println("No users found")
		http.Error(w, "No users found", http.StatusNotFound)
		return
	}

	fmt.Println("All users and their roles fetched successfully:", userRoles)
	utils.WriteJSONSuccessResponse(w, http.StatusOK, "All users and their roles fetched successfully", userRoles)
}

func (urc *UserRoleController) AssignRoleToUserController(w http.ResponseWriter, r *http.Request) {
	fmt.Println("AssignRoleToUserController called in controller layer")
	userId := chi.URLParam(r, "userId")
	roleId := chi.URLParam(r, "roleId")
	if userId == "" {
		utils.WriteJSONErrorResponse(w, http.StatusBadRequest, "User ID is required", fmt.Errorf("missing user ID"))
		return
	}
	if roleId == "" {
		utils.WriteJSONErrorResponse(w, http.StatusBadRequest, "Role ID is required", fmt.Errorf("missing role ID"))
		return
	}

	roleIdInt, err := strconv.ParseInt(roleId, 10, 64)
	if err != nil {
		utils.WriteJSONErrorResponse(w, http.StatusBadRequest, "Invalid role ID", err)
		return
	}

	userIdInt, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		utils.WriteJSONErrorResponse(w, http.StatusBadRequest, "Invalid user ID", err)
		return
	}

	assignedRole, err := urc.UserRoleService.AssignRoleToUserService(userIdInt, roleIdInt)
	if err != nil {
		fmt.Println("Got error while assigning role to user from controller layer", err)
		http.Error(w, "Failed to assign role to user", http.StatusInternalServerError)
		return
	}

	if assignedRole == nil {
		fmt.Println("Role assignment failed")
		http.Error(w, "Role assignment failed", http.StatusInternalServerError)
		return
	}

	fmt.Println("Role assigned to user successfully")
	utils.WriteJSONSuccessResponse(w, http.StatusOK, "Role assigned to user successfully", assignedRole)
}

func (urc *UserRoleController) RemoveRoleFromUserController(w http.ResponseWriter, r *http.Request) {
	fmt.Println("RemoveRoleFromUserController called in controller layer")
	userId := chi.URLParam(r, "userId")
	roleId := chi.URLParam(r, "roleId")
	if userId == "" {
		utils.WriteJSONErrorResponse(w, http.StatusBadRequest, "User ID is required", fmt.Errorf("missing user ID"))
		return
	}
	if roleId == "" {
		utils.WriteJSONErrorResponse(w, http.StatusBadRequest, "Role ID is required", fmt.Errorf("missing role ID"))
		return
	}

	roleIdInt, err := strconv.ParseInt(roleId, 10, 64)
	if err != nil {
		utils.WriteJSONErrorResponse(w, http.StatusBadRequest, "Invalid role ID", err)
		return
	}

	userIdInt, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		utils.WriteJSONErrorResponse(w, http.StatusBadRequest, "Invalid user ID", err)
		return
	}

	err = urc.UserRoleService.RemoveRoleFromUserService(userIdInt, roleIdInt)
	if err != nil {
		fmt.Println("Got error while removing role from user from controller layer", err)
		http.Error(w, "Failed to remove role from user", http.StatusInternalServerError)
		return
	}

	fmt.Println("Role removed from user successfully")
	utils.WriteJSONSuccessResponse(w, http.StatusOK, "Role removed from user successfully", nil)
}

func (urc *UserRoleController) HasRoleController(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HasRoleController called in controller layer")
	userId := chi.URLParam(r, "userId")

	fmt.Println("userId from URL param:", userId)

	if userId == "" {
		fmt.Println("User ID is required")
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	userIdInt, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		fmt.Println("Invalid user ID:", err)
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	roleName := chi.URLParam(r, "roleName")

	fmt.Println("roleName from URL param:", roleName)

	if roleName == "" {
		fmt.Println("Role name is required")
		http.Error(w, "Role name is required", http.StatusBadRequest)
		return
	}

	hasRole, err := urc.UserRoleService.HasRoleService(userIdInt, roleName)
	if err != nil {
		fmt.Println("Got error while checking user role from controller layer", err)
		http.Error(w, "Failed to check user role", http.StatusInternalServerError)
		return
	}

	if !hasRole {
		fmt.Println("User does not have role")
		http.Error(w, "User does not have role", http.StatusNotFound)
		return
	}

	fmt.Println("User has role")
	utils.WriteJSONSuccessResponse(w, http.StatusOK, "User has role of ----> "+roleName, nil)
}

func (urc *UserRoleController) HasAllRolesController(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HasAllRolesController called in controller layer")
	userId := chi.URLParam(r, "userId")

	fmt.Println("userId from URL param:", userId)

	if userId == "" {
		fmt.Println("User ID is required")
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	userIdInt, err := strconv.ParseInt(userId, 10, 64)
	if err != nil {
		fmt.Println("Invalid user ID:", err)
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	roleNames := r.URL.Query()["roleName"]

	fmt.Println("roleNames from query param:", roleNames)

	if len(roleNames) == 0 {
		fmt.Println("Role names are required")
		http.Error(w, "Role names are required", http.StatusBadRequest)
		return
	}

	hasAllRoles, err := urc.UserRoleService.HasAllRolesService(userIdInt, roleNames)
	if err != nil {
		fmt.Println("Got error while checking user roles from controller layer", err)
		http.Error(w, "Failed to check user roles", http.StatusInternalServerError)
		return
	}

	if !hasAllRoles {
		fmt.Println("User does not have all roles")
		http.Error(w, "User does not have all roles", http.StatusNotFound)
		return
	}

	fmt.Println("User has all roles")
	utils.WriteJSONSuccessResponse(w, http.StatusOK, "User has all roles", nil)
}

func (urc *UserRoleController) HasAnyRoleController(w http.ResponseWriter, r *http.Request) {
	fmt.Println("HasAnyRoleController called in controller layer")
	userIdStr := chi.URLParam(r, "userId")

	fmt.Println("userId from URL param:", userIdStr)

	if userIdStr == "" {
		fmt.Println("User ID is required")
		http.Error(w, "User ID is required", http.StatusBadRequest)
		return
	}

	userIdInt, err := strconv.ParseInt(userIdStr, 10, 64)
	if err != nil {
		fmt.Println("Invalid user ID:", err)
		http.Error(w, "Invalid user ID", http.StatusBadRequest)
		return
	}

	roleNames := r.URL.Query()["roleName"]

	fmt.Println("roleNames from query param:", roleNames)

	if len(roleNames) == 0 {
		fmt.Println("Role names are required")
		http.Error(w, "Role names are required", http.StatusBadRequest)
		return
	}

	hasAnyRole, err := urc.UserRoleService.HasAnyRoleService(userIdInt, roleNames)
	if err != nil {
		fmt.Println("Got error while checking user roles from controller layer", err)
		http.Error(w, "Failed to check user roles", http.StatusInternalServerError)
		return
	}

	if !hasAnyRole {
		fmt.Println("User does not have any role")
		http.Error(w, "User does not have any role", http.StatusNotFound)
		return
	}

	fmt.Println("User has at least one role")
	utils.WriteJSONSuccessResponse(w, http.StatusOK, "User has at least one role", nil)
}
