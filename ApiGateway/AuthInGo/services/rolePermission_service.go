package service

import (
	db "AuthInGo/db/repositories"
	"AuthInGo/models"
	"fmt"
)

type RolePermissionService interface {
	GetRolePermissionByIdService(id int64) (*models.RolePermission, error)
	GetRolePermissionByRoleIdService(roleId int64) ([]*models.RolePermission, error)
	AddPermissionToRoleService(roleId int64, permissionId int64) (*models.RolePermission, error)
	RemovePermissionFromRoleService(roleId int64, permissionId int64) error
	GetAllRolePermissionsService() ([]*models.RolePermission, error)
}

type RolePermissionServiceImpl struct {
	rolePermissionRepository db.RolePermissionRepository
}

func NewRolePermissionService(_rolePermissionRepository db.RolePermissionRepository) RolePermissionService {
	return &RolePermissionServiceImpl{
		rolePermissionRepository: _rolePermissionRepository,
	}
}

func (rp *RolePermissionServiceImpl) GetRolePermissionByIdService(id int64) (*models.RolePermission, error) {
	fmt.Println("Fetching role permission by ID from service layer")

	rolePermission, err := rp.rolePermissionRepository.GetRolePermissionById(id)
	if err != nil {
		fmt.Println("Got error while fetching role permission by ID from service layer", err)
		return nil, err
	}
	return rolePermission, nil
}

func (rp *RolePermissionServiceImpl) GetRolePermissionByRoleIdService(roleId int64) ([]*models.RolePermission, error) {
	fmt.Println("Fetching role permissions by Role ID from service layer")
	rolePermissions, err := rp.rolePermissionRepository.GetRolePermissionByRoleId(roleId)
	if err != nil {
		fmt.Println("Got error while fetching role permissions by Role ID from service layer", err)
		return nil, err
	}
	return rolePermissions, nil
}

func (rp *RolePermissionServiceImpl) AddPermissionToRoleService(roleId int64, permissionId int64) (*models.RolePermission, error) {
	fmt.Println("Adding permission to role from service layer")
	rolePermission, err := rp.rolePermissionRepository.AddPermissionToRole(roleId, permissionId)
	if err != nil {
		fmt.Println("Got error while adding permission to role from service layer", err)
		return nil, err
	}
	return rolePermission, nil
}

func (rp *RolePermissionServiceImpl) RemovePermissionFromRoleService(roleId int64, permissionId int64) error {
	fmt.Println("Removing permission from role from service layer")
	err := rp.rolePermissionRepository.RemovePermissionFromRole(roleId, permissionId)
	if err != nil {
		fmt.Println("Got error while removing permission from role from service layer", err)
		return err
	}
	return nil
}

func (rp *RolePermissionServiceImpl) GetAllRolePermissionsService() ([]*models.RolePermission, error) {
	fmt.Println("Fetching all role permissions from service layer")
	rolePermissions, err := rp.rolePermissionRepository.GetAllRolePermissions()
	if err != nil {
		fmt.Println("Got error while fetching all role permissions from service layer", err)
		return nil, err
	}
	return rolePermissions, nil
}
