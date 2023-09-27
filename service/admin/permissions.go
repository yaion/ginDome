package admin

import (
	"fmt"
	"ginDome/model/entity"
	"ginDome/service/admin/dto"
	"net/http"
)

// todo
func IndexPermissions(params *dto.IndexPermissionsParams) dto.JsonResult {
	result := dto.JsonResult{}
	fmt.Println(params)
	return result
}

// CreatePermissions 创建权限
func CreatePermissions(permissions *entity.Permissions) dto.JsonResult {
	result := dto.JsonResult{}
	err := permissions.Create()
	if err != nil {
		//todo 记录日志
		result.Code = http.StatusInternalServerError
		result.Msg = "System error please try again!"
		return result
	}

	result.Code = http.StatusOK
	result.Msg = "success"
	result.Data = nil
	return result
}

func DeletePermissions(permissions []entity.Permissions) dto.JsonResult {
	result := dto.JsonResult{}
	Permissions := entity.Permissions{}
	err := Permissions.Del(permissions)
	if err != nil {
		//todo 记录日志
		result.Code = http.StatusInternalServerError
		result.Msg = "System error please try again!"
		return result
	}

	result.Code = http.StatusOK
	result.Msg = "success"
	result.Data = nil
	return result
}

func ListPermissions(params *dto.ListPermissionsParams) *dto.JsonResult {
	result := dto.JsonResult{}
	perm := entity.Permissions{}
	users, err := perm.List(params)
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

func LinkMenu(params *dto.ListPermissionsParams) *dto.JsonResult {
	result := dto.JsonResult{}
	perm := entity.Permissions{}
	users, err := perm.List(params)
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
