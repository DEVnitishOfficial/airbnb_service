package service

import (
	db "AuthInGo/db/repositories"
	"fmt"
)

// interface for user service
type UserService interface {
	CreateUser() error
}

// The UserServiceImpl struct has a field named userRepository.
//
//	This field can hold any object that satisfies the db.UserRepository interface
type UserServiceImpl struct {
	userRepository db.UserRepository // dependency injection of UserRepository
}

// constructor for UserService
func NewUserService(_userRepository db.UserRepository) UserService {
	return &UserServiceImpl{
		userRepository: _userRepository,
	}
}

// CreateUser method implementation
func (u *UserServiceImpl) CreateUser() error {
	fmt.Println("Creating user in UserService")
	u.userRepository.Create() // calling the Create method of UserRepository
	return nil
}
