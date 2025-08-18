package service

import (
	env "AuthInGo/config/env"
	db "AuthInGo/db/repositories"
	"AuthInGo/dto"
	"AuthInGo/utils"
	"fmt"

	// "github.com/golang-jwt/jwt"
	"github.com/golang-jwt/jwt/v5"
)

// interface for user service
type UserService interface {
	GetUserById() error
	CreateUser() error
	LoginUser(payload *dto.LoginUserRequestDto) (string, error)
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

func (u *UserServiceImpl) LoginUser(payload *dto.LoginUserRequestDto) (string, error) {
	email := payload.Email
	password := payload.Password
	// step:1 Make a repository call to get the user by email
	fmt.Println("Getting user by email")
	user, err := u.userRepository.GetByEmail(email)

	// step:2 if user not exist, return the error
	if err != nil {
		fmt.Println("User not found")
		return "", err
	}

	// step:3 if user exist, check password using utils.checkPasswordHash
	if !utils.CheckPasswordHash(password, user.Password) {
		fmt.Println("Invalid password")
		return "", fmt.Errorf("invalid password")
	}

	// step:4 if password matches, print a JWT token else return error saying password not match

	jwtPayload := jwt.MapClaims{
		"email": user.Email,
		"id":    user.ID,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtPayload)
	tokenString, err := token.SignedString([]byte(env.GetString("JWT_SECRET", "DNKN_TOKEN")))

	if err != nil {
		fmt.Println("Got error while generating token", err)
		return "", err
	}

	fmt.Println("JWT TOKEN", tokenString)
	return tokenString, nil

}
