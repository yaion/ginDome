package dto

type IndexMenuParams struct {
}

type CreateMenuParams struct {
	MenuName string `json:"menu_name" validate:"required"`
	Pid      uint64 `json:"pid" validate:"required,numeric"`
	Level    uint8  `json:"level" validate:"required,numeric"`
	Remark   string `json:"remark" `
}

type MenuListParams struct {
	MenuName string   `json:"menu_name" validate:"omitempty"`
	Level    uint8    `json:"level" validate:"omitempty,numeric"`
	Remark   string   `json:"remark"`
	Sort     MenuSort `json:"sort"`
	Page     uint64   `json:"page" validate:"required,numeric"`
	PageSize uint64   `json:"page_size" validate:"required,numeric"`
}

type MenuSort struct {
	MenuName string `json:"menu_name" validate:"omitempty,oneof=asc desc"`
	Level    string `json:"level" validate:"omitempty,oneof=asc desc"`
}

type UpdateMenuParams struct {
	MenuId   uint64 `json:"menu_id" validate:"required,numeric"`
	MenuName string `json:"menu_name" validate:"required"`
	Pid      uint64 `json:"pid" validate:"required,numeric"`
	Level    uint8  `json:"level" validate:"required,numeric"`
	Status   uint8  `json:"status" validate:"required,numeric"`
	Remark   string `json:"remark" `
}

type DeleteMenuParams struct {
	MenuIds []int64 `json:"menu_id" validate:"required,dive,gt=0"`
}

type MenuLinkPermissionsParams struct {
}

type LinkMenuRoleParams struct {
}
