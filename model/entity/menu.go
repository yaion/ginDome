package entity

import (
	"fmt"
	"ginDome/global"
	"ginDome/service/admin/dto"
)

type Menu struct {
	MenuId   uint64 `json:"menu_id" gorm:"primarKey"`
	MenuName string `json:"menu_name" `
	Pid      uint64 `json:"pid"`
	Level    uint8  `json:"level"`
	Status   uint8  `json:"status"`
	Remark   string `json:"remark"`
	Model
}

// Create menu
func (m *Menu) Create() error {
	return global.GvaMysqlClient.Create(m).Error
}

// Update 更新
func (m *Menu) Update() error {
	return global.GvaMysqlClient.Where("status = 1 menu_id = ? ", m.MenuId).Save(m).Error
}

// Del 批量删除
func (m *Menu) Del(del *[]Menu) error {
	return global.GvaMysqlClient.Delete(del).Error
}

// List 列表 todo
func (m *Menu) List(params *dto.MenuListParams) ([]Menu, error) {
	var result []Menu
	db := global.GvaMysqlClient
	if params.MenuName != "" {
		db = db.Where("menu_name like ?", "%"+params.MenuName+"%")
	}
	if params.Remark != "" {
		db = db.Where("remark like ?", "%"+params.Remark+"%")
	}
	if params.Level != 0 {
		db = db.Where("level = ?", params.Level)
	}
	//排序
	if params.Sort.MenuName != "" {
		db = db.Order(fmt.Sprintf("menu_name %s", params.Sort.MenuName))
	}
	if params.Sort.Level != "" {
		db = db.Order(fmt.Sprintf("level %s", params.Sort.Level))
	}

	db = db.Offset(int((params.Page - 1) * params.PageSize)).Limit(int(params.Page * params.PageSize))
	err := db.Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}

// LinkMenuRole 菜单连接角色 todo
func LinkMenuRole() {

}

// LinkMenuPermissions 菜单连接权限 todo
func LinkMenuPermissions() {

}
