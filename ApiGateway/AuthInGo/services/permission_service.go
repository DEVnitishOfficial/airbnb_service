package service

import (
	db "AuthInGo/db/repositories"
	"AuthInGo/models"
	"fmt"
)

// similir to typescript interface
type PermissionService interface {
	GetPermissionByIdService(id int64) (*models.Permission, error)
	GetPermissionByNameService(name string) (*models.Permission, error)
	GetAllPermissionsService() ([]*models.Permission, error)
	CreatePermissionService(name string, description string, resource string, action string) (*models.Permission, error)
	UpdatePermissionByIdService(id int64, name string, description string, resource string, action string) (*models.Permission, error)
	DeletePermissionByIdService(id int64) error
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

func (p *PermissionServiceImpl) GetPermissionByIdService(id int64) (*models.Permission, error) {
	fmt.Println("Fetching permission using id from service layer")

	permission, error := p.permissionRepository.GetPermissionById(id)

	if error != nil {
		fmt.Println("Got error while fetching permisssion from service layer", error)
		return nil, error
	}
	return permission, nil
}

func (p *PermissionServiceImpl) GetPermissionByNameService(name string) (*models.Permission, error) {
	fmt.Println("Fetching permission using name from service layer")

	permission, error := p.permissionRepository.GetPermissionByName(name)

	if error != nil {
		fmt.Println("Got error while fetching permisssion from service layer", error)
		return nil, error
	}
	return permission, nil
}

func (p *PermissionServiceImpl) GetAllPermissionsService() ([]*models.Permission, error) {
	fmt.Println("Fetching all permission from the service layer")

	allPermission, error := p.permissionRepository.GetAllPermissions()

	if error != nil {
		fmt.Println("Got error while fetching permission from service layer")
		return nil, error
	}
	fmt.Println("All permissions fetched successfully", allPermission)
	return allPermission, nil
}

func (p *PermissionServiceImpl) CreatePermissionService(name string, description string, resource string, action string) (*models.Permission, error) {
	fmt.Println("create permission called in service layer")
	createdPermission, error := p.permissionRepository.CreatePermission(name, description, resource, action)

	if createdPermission != nil {
		fmt.Println("Got error while creating permission form service layer")
		return nil, error
	}

	fmt.Println("Permission created successfully")
	return createdPermission, nil
}

func (p *PermissionServiceImpl) UpdatePermissionByIdService(id int64, name string, description string, resource string, action string) (*models.Permission, error) {
	fmt.Println("Update permission called in service layer")
	updatedPermission, error := p.permissionRepository.UpdatePermissionById(id, name, description, resource, action)

	if error != nil {
		fmt.Println("Got error while updating permission from service layer")
		return nil, error
	}

	fmt.Println("Permission updated successfully")
	return updatedPermission, nil
}

func (p *PermissionServiceImpl) DeletePermissionByIdService(id int64) error {
	fmt.Println("Delete permission called in service layer")
	error := p.permissionRepository.DeletePermissionById(id)

	if error != nil {
		fmt.Println("Got error while deleting permission from service layer", error)
		return error
	}
	return nil
}
