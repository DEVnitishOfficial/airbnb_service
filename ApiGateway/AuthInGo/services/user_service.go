package service

import (
	db "AuthInGo/db/repositories"
)

// interface for user service
type UserService interface {
	CreateUser() error
}

// implementation of user service or simply it's a class that implements UserService interface
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
	// Implementation for creating a user
	return nil
}
