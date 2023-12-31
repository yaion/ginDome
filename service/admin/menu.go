package admin

import (
	"ginDome/model/entity"
	"ginDome/service/admin/dto"
	"net/http"
)

func CreateMenu(menu *entity.Menu) dto.JsonResult {
	result := dto.JsonResult{}
	err := menu.Create()
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

func UpdateMenu(menu *entity.Menu) dto.JsonResult {
	result := dto.JsonResult{}
	err := menu.Update()
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

func DeleteMenu(menus *[]entity.Menu) dto.JsonResult {
	result := dto.JsonResult{}
	menu := entity.Menu{}
	err := menu.Del(menus)
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

func ListMenu(params *dto.MenuListParams) dto.JsonResult {
	result := dto.JsonResult{}
	menu := entity.Menu{}
	list, err := menu.List(params)
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
