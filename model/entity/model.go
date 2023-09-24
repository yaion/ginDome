package entity

import (
	"gorm.io/gorm"
	"time"
)

type Model struct {
	Modifier uint64 `json:"modifier"`
	CreatedTime time.Time `json:"created_time"`
	ModifiedTime time.Time `json:"modified_time"`
	Creater uint64 `json:"creater"`
	ModifierName string `json:"modifier_name"`
	CreaterName string `json:"creater_name"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`

}
