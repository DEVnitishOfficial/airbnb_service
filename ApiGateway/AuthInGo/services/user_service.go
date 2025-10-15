package service

import (
	env "AuthInGo/config/env"
	db "AuthInGo/db/repositories"
	"AuthInGo/dto"
	"AuthInGo/models"
	"AuthInGo/utils"
	"fmt"
	"strconv"

	// "github.com/golang-jwt/jwt"
	"github.com/golang-jwt/jwt/v5"
)

// interface for user service
type UserService interface {
	GetAllUserService() ([]*models.User, error)
	GetUserById(id string) (*models.User, error)
	CreateUser(payload *dto.CreateUserRequestDto) (*models.User, error)
	LoginUser(payload *dto.LoginUserRequestDto) (string, error)
	GetBulkUserInfoByIdsService(ids []int64) (map[int64]*models.User, error)
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

// service layer of getAll user

func (u *UserServiceImpl) GetAllUserService() ([]*models.User, error) {
	fmt.Println("Fetching all users from service layer")
	user, err := u.userRepository.GetAll()
	if err != nil {
		fmt.Println("Got error while fetching user from service layer", err)
		return nil, err
	}
	return user, nil
}

// GetUserById is a method that fetches a user by ID using the userRepository
func (u *UserServiceImpl) GetUserById(id string) (*models.User, error) {
	fmt.Println("Fetching user in UserService")
	user, err := u.userRepository.GetById(id)
	if err != nil {
		fmt.Println("Got error while fetching user from service layer", err)
		return nil, err
	}
	return user, nil
}

func (u *UserServiceImpl) CreateUser(payload *dto.CreateUserRequestDto) (*models.User, error) {
	fmt.Println("Creating user from userService")

	// hashing password using utils.HashPassword
	hashedPassword, err := utils.HashedPassword(payload.Password)

	if err != nil {
		fmt.Println("Unable to hash password")
		return nil, err
	}

	// call repository layer to create the user
	createdUser, err := u.userRepository.Create(payload.Username, payload.Email, hashedPassword)
	if err != nil {
		fmt.Println("Error while creating user", err)
		return nil, err
	}

	// return the created user
	return createdUser, nil
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

func (u *UserServiceImpl) GetBulkUserInfoByIdsService(ids []int64) (map[int64]*models.User, error) {
	fmt.Println("Fetching bulk user info by IDs in UserService")

	users := make(map[int64]*models.User)
	for _, id := range ids {
		user, err := u.userRepository.GetById(strconv.FormatInt(id, 10))
		if err != nil {
			fmt.Println("Error fetching user by ID:", err)
			return nil, err
		}
		users[id] = user
	}
	return users, nil
}
