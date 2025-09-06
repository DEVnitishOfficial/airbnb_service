package dto

type CreateRoleRequestDto struct {
	RoleName    string `json:"roleName" validate:"required,min=3,max=20"`
	Description string `json:"description" validate:"required,min=3,max=200"`
}

type UpdateRoleRequestDto struct {
	RoleName    string `json:"roleName" validate:"required,min=3,max=20"`
	Description string `json:"description" validate:"required,min=3,max=200"`
}
