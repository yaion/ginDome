package entity

type Permissions struct {
	PermissionsId uint64 `gorm:"primarKey"`
	Handle  string
	PermissionsName string
	Path string
	Action string
	Status uint8
	Remark string
	Model
}
