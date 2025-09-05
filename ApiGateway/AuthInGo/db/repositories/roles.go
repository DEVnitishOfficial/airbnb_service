package db

import (
	"AuthInGo/models"
	"database/sql"
	"fmt"
)

type RoleRepository interface {
	GetById(id int64) (*models.Role, error)
}

type RoleRepositoryImpl struct {
	db *sql.DB
}

func NewRoleRepository(_db *sql.DB) RoleRepository {
	return &RoleRepositoryImpl{
		db: _db,
	}
}

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
