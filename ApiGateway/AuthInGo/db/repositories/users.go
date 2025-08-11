package db

import (
	"fmt"
)

type UserRepository interface {
	Create() error
	// Other methods like Get, Update, Delete can be added here
}

type UserRepositoryImpl struct {
	// db *sql.DB
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{
		// db: db,
	}
}

func (u *UserRepositoryImpl) Create() error {
	fmt.Println("Creating user in UserRepository")
	return nil
}
