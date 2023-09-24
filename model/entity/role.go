package entity

type Role struct {
	RoleId uint64 `gorm:"primarKey""`
	RoleName string
	Status uint8
	Remark string
	Model
}
