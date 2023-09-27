package dto

type IndexRoleParams struct {
}

type CreateRoleParams struct {
	RoleName string `json:"role_name"`
	Status   uint8  `json:"status"`
	Remark   string `json:"remark"`
}

type DelRoleParams struct {
	Ids []uint64 `json:"ids" validate:"required,numeric"`
}

type UpdateRoleParams struct {
	RoleId   uint64 `json:"role_id"`
	RoleName string `json:"role_name"`
	Status   uint8  `json:"status"`
	Remark   string `json:"remark"`
}

type ListRoleParams struct {
	RoleName string   `json:"role_name"`
	Status   uint8    `json:"status"`
	Remark   string   `json:"remark"`
	Sort     RoleSort `json:"sore"`
	Page     uint64   `json:"page" validate:"required,numeric"`
	PageSize uint64   `json:"page_size" validate:"required,numeric"`
}

type RoleSort struct {
	RoleName string `json:"role_name" validate:"omitempty,oneof=asc desc"`
}

type LinkUserParams struct {
}

type RoleLinkPermissionsParams struct {
}
