package service

import (
	"ginDome/service/admin/dto"
	"ginDome/utils"
	"net/http"
)

func Verify() dto.JsonResult {
	var result dto.JsonResult
	verify, err := utils.CreateVerifyImg(100, 200)
	if err != nil {
		result.Code = http.StatusInternalServerError
		result.Msg = err.Error()
		return result
	}
	result.Code = http.StatusOK
	result.Msg = "success"
	result.Data = verify
	return result
}
