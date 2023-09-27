package middleware

import (
	"context"
	"fmt"
	"ginDome/global"
	"ginDome/service/admin/dto"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"net/http"
)

var RouterParams map[string]interface{}

func ValidatorMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		path := ctx.Request.URL.Path
		params := getParams(path)
		jsonRes := dto.JsonResult{
			Code: http.StatusInternalServerError,
			Data: nil,
		}
		// 请求前业务逻辑...
		if ctx.Request.ContentLength == 0 {
			jsonRes.Msg = "当前请求参数不能为空！"
			ctx.JSONP(http.StatusInternalServerError, jsonRes)
			ctx.Abort() //停止往下执行
			return
		}
		err := ctx.ShouldBindJSON(&params)
		if err != nil {
			jsonRes.Msg = err.Error()
			ctx.JSONP(http.StatusInternalServerError, jsonRes)
			ctx.Abort() //停止往下执行
			return
		}
		err = global.GvaValidate.Struct(params)
		if err != nil {
			// 处理验证错误
			validationErrors := err.(validator.ValidationErrors)
			for _, e := range validationErrors {
				errStr := fmt.Sprintf("Validation error for field %s: %s\n", e.Field(), e.Tag())
				jsonRes.Msg = errStr
				ctx.JSONP(http.StatusInternalServerError, jsonRes)
				ctx.Abort() //停止往下执行
				return
			}

		}

		c := context.WithValue(ctx.Request.Context(), global.GvaParams, params)
		ctx.Request = ctx.Request.WithContext(c)
		// 继续往下执行
		ctx.Next()
		// 请求后业务逻辑...
	}
}

func getParams(path string) interface{} {
	RouterParams = make(map[string]interface{})
	//user
	RouterParams["/v1/user/login"] = &dto.LoginParams{}
	RouterParams["/v1/user/list"] = &dto.UserListParams{}
	RouterParams["/v1/user/register"] = &dto.RegisterParams{}
	RouterParams["/v1/user/update"] = &dto.UpdateUserParams{}
	//permissions
	RouterParams["/v1/permissions/Index"] = &dto.IndexPermissionsParams{}
	RouterParams["/v1/permissions/create"] = &dto.CreatePermissionsParams{}
	RouterParams["/v1/permissions/del"] = &dto.DelPermissionsParams{}
	RouterParams["/v1/permissions/update"] = &dto.UpdatePermissionsParams{}
	RouterParams["/v1/permissions/list"] = &dto.ListPermissionsParams{}
	//role
	RouterParams["/v1/role/Index"] = &dto.IndexRoleParams{}
	RouterParams["/v1/role/create"] = &dto.CreateRoleParams{}
	RouterParams["/v1/role/del"] = &dto.DelRoleParams{}
	RouterParams["/v1/role/update"] = &dto.UpdateRoleParams{}
	RouterParams["/v1/role/list"] = &dto.ListRoleParams{}
	RouterParams["/v1/role/linkUser"] = &dto.LinkUserParams{}
	RouterParams["/v1/role/linkPermissions"] = &dto.RoleLinkPermissionsParams{}
	//menu
	RouterParams["/v1/menu/Index"] = &dto.IndexMenuParams{}
	RouterParams["/v1/menu/create"] = &dto.CreateMenuParams{}
	RouterParams["/v1/menu/del"] = &dto.DeleteMenuParams{}
	RouterParams["/v1/menu/update"] = &dto.UpdateMenuParams{}
	RouterParams["/v1/menu/list"] = &dto.MenuListParams{}
	RouterParams["/v1/menu/linkRole"] = &dto.LinkMenuRoleParams{}
	RouterParams["/v1/menu/linkPermissions"] = &dto.LinkMenuPermissionsParams{}

	return RouterParams[path]
}
