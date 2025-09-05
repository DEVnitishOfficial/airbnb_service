package models

type Role struct {
	ID          int64
	RoleName    string
	Description string
	CreatedAt   string
	UpdatedAt   string
}

type Permission struct {
	ID             int64
	PermissionName string
	Description    string
	Resource       string
	Action         string
	CreatedAt      string
	UpdatedAt      string
}

type RolePermission struct {
	ID           int64
	RoleId       int64
	PermissionId int64
	CreatedAt    string
	UpdatedAt    string
}

type UserRole struct {
	ID        int64
	UserId    int64
	RoleId    int64
	CreatedAt string
	UpdatedAt string
}
