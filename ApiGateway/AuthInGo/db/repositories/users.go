package db

import (
	"AuthInGo/models"
	"database/sql"
	"fmt"
)

type UserRepository interface {
	GetById(id string) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	Create(username string, email string, hashedPassword string) (*models.User, error)
	GetAll() ([]*models.User, error)
	DeleteById(id int64) error
	UpdateById(id int64, user *models.User) (*models.User, error)
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

	query := "SELECT id, username, email, created_at, updated_at FROM users"
	rows, err := u.db.Query(query)

	if err != nil {
		fmt.Println("Got error while fetching user", err)
		return nil, err
	}

	defer rows.Close()

	var users []*models.User
	for rows.Next() {
		user := &models.User{}
		if err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.CreatedAt, &user.UpdatedAt); err != nil {
			fmt.Println("Got error while scanning user", err)
			return nil, err
		}
		users = append(users, user)
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error with rows:", err)
		return nil, err
	}

	return users, nil
}

func (u *UserRepositoryImpl) DeleteById(id int64) error {

	return nil
}

func (u *UserRepositoryImpl) UpdateById(id int64, user *models.User) (*models.User, error) {
	query := "UPDATE users SET username = ?, email = ?, password = ? WHERE id = ?"
	result, err := u.db.Exec(query, user.Username, user.Email, user.Password, id)

	if err != nil {
		fmt.Println("Error updating user:", err)
		return nil, err
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("Error fetching rows affected:", err)
		return nil, err
	}
	if rowsAffected == 0 {
		fmt.Println("No user updated with the given ID")
		return nil, fmt.Errorf("no user updated")
	}

	updatedUser := &models.User{
		Username: user.Username,
		Email:    user.Email,
		Password: user.Password,
	}

	return updatedUser, nil
}

func (u *UserRepositoryImpl) Create(username string, email string, hashedPassword string) (*models.User, error) {
	query := "INSERT INTO users (username, email, password) VALUES (?, ?, ?)"
	result, err := u.db.Exec(query, username, email, hashedPassword)

	fmt.Println("printing the  result after executing query>>>>", result)

	if err != nil {
		fmt.Println("Error inserting user:", err)
		return nil, err
	}

	lastInsertID, rowErr := result.LastInsertId()
	if rowErr != nil {
		fmt.Println("Got error while fetching lastInsertID", rowErr)
		return nil, rowErr
	}

	user := &models.User{
		ID:       lastInsertID,
		Username: username,
		Email:    email,
	}

	fmt.Println("User created successfully", user)

	return user, nil
}

func (u *UserRepositoryImpl) GetById(id string) (*models.User, error) {
	fmt.Println("Fetching user in UserRepository")

	// step 1 : Prepare the SQL query
	query := "SELECT id, username, email, created_at, updated_at FROM users WHERE id = ?"
	// step 2 : Execute the query
	row := u.db.QueryRow(query, id)
	// step 3 : Process the result
	// it simply creates a user object and scans the row into it
	user := &models.User{}

	// here row.Scan will scan the row into the user object
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

func (u *UserRepositoryImpl) GetByEmail(email string) (*models.User, error) {

	query := "SELECT id, email, password from users where email = ?"

	row := u.db.QueryRow(query, email)

	user := &models.User{}

	err := row.Scan(&user.ID, &user.Email, &user.Password)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No user found with the given email")
			return nil, err
		} else {
			fmt.Println("Error fetching user:", err)
			return nil, err
		}
	}
	fmt.Println("User fetched successfully:", user)
	return user, nil
}
