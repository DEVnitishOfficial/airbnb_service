package dto

type AssignRoleRequestDTO struct {
	RoleId int64 `json:"roleId" validate:"required"`
}
type RemoveRoleRequestDTO struct {
	RoleId int64 `json:"roleId" validate:"required"`
}
