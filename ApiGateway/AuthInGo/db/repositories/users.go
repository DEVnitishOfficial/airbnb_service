package db

import (
	"AuthInGo/models"
	"database/sql"
	"fmt"
	"strings"
)

// Similir kind of implementation done in typescript for better understanding of go code.
/*
// The `User` model, analogous to your Go `models.User`
interface User {
    id: number;
    username: string;
    email: string;
    // ... other properties
}

// 1. The Interface (The Contract) - Exactly like Go
interface UserRepository {
    getById(id: string): Promise<User | null>;
    getByEmail(email: string): Promise<User | null>;
    create(username: string, email: string, hashedPassword: string): Promise<User>;
    getAll(): Promise<User[]>;
    deleteById(id: number): Promise<void>;
    updateById(id: number, user: User): Promise<User>;
}

// 2. The Class (Implementation) - The equivalent of your `UserRepositoryImpl` struct + methods
class UserRepositoryImpl implements UserRepository {
    // A private property for the database connection
    private db: any;

    // The constructor, equivalent to your `NewUserRepository` function
    constructor(db: any) {
        this.db = db;
    }

    // The method implementations, defined directly inside the class
    async getById(id: string): Promise<User | null> {
        // Your implementation logic here
        return null;
    }

    async getByEmail(email: string): Promise<User | null> {
        // Your implementation logic here
        return null;
    }

    // And so on for all other methods...
    async create(username: string, email: string, hashedPassword: string): Promise<User> {
      // Your implementation logic here
      return { id: 1, username, email } as User;
    }

    async getAll(): Promise<User[]> {
      // Your implementation logic here
      return [];
    }

    async deleteById(id: number): Promise<void> {
        // Your implementation logic here
    }

    async updateById(id: number, user: User): Promise<User> {
        // Your implementation logic here
        return user;
    }
}

// 3. Object Creation - The equivalent of calling your `NewUserRepository` function
const myDbConnection = {}; // Assume this is your database client
const userRepository: UserRepository = new UserRepositoryImpl(myDbConnection);
*/

// similir to typescript interface
type UserRepository interface {
	GetById(id string) (*models.User, error)
	GetByEmail(email string) (*models.User, error)
	Create(username string, email string, hashedPassword string) (*models.User, error)
	GetAll() ([]*models.User, error)
	DeleteById(id int64) error
	UpdateById(id int64, user *models.User) (*models.User, error)
	GetBulkUserInfoByIds(ids []int64) (map[int64]*models.User, error)
}

// similir to class in typescript, here in go it only take properties not consturctor and methods
type UserRepositoryImpl struct {
	db *sql.DB
}

// this is similir to the typescript consturctor, here it is not defined inside the struct rather it
// is attached to the struct using &UserRepositoryImpl
func NewUserRepository(_db *sql.DB) UserRepository {
	return &UserRepositoryImpl{
		db: _db,
	}
}

// this is similir to methods definiton in typescript which is written inside the class but here in
// golang it is attached to the class/struct using *UserRepositoryImpl
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

	userID, _ := result.LastInsertId()
	roleName := "user"

	// get role ID dynamically
	var roleID int64
	err = u.db.QueryRow("SELECT id FROM roles WHERE name = ?", roleName).Scan(&roleID)
	if err != nil {
		fmt.Println("Error fetching role ID:", err)
		return nil, err
	}

	// assign role
	_, err = u.db.Exec("INSERT INTO user_roles (user_id, role_id) VALUES (?, ?)", userID, roleID)
	if err != nil {
		fmt.Println("Error assigning role:", err)
		return nil, err
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

// Get userInfo by ID in bulk for ReviewService
func (u *UserRepositoryImpl) GetBulkUserInfoByIds(ids []int64) (map[int64]*models.User, error) {
	fmt.Println("Fetching users in UserRepository by IDs:", ids)
	if len(ids) == 0 {
		return map[int64]*models.User{}, nil
	}
	// Create a query with the appropriate number of placeholders
	placeholders := make([]string, len(ids))
	args := make([]interface{}, len(ids))
	for i, id := range ids {
		placeholders[i] = "?"
		args[i] = id
	}
	query := fmt.Sprintf("SELECT id, username, email FROM users WHERE id IN (%s)", strings.Join(placeholders, ","))

	rows, err := u.db.Query(query, args...)
	if err != nil {
		fmt.Println("Got error while fetching users by IDs", err)
		return nil, err
	}
	defer rows.Close()

	userMap := make(map[int64]*models.User)
	for rows.Next() {
		user := &models.User{}
		if err := rows.Scan(&user.ID, &user.Username, &user.Email); err != nil {
			fmt.Println("Error scanning user:", err)
			return nil, err
		}
		userMap[user.ID] = user
	}

	if err := rows.Err(); err != nil {
		fmt.Println("Error with rows:", err)
		return nil, err
	}

	return userMap, nil
}
