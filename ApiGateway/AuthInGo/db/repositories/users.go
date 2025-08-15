package db

import (
	"AuthInGo/models"
	"database/sql"
	"fmt"
)

type UserRepository interface {
	GetById() (*models.User, error)
	Create(username string, email string, hashedPassword string) error
	GetAll() ([]*models.User, error)
	DeleteById(id int64) error
	UpdateById(id int64, user *models.User) error
}

type UserRepositoryImpl struct {
	db *sql.DB
}

func NewUserRepository(_db *sql.DB) UserRepository {
	return &UserRepositoryImpl{
		db: _db,
	}
}

func (u *UserRepositoryImpl) GetAll() ([]*models.User, error) {
	return nil, nil
}

func (u *UserRepositoryImpl) DeleteById(id int64) error {
	return nil
}

func (u *UserRepositoryImpl) UpdateById(id int64, user *models.User) error {
	return nil
}

func (u *UserRepositoryImpl) Create(username string, email string, hashedPassword string) error {
	query := "INSERT INTO users (username, email, password) VALUES (?, ?, ?)"
	result, err := u.db.Exec(query, username, email, hashedPassword)

	if err != nil {
		fmt.Println("Error inserting user:", err)
		return err
	}
	rowsAffected, rowErr := result.RowsAffected()
	if rowErr != nil {
		fmt.Println("Error getting rows affected:", rowErr)
		return rowErr
	}

	if rowsAffected == 0 {
		fmt.Println("No rows were affected, user not created")
		return nil
	}

	fmt.Println("User created successfully, rows affected:", rowsAffected)
	return nil
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
