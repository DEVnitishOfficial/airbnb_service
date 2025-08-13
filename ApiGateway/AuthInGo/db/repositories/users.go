package db

import (
	"AuthInGo/models"
	"database/sql"
	"fmt"
)

type UserRepository interface {
	GetById() (*models.User, error)
	// Other methods like Get, Update, Delete can be added here
}

type UserRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(_db *sql.DB) UserRepository {
	return &UserRepositoryImpl{
		db: _db,
	}
}

func (u *UserRepositoryImpl) GetById() (*models.User, error) {
	fmt.Println("Fetching user in UserRepository")

	// step 1 : Prepare the SQL query
	query := "SELECT id, username, email, created_at, updated_at FROM users WHERE id = ?"
	// step 2 : Execute the query
	row := u.db.QueryRow(query, 1)
	// step 3 : Process the result
	user := &models.User{}

	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt, &user.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No user found with the given ID")
			return nil, err
		} else {
			fmt.Println("Error fetching user:", err)
			return nil, err
		}
	}
	// step 4 : Return the user
	fmt.Println("User fetched successfully:", user)
	return user, nil
}
