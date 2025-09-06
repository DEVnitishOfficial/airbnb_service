package service

import (
	db "AuthInGo/db/repositories"
	"AuthInGo/models"
	"fmt"
)

type RoleService interface {
	GetRoleByIdService(id int64) (*models.Role, error)
	GetRoleByNameService(name string) (*models.Role, error)
	GetAllRolesService() ([]*models.Role, error)
	CreateRoleService(name string, description string) (*models.Role, error)
	DeleteRoleByIdService(id int64) error
	UpdateRoleByIdService(id int64, name string, description string) (*models.Role, error)
}

type RoleServiceImpl struct {
	roleRepository db.RoleRepository
}

func NewRoleService(_roleRepository db.RoleRepository) RoleService {
	return &RoleServiceImpl{
		roleRepository: _roleRepository,
	}
}

func (r *RoleServiceImpl) GetRoleByIdService(id int64) (*models.Role, error) {
	fmt.Println("Fetching user role from service layer")

	role, error := r.roleRepository.GetById(id)

	if error != nil {
		fmt.Println("Got error while fetching role from service layer", error)
		return nil, error
	}
	return role, nil
}

func (r *RoleServiceImpl) GetRoleByNameService(name string) (*models.Role, error) {
	fmt.Println("Fetching role by name from service layer")
	role, err := r.roleRepository.GetRoleByName(name)
	if err != nil {
		fmt.Println("Got error while fetching role by name from service layer", err)
		return nil, err
	}
	fmt.Println("Role fetched successfully:", role)
	return role, nil
}

func (r *RoleServiceImpl) GetAllRolesService() ([]*models.Role, error) {
	fmt.Println("Fetching all roles from service layer")
	roles, err := r.roleRepository.GetAllRoles()
	if err != nil {
		fmt.Println("Got error while fetching all roles from service layer", err)
		return nil, err
	}
	fmt.Println("All roles fetched successfully:", roles)
	return roles, nil
}

func (r *RoleServiceImpl) CreateRoleService(name string, description string) (*models.Role, error) {
	fmt.Println("Creating role in service layer")
	role, err := r.roleRepository.CreateRole(name, description)
	if err != nil {
		fmt.Println("Got error while creating role in service layer", err)
		return nil, err
	}
	fmt.Println("Role created successfully:", role)
	return role, nil
}

func (r *RoleServiceImpl) DeleteRoleByIdService(id int64) error {
	fmt.Println("Deleting role in service layer")
	err := r.roleRepository.DeleteRoleById(id)
	if err != nil {
		fmt.Println("Got error while deleting role in service layer", err)
		return err
	}
	return nil
}

func (r *RoleServiceImpl) UpdateRoleByIdService(id int64, name string, description string) (*models.Role, error) {
	fmt.Println("Updating role in service layer")
	role, err := r.roleRepository.UpdateRoleById(id, name, description)
	if err != nil {
		fmt.Println("Got error while updating role in service layer", err)
		return nil, err
	}
	fmt.Println("Role updated successfully:", role)
	return role, nil
}
