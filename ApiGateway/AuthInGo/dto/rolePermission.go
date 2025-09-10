package dto

type AssignPermissionRequestDTO struct {
	PermissionId int64 `json:"permissionId" validate:"required"`
}

type RemovePermissionRequestDTO struct {
	PermissionId int64 `json:"permissionId" validate:"required"`
}
