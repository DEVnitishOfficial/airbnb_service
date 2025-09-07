package dto

type CreatePermissionRequestDto struct {
	PermissionName string `json:"permissionName" validate:"required,min=3,max=100"`
	Description    string `json:"description" validate:"required,min=3,max=200"`
	Resource       string `json:"resource" validate:"required,min=4,max=50"`
	Action         string `json:"action" validate:"required,min=4,max=50"`
}

//	type UpdatePermissionRequestDto struct {
//		PermissionName string `json:"permissionName" validate:"required,min=3,max=50"`
//		Description    string `json:"description" validate:"required,min=3,max=200"`
//		Resource       string `json:"resource" validate:"required,min=4,max=50"`
//		Action         string `json:"action" validate:"required,min=4,max=50"`
//	}

type UpdatePermissionRequestDto struct {
	PermissionName string `json:"permissionName"`
	Description    string `json:"description"`
	Resource       string `json:"resource"`
	Action         string `json:"action"`
}
