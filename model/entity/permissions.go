package entity

import (
	"fmt"
	"ginDome/global"
	"ginDome/service/admin/dto"
)

type Permissions struct {
	PermissionsId   uint64 `gorm:"primarKey"`
	Handle          string
	PermissionsName string
	Path            string
	Action          string
	Status          uint8
	Remark          string
	Model
}

// Create  创建权限
func (perm *Permissions) Create() error {
	return global.GvaMysqlClient.Create(perm).Error
}

// Update 更新
func (perm *Permissions) Update() error {
	return global.GvaMysqlClient.Where("status = 1 permissions_id = ? ", perm.PermissionsId).Save(perm).Error
}

// Del 批量删除
func (perm *Permissions) Del(del []Permissions) error {
	return global.GvaMysqlClient.Delete(del).Error
}

// List 列表
func (perm *Permissions) List(params *dto.ListPermissionsParams) ([]Permissions, error) {
	var result []Permissions
	db := global.GvaMysqlClient
	if params.Handle != "" {
		db = db.Where("handle like ?", "%"+params.Handle+"%")
	}
	if params.PermissionsName != "" {
		db = db.Where("permissions like ?", "%"+params.PermissionsName+"%")
	}
	if params.Path != "" {
		db = db.Where("path = ?", params.Path)
	}
	//排序
	if params.Sort.Handle != "" {
		db = db.Order(fmt.Sprintf("handle %s", params.Sort.Handle))
	}
	if params.Sort.PermissionsName != "" {
		db = db.Order(fmt.Sprintf("permissions_name %s", params.Sort.PermissionsName))
	}
	if params.Sort.Path != "" {
		db = db.Order(fmt.Sprintf("path %s", params.Sort.Path))
	}
	db = db.Offset(int((params.Page - 1) * params.PageSize)).Limit(int(params.Page * params.PageSize))
	err := db.Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}
