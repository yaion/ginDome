package entity

import (
	"fmt"
	"ginDome/global"
	"ginDome/service/admin/dto"
)

type User struct {
	UserId   uint64 `gorm:"primarKey" json:"user_id"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
	NickName string `json:"nick_name"`
	Phone    string `json:"phone"`
	Avatar   string `json:"avatar"`
	Sex      uint8  `json:"sex"`
	Email    string `json:"email"`
	Status   uint8  `json:"status"`
	Remark   string `json:"remark"`
	Model
}

func (user *User) TableName() string {
	return "gd_user"
}

func (user *User) GetUserOne() error {
	user.Status = 1
	return global.GvaMysqlClient.Where(user).Find(user).Error
}

func (user *User) Create() error {
	return global.GvaMysqlClient.Create(user).Error
}

func (user *User) Delete() error {
	return global.GvaMysqlClient.Where("user_id = ?", user.UserId).Delete(user).Error
}

func (user *User) Update() error {
	return global.GvaMysqlClient.Where("status = 1 user_id = ? ", user.UserId).Save(user).Error
}

func (user *User) GetUserList(params *dto.UserListParams) (*map[string]interface{}, error) {
	var users []User
	var total int64
	var err error
	result := make(map[string]interface{})
	db := global.GvaMysqlClient
	if params.UserName != "" {
		db = db.Where("user_name like ?", "%"+params.UserName+"%")
	}
	if params.NickName != "" {
		db = db.Where("nick_name like ?", "%"+params.NickName+"%")
	}
	if params.Phone != "" {
		db = db.Where("phone = ?", params.Phone)
	}
	//排序
	if params.Sort.UserName != "" {
		db = db.Order(fmt.Sprintf("user_name %s", params.Sort.UserName))
	}
	if params.Sort.NickName != "" {
		db = db.Order(fmt.Sprintf("nick_name %s", params.Sort.NickName))
	}
	if params.Sort.Phone != "" {
		db = db.Order(fmt.Sprintf("phone %s", params.Sort.Phone))
	}
	if params.Sort.Id != "" {
		db = db.Order(fmt.Sprintf("user_id %s", params.Sort.Id))
	}
	// todo 优化
	err = db.Table(user.TableName()).Where("deleted_at is null").Count(&total).Error
	if err != nil {
		return nil, err
	}
	db = db.Offset(int((params.Page - 1) * params.PageSize)).Limit(int(params.Page * params.PageSize))
	err = db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	result["list"] = users
	result["total"] = total
	result["page"] = params.Page
	result["page_size"] = params.PageSize
	return &result, nil
}
