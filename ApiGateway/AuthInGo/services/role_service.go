package service

import (
	db "AuthInGo/db/repositories"
	"AuthInGo/models"
	"fmt"
)

type RoleService interface {
	GetRoleById(id int64) (*models.Role, error)
}

type RoleServiceImpl struct {
	roleRepository db.RoleRepository
}

func NewRoleService(_roleRepository db.RoleRepository) RoleService {
	return &RoleServiceImpl{
		roleRepository: _roleRepository,
	}
}

func (r *RoleServiceImpl) GetRoleById(id int64) (*models.Role, error) {
	fmt.Println("Fetching user role from service layer")

	role, error := r.roleRepository.GetById(id)

	if error != nil {
		fmt.Println("Got error while fetching role from service layer", error)
		return nil, error
	}
	return role, nil
}
