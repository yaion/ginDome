package admin

import (
	"ginDome/model/entity"
	"ginDome/service/admin/dto"
	"net/http"
)

func CreateRole(role *entity.Role) dto.JsonResult {
	result := dto.JsonResult{}
	err := role.Create()
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

func DeleteRole(roles *[]entity.Role) dto.JsonResult {
	result := dto.JsonResult{}
	role := entity.Role{}
	err := role.Del(roles)
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

func UpdateRole(role *entity.Role) dto.JsonResult {
	result := dto.JsonResult{}
	err := role.Update()
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

func ListRole(params *dto.ListRoleParams) dto.JsonResult {
	var result dto.JsonResult
	role := entity.Role{}
	list, err := role.List(params)
	if err != nil {
		//todo 记录日志
		result.Code = http.StatusInternalServerError
		result.Msg = "System error please try again!"
		return result
	}

	result.Code = http.StatusOK
	result.Msg = "success"
	result.Data = list
	return result
}

func LinkUser() {

}

func LinkPermissions() {

}
