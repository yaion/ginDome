package admin

import (
	"fmt"
	"ginDome/model/entity"
	"ginDome/service/admin/dto"
	"ginDome/utils"
	"net/http"
)

// Login v1 登录逻辑
func Login(params *dto.LoginParams) dto.JsonResult {
	var data dto.LoginBack
	var result dto.JsonResult
	var err error
	// 验证验证码 todo
	//验证密码
	user := entity.User{UserName: params.Name}
	err = user.GetUserOne()
	if err != nil {
		result.Code = http.StatusInternalServerError
		result.Msg = err.Error()
		return result
	}
	md5str := utils.Md5str(params.Password)
	if user.Password != md5str {
		result.Code = http.StatusInternalServerError
		result.Msg = "login err please try again!"
		return result
	}
	//发放jwt
	data.Token, err = utils.CreateJwt(user.UserId)
	if err != nil {
		result.Code = http.StatusInternalServerError
		result.Msg = "System error please try again!"
		return result
	}
	data.Msg = "login success!"
	result.Code = http.StatusOK
	result.Msg = "success"
	result.Data = data
	return result
}

func UserIndex(user *entity.User) *dto.JsonResult {
	return &dto.JsonResult{
		Code: http.StatusOK,
		Msg:  "success",
		Data: user,
	}
}

func Register(user *entity.User) *dto.JsonResult {
	result := dto.JsonResult{}
	err := user.Create()
	if err != nil {
		//记录日志 todo
		//global.GvaLogger.Error("error",zap.)
		result.Code = 4001
		result.Msg = "Register fail!"
		return &result
	}
	result.Code = http.StatusOK
	result.Msg = "success"
	result.Data = nil
	return &result
}

// UserList 用户列表
func UserList(params *dto.UserListParams) *dto.JsonResult {
	result := dto.JsonResult{}
	user := entity.User{}
	users, err := user.GetUserList(params)
	if err != nil {
		//记录日志 todo
		//global.GvaLogger.Error("error",zap.)
		result.Code = 4001
		result.Msg = "System error please try again!"
		return &result
	}
	result.Code = http.StatusOK
	result.Msg = "success"
	result.Data = users
	return &result
}

// Destroy 注销
func Destroy(user *entity.User) *dto.JsonResult {
	result := dto.JsonResult{}
	err := user.Delete()
	if err != nil {
		//记录日志 todo
		//global.GvaLogger.Error("error",zap.)
		fmt.Println(err)
		result.Code = 4001
		result.Msg = "Register fail!"
		return &result
	}
	result.Code = http.StatusOK
	result.Msg = "destroy success！"
	result.Data = nil
	return &result
}

func UpdateUser(user *entity.User) *dto.JsonResult {
	result := dto.JsonResult{}
	err := user.Update()
	if err != nil {
		//记录日志 todo
		//global.GvaLogger.Error("error",zap.)
		result.Code = 4001
		result.Msg = "update fail!"
		return &result
	}
	result.Code = http.StatusOK
	result.Msg = "update success！"
	result.Data = nil
	return &result
}

// Verify 验证码 生产验证码放回 todo
func Verify() {

}
