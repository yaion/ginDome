package entity

import (
	"fmt"
	"ginDome/global"
	"ginDome/service/admin/dto"
)

type Role struct {
	RoleId   uint64 `gorm:"primarKey" json:"role_id"`
	RoleName string `json:"role_name"`
	Status   uint8  `json:"status"`
	Remark   string `json:"remark"`
	Model
}

// Create menu
func (r *Role) Create() error {
	return global.GvaMysqlClient.Create(r).Error
}

// Update 更新
func (r *Role) Update() error {
	return global.GvaMysqlClient.Where("status = 1 role_id = ? ", r.RoleId).Save(r).Error
}

// Del 批量删除
func (r *Role) Del(del *[]Role) error {
	return global.GvaMysqlClient.Delete(del).Error
}

// List 列表 todo
func (r *Role) List(params *dto.ListRoleParams) ([]Role, error) {
	var result []Role
	db := global.GvaMysqlClient
	if params.RoleName != "" {
		db = db.Where("role_name like ?", "%"+params.RoleName+"%")
	}
	if params.Remark != "" {
		db = db.Where("remark like ?", "%"+params.Remark+"%")
	}
	if params.Status != 0 {
		db = db.Where("level = ?", params.Status)
	}
	//排序
	if params.Sort.RoleName != "" {
		db = db.Order(fmt.Sprintf("role_name %s", params.Sort.RoleName))
	}

	db = db.Offset(int((params.Page - 1) * params.PageSize)).Limit(int(params.Page * params.PageSize))
	err := db.Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
	return result, nil
}

// LinkUserRole 用户连接角色 todo
func LinkUserRole() {

}

// LinkRolePermissions 用户连接权限 todo
func LinkRolePermissions() {

}
