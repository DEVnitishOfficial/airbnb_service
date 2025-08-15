package service

import (
	db "AuthInGo/db/repositories"
	"AuthInGo/utils"
	"fmt"
)

// interface for user service
type UserService interface {
	GetUserById() error
	CreateUser() error
	LoginUser() error
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
func (u *UserServiceImpl) GetUserById() error {
	fmt.Println("Fetching user in UserService")
	u.userRepository.GetById()
	return nil
}

func (u *UserServiceImpl) CreateUser() error {
	fmt.Println("Creating user from userService")
	password := "pass4user4"
	hashedPassword, err := utils.HashedPassword(password)
	if err != nil {
		fmt.Println("Unable to hash password")
		return err
	}
	u.userRepository.Create(
		"user4",
		"user4@gmail.com",
		hashedPassword,
	)
	return nil
}

func (u *UserServiceImpl) LoginUser() error {
	response := utils.CheckPasswordHash("pass4user444", "$2a$10$r3QS7wjYQ4qbZDDDbWbRsOL7UpIsgkb.VU935IX5Xe2VihNMV9bfa")
	fmt.Println("Response result", response)
	return nil
}
