package db

import (
	"AuthInGo/models"
	"database/sql"
	"fmt"
)

type PermissionRepository interface {
	GetPermissionById(id int64) (*models.Permission, error)
	GetPermissionByName(name string) (*models.Permission, error)
	GetAllPermissions() ([]*models.Permission, error)
	CreatePermission(name string, description string, resource string, action string) (*models.Permission, error)
	DeletePermissionById(id int64) error
	UpdatePermissionById(id int64, name string, description string, resource string, action string) (*models.Permission, error)
}

type PermissionRepositoryImpl struct {
	db *sql.DB
}

func NewPermissionRepository(_db *sql.DB) PermissionRepository {
	return &PermissionRepositoryImpl{
		db: _db,
	}
}

func (p *PermissionRepositoryImpl) GetPermissionById(id int64) (*models.Permission, error) {
	query := "SELECT id, name, description, resource, action, created_at, updated_at FROM permissions WHERE id = ?"
	row := p.db.QueryRow(query, id)
	permission := &models.Permission{}
	err := row.Scan(&permission.ID, &permission.PermissionName, &permission.Description, &permission.Resource, &permission.Action, &permission.CreatedAt, &permission.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No permission found with the given id")
			return nil, err
		} else {
			fmt.Println("Got error while fetching permission", err)
			return nil, err
		}
	}
	return permission, nil
}

func (p *PermissionRepositoryImpl) GetPermissionByName(name string) (*models.Permission, error) {
	query := "SELECT id, name, description, resource, action, created_at, updated_at FROM permissions WHERE name = ?"
	row := p.db.QueryRow(query, name)
	permission := &models.Permission{}
	err := row.Scan(&permission.ID, &permission.PermissionName, &permission.Description, &permission.Resource, &permission.Action, &permission.CreatedAt, &permission.UpdatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No permission found with the given name")
			return nil, err
		} else {
			fmt.Println("Got error while fetching permission", err)
			return nil, err
		}
	}
	return permission, nil
}

func (p *PermissionRepositoryImpl) GetAllPermissions() ([]*models.Permission, error) {
	query := "SELECT id, name, description, resource, action, created_at, updated_at FROM permissions"
	rows, err := p.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var permissions []*models.Permission
	for rows.Next() {
		permission := &models.Permission{}
		if err := rows.Scan(&permission.ID, &permission.PermissionName, &permission.Description, &permission.Resource, &permission.Action, &permission.CreatedAt, &permission.UpdatedAt); err != nil {
			return nil, err
		}
		permissions = append(permissions, permission)
	}
	return permissions, nil
}

func (p *PermissionRepositoryImpl) CreatePermission(name string, description string, resource string, action string) (*models.Permission, error) {
	query := "INSERT INTO permissions (name, description, resource, action) VALUES (?, ?, ?, ?)"
	result, err := p.db.Exec(query, name, description, resource, action)
	if err != nil {
		return nil, err
	}
	lastInsertID, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}
	permission := &models.Permission{
		ID:             lastInsertID,
		PermissionName: name,
		Description:    description,
		Resource:       resource,
		Action:         action,
	}
	return permission, nil
}

func (p *PermissionRepositoryImpl) DeletePermissionById(id int64) error {
	query := "DELETE FROM permissions WHERE id = ?"
	result, err := p.db.Exec(query, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("Got error while fetching rows affected", err)
		return err
	}
	if rowsAffected == 0 {
		fmt.Println("No permission deleted")
		return fmt.Errorf("no permission deleted")
	}
	return nil
}

func (p *PermissionRepositoryImpl) UpdatePermissionById(id int64, name string, description string, resource string, action string) (*models.Permission, error) {
	query := "UPDATE permissions SET name = ?, description = ?, resource = ?, action = ? WHERE id = ?"
	result, err := p.db.Exec(query, name, description, resource, action, id)
	if err != nil {
		fmt.Println("Got error while executing update command Exec")
		return nil, err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		fmt.Println("Got error while fetching rows affected", err)
		return nil, err
	}
	if rowsAffected == 0 {
		return nil, fmt.Errorf("no permission updated")
	}
	return &models.Permission{
		ID:             id,
		PermissionName: name,
		Description:    description,
		Resource:       resource,
		Action:         action,
	}, nil
}
