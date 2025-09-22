package db

import (
	"AuthInGo/models"
	"database/sql"
	"fmt"
	"strings"
)

type UserRoleRepository interface {
	GetUserRoles(userId int64) ([]*models.Role, error)
	GetAllUserAndTheirRoles() ([]*models.UserRole, error)
	AssignRoleToUser(userId int64, roleId int64) error
	RemoveRoleFromUser(userId int64, roleId int64) error
	HasRole(userId int64, roleName string) (bool, error)
	HasAllRoles(userId int64, roleNames []string) (bool, error)
	HasAnyRole(userId int64, roleNames []string) (bool, error)
}

type UserRoleRepositoryImpl struct {
	db *sql.DB
}

func NewUserRoleRepository(_db *sql.DB) UserRoleRepository {
	return &UserRoleRepositoryImpl{
		db: _db,
	}
}

func (u *UserRoleRepositoryImpl) GetUserRoles(userId int64) ([]*models.Role, error) {
	query := `
		SELECT u.id, u.username, u.email, r.id, r.name, r.description
		FROM user_roles ur
		INNER JOIN roles r ON ur.role_id = r.id
		JOIN users u ON ur.user_id = u.id
		WHERE ur.user_id = ?`
	rows, err := u.db.Query(query, userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var roles []*models.Role
	for rows.Next() {
		role := &models.Role{}
		if err := rows.Scan(&role.UserId, &role.UserName, &role.Email, &role.RoleId, &role.RoleName, &role.Description); err != nil {
			return nil, err
		}
		roles = append(roles, role)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return roles, nil
}

func (u *UserRoleRepositoryImpl) GetAllUserAndTheirRoles() ([]*models.UserRole, error) {

	query := `SELECT u.id as userId, u.username, u.email, r.id as roleId, r.name, r.description
		FROM user_roles ur
		INNER JOIN roles r ON ur.role_id = r.id
		JOIN users u ON ur.user_id = u.id`

	rows, err := u.db.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var userRoles []*models.UserRole
	for rows.Next() {
		userRole := &models.UserRole{}
		if err := rows.Scan(&userRole.UserId, &userRole.UserName, &userRole.Email, &userRole.RoleId, &userRole.RoleName, &userRole.Description); err != nil {
			return nil, err
		}
		userRoles = append(userRoles, userRole)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return userRoles, nil
}

func (u *UserRoleRepositoryImpl) AssignRoleToUser(userId int64, roleId int64) error {
	query := "INSERT INTO user_roles (user_id, role_id) VALUES (?, ?)"
	_, err := u.db.Exec(query, userId, roleId)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserRoleRepositoryImpl) RemoveRoleFromUser(userId int64, roleId int64) error {
	query := "DELETE FROM user_roles WHERE user_id = ? AND role_id = ?"
	_, err := u.db.Exec(query, userId, roleId)
	if err != nil {
		return err
	}
	return nil
}

func (u *UserRoleRepositoryImpl) HasRole(userId int64, roleName string) (bool, error) {
	query := `
		SELECT COUNT(*) > 0
		FROM user_roles ur
		INNER JOIN roles r ON ur.role_id = r.id
		WHERE ur.user_id = ? AND r.name = ?`
	var exists bool
	err := u.db.QueryRow(query, userId, roleName).Scan(&exists)
	if err != nil {
		return false, err
	}
	return exists, nil
}

func (u *UserRoleRepositoryImpl) HasAllRoles(userId int64, roleNames []string) (bool, error) {

	if len(roleNames) == 0 {
		return true, nil // If no roles are specified, return true
	}

	fmt.Println("roleNames in repository layer:", roleNames)

	// 1. Construct the placeholders for the IN clause
	// e.g., if len(roleNames) is 3, this creates "?,?,?"
	placeholders := make([]string, len(roleNames))
	for i := range placeholders {
		placeholders[i] = "?"
	}
	inClause := strings.Join(placeholders, ",")

	// 2. Build the full query
	query := fmt.Sprintf(`
	SELECT COUNT(*) = ?
	FROM user_roles ur
	INNER JOIN roles r ON ur.role_id = r.id
	WHERE ur.user_id = ? AND r.name IN (%s)
	GROUP BY ur.user_id`, inClause)

	// 3. Prepare the arguments for the query
	// The first argument is the number of roles (len(roleNames))
	// The second argument is the userId
	// The remaining arguments are the elements of the roleNames slice
	args := make([]interface{}, 0, 2+len(roleNames)) // defining the capacity of slice
	args = append(args, len(roleNames), userId)
	for _, roleName := range roleNames {
		args = append(args, roleName)
	}

	// 4. Execute the query with the prepared arguments
	row := u.db.QueryRow(query, args...)

	var hasAllRoles bool
	if err := row.Scan(&hasAllRoles); err != nil {
		if err == sql.ErrNoRows {
			fmt.Println("No rows returned from query, user does not have all roles.")
			return false, nil
		}
		return false, err
	}

	return hasAllRoles, nil
}

func (u *UserRoleRepositoryImpl) HasAnyRole(userId int64, roleNames []string) (bool, error) {

	if len(roleNames) == 0 {
		return true, nil // If no roles are specified, return true
	}
	placeholders := strings.Repeat("?,", len(roleNames)) // it create parameter placeholders(?) based on how many role passed to it.
	placeholders = placeholders[:len(placeholders)-1]    // remove the last comma("?","?", to "?","?")
	query := fmt.Sprintf("SELECT COUNT(*) > 0 FROM user_roles ur INNER JOIN roles r ON ur.role_id = r.id WHERE ur.user_id = ? AND r.name IN (%s)", placeholders)

	/* finally query will look like
	   SELECT COUNT(*) > 0
	   FROM user_roles ur
	   INNER JOIN roles r ON ur.role_id = r.id
	   WHERE ur.user_id = 42
	   AND r.name IN ('user', 'admin');

	*/

	// Create args slice with userId first, then all roleNames
	args := make([]interface{}, 0, 1+len(roleNames)) // creating a slice just like array in js.
	args = append(args, userId)
	for _, roleName := range roleNames {
		args = append(args, roleName)
	}

	row := u.db.QueryRow(query, args...)

	var hasAnyRole bool
	if err := row.Scan(&hasAnyRole); err != nil {
		if err == sql.ErrNoRows {
			return false, nil // No roles found for the user
		}
		return false, err // Return any other error
	}

	fmt.Println("hasAnyRole", hasAnyRole)

	return hasAnyRole, nil
}
