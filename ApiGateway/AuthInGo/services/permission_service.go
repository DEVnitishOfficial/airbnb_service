package service

import (
	db "AuthInGo/db/repositories"
	"AuthInGo/models"
)

// similir to typescript interface
type PermissionService interface {
	GetPermissionByIdService(id int64) (*models.Permission, error)
	GetPermissionByNameService(name string) (*models.Permission, error)
	GetAllPermissionsService() ([]*models.Permission, error)
	CreatePermissionService(name string, description string, resource string, action string) (*models.Permission, error)
	DeletePermissionByIdService(id int64) error
	UpdatePermissionByIdService(id int64, name string, description string, resource string, action string) (*models.Permission, error)
}

// similir to typescript class, here we can only define the properties not methods
type PermissionServiceImpl struct {
	permissionRepository db.PermissionRepository
}

// similir to typescript constructor in golang it is attached to the class/struct rather defining
// inside the class/sturct using &PermissionServiceImpl
func NewPermissionService(_permissionRepository db.PermissionRepository) PermissionService {
	return &PermissionServiceImpl{
		permissionRepository: _permissionRepository,
	}
}
