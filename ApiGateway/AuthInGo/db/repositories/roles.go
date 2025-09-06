package db

import (
	"AuthInGo/models"
	"database/sql"
	"fmt"
)

// 1. Interface (The Contract)
//
//	interface RoleRepository {
//	    getById(id: string): Promise<Role | null>;
//	    // ... other methods
//	}
type RoleRepository interface {
	GetById(id int64) (*models.Role, error)
	GetRoleByName(name string) (*models.Role, error)
	GetAllRoles() ([]*models.Role, error)
	CreateRole(name string, description string) (*models.Role, error)
	DeleteRoleById(id int64) error
	UpdateRoleById(id int64, name string, description string) (*models.Role, error)
}

// A Go struct is a composite data type that just holds data (your db connection). A TypeScript class
//  combines both data (properties like db) and behavior (methods like getById) into a single blueprint.
// 2. Struct (The Blueprint of the "Class")
// 2. In TS ---> Class (The Blueprint of the "Class" and Methods)
// class UserRepositoryImpl implements UserRepository {
//     private db: any;
// }

type RoleRepositoryImpl struct {
	db *sql.DB
}

// 3.Object Creation. Go has a convention of using a New function as a constructor or factory.
//
//	This function returns an instance of the struct but with the interface type
//
// (return &RoleRepositoryImpl{}). TypeScript uses the standard new keyword with the class name
//
//	(new RoleRepositoryImpl(...)).
//	 3. Constructor/Factory Function
//	    constructor(db: any) {
//	    this.db = db;
//	    }
func NewRoleRepository(_db *sql.DB) RoleRepository {
	return &RoleRepositoryImpl{
		db: _db,
	}
}

// 4. Method Implementation (Attached to the struct)
// These methods "implement" the interface implicitly
//
//	async getById(id: string): Promise<User | null> {
//		// Implementation logic here
//		return null;
//	}
func (r *RoleRepositoryImpl) GetById(id int64) (*models.Role, error) {
	query := "SELECT id, name, description, created_at, updated_at FROM roles WHERE id = ?"

	row := r.db.QueryRow(query, id)

	role := &models.Role{}

	err := row.Scan(&role.ID, &role.RoleName, &role.Description, &role.CreatedAt, &role.UpdatedAt)

	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No role found with the given id")
			return nil, err
		} else {
			fmt.Println("Got error while fetching role", err)
			return nil, err
		}
	}
	fmt.Println("Role fetched successfully:", role)
	return role, nil
}

func (r *RoleRepositoryImpl) GetRoleByName(name string) (*models.Role, error) {
	query := "SELECT id, name, description, created_at, updated_at FROM roles WHERE name = ?"
	row := r.db.QueryRow(query, name)
	role := &models.Role{}
	err := row.Scan(&role.ID, &role.RoleName, &role.Description, &role.CreatedAt, &role.UpdatedAt)
	if err != nil {
		return nil, err
	}
	return role, nil
}

func (r *RoleRepositoryImpl) GetAllRoles() ([]*models.Role, error) {
	query := "SELECT id, name, description, created_at, updated_at FROM roles"
	rows, err := r.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var roles []*models.Role
	for rows.Next() {
		role := &models.Role{}
		if err := rows.Scan(&role.ID, &role.RoleName, &role.Description, &role.CreatedAt, &role.UpdatedAt); err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}
	return roles, nil
}

func (r *RoleRepositoryImpl) CreateRole(name string, description string) (*models.Role, error) {
	query := "INSERT INTO roles (name, description) VALUES (?, ?)"
	result, err := r.db.Exec(query, name, description)
	if err != nil {
		return nil, err
	}
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	role := &models.Role{
		ID:          lastInsertID,
		RoleName:    name,
		Description: description,
	}
	return role, nil
}

func (r *RoleRepositoryImpl) DeleteRoleById(id int64) error {
	query := "DELETE FROM roles WHERE id = ?"
	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return fmt.Errorf("no role deleted")
	}
	return nil
}

func (r *RoleRepositoryImpl) UpdateRoleById(id int64, name string, description string) (*models.Role, error) {
	query := "UPDATE roles SET name = ?, description = ? WHERE id = ?"
	result, err := r.db.Exec(query, name, description, id)
	if err != nil {
		return nil, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return nil, err
	}
	if rowsAffected == 0 {
		return nil, fmt.Errorf("no role updated")
	}
	role := &models.Role{
		ID:          id,
		RoleName:    name,
		Description: description,
	}
	return role, nil
}
