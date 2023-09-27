package dto

type IndexPermissionsParams struct {
}

type CreatePermissionsParams struct {
	//Handle  string `json:"handle" validate:"required"`
	PermissionsName string `json:"permissions_name" validate:"required"`
	Path            string `json:"path" validate:"required"`
	Action          string `json:"action" validate:"required"`
	Status          uint8  `json:"status" validate:"required"`
	Remark          string `json:"remark" `
}

type ListPermissionsParams struct {
	Handle          string          `json:"handle" validate:"omitempty"`
	PermissionsName string          `json:"permissions_name" validate:"omitempty"`
	Path            string          `json:"path" validate:"omitempty"`
	Action          string          `json:"action" validate:"omitempty"`
	Status          uint8           `json:"status" validate:"omitempty"`
	Remark          string          `json:"remark" `
	Sort            PermissionsSort `json:"sore"`
	Page            uint64          `json:"page" validate:"required,numeric"`
	PageSize        uint64          `json:"page_size" validate:"required,numeric"`
}

type PermissionsSort struct {
	Handle          string `json:"handle" validate:"omitempty,oneof=asc desc"`
	PermissionsName string `json:"permissions_name" validate:"omitempty,oneof=asc desc"`
	Path            string `json:"path" validate:"omitempty,oneof=asc desc"`
}

type DelPermissionsParams struct {
	Ids []uint64 `json:"ids" validate:"required,numeric"`
}

type UpdatePermissionsParams struct {
	PermissionsId   uint64 `json:"permissions_id" validate:"required,numeric"`
	Handle          string `json:"handle" validate:"required"`
	PermissionsName string `json:"permissions_name" validate:"required"`
	Path            string `json:"path" validate:"required"`
	Action          string `json:"action" validate:"required"`
	Status          uint8  `json:"status" validate:"required"`
	Remark          string `json:"remark"`
}

type LinkMenuPermissionsParams struct {
	PermissionsId uint64 `json:"permissions_id" validate:"required,numeric"`
	MenuId        uint64 `json:"menu_id" validate:"required,numeric"`
}
