package service

import (
	db "AuthInGo/db/repositories"
	"AuthInGo/models"
	"fmt"
)

type UserRolesService interface {
	// return the list of roles for a user
	GetUserRoleService(userId int64) ([]*models.Role, error)
	GetAllUserAndTheirRolesService() ([]*models.UserRole, error)
	AssignRoleToUserService(userId int64, roleId int64) ([]*models.Role, error)
	RemoveRoleFromUserService(userId int64, roleId int64) error
	HasRoleService(userId int64, roleName string) (bool, error)
	HasAllRolesService(userId int64, roleNames []string) (bool, error)
	HasAnyRoleService(userId int64, roleNames []string) (bool, error)
}

type UserRolesServiceImpl struct {
	userRolesService db.UserRoleRepository
}

func NewUserRolesService(_userRolesService db.UserRoleRepository) UserRolesService {
	return &UserRolesServiceImpl{
		userRolesService: _userRolesService,
	}
}

func (ur *UserRolesServiceImpl) GetUserRoleService(userId int64) ([]*models.Role, error) {
	fmt.Println("Fetching user roles from service layer")
	userRole, err := ur.userRolesService.GetUserRoles(userId)
	if err != nil {
		fmt.Println("Got error while fetching user roles from service layer", err)
		return nil, err
	}
	return userRole, nil
}

func (ur *UserRolesServiceImpl) GetAllUserAndTheirRolesService() ([]*models.UserRole, error) {
	fmt.Println("Fetching all users and their roles from service layer")
	userRoles, err := ur.userRolesService.GetAllUserAndTheirRoles()
	if err != nil {
		fmt.Println("Got error while fetching all users and their roles from service layer", err)
		return nil, err
	}
	return userRoles, nil
}

func (ur *UserRolesServiceImpl) AssignRoleToUserService(userId int64, roleId int64) ([]*models.Role, error) {
	fmt.Println("Assigning role to user from service layer")
	err := ur.userRolesService.AssignRoleToUser(userId, roleId)
	if err != nil {
		fmt.Println("Got error while assigning role to user from service layer", err)
		return nil, err
	}
	return ur.GetUserRoleService(userId)
}

func (ur *UserRolesServiceImpl) RemoveRoleFromUserService(userId int64, roleId int64) error {
	fmt.Println("Removing role from user from service layer")
	err := ur.userRolesService.RemoveRoleFromUser(userId, roleId)
	if err != nil {
		fmt.Println("Got error while removing role from user from service layer", err)
		return err
	}
	return nil
}

func (ur *UserRolesServiceImpl) HasRoleService(userId int64, roleName string) (bool, error) {
	fmt.Println("Checking if user has role from service layer")
	hasRole, err := ur.userRolesService.HasRole(userId, roleName)
	if err != nil {
		fmt.Println("Got error while checking if user has role from service layer", err)
		return false, err
	}
	return hasRole, nil
}

func (ur *UserRolesServiceImpl) HasAllRolesService(userId int64, roleNames []string) (bool, error) {
	fmt.Println("Checking if user has all roles from service layer")
	hasAllRoles, err := ur.userRolesService.HasAllRoles(userId, roleNames)
	if err != nil {
		fmt.Println("Got error while checking if user has all roles from service layer", err)
		return false, err
	}
	return hasAllRoles, nil
}

func (ur *UserRolesServiceImpl) HasAnyRoleService(userId int64, roleNames []string) (bool, error) {
	fmt.Println("Checking if user has any role from service layer")
	hasAnyRole, err := ur.userRolesService.HasAnyRole(userId, roleNames)
	if err != nil {
		fmt.Println("Got error while checking if user has any role from service layer", err)
		return false, err
	}
	return hasAnyRole, nil
}
