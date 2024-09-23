package admin

import (
	"ginDome/global"
	"ginDome/model/entity"
	"ginDome/service/admin"
	"ginDome/service/admin/dto"
	"ginDome/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// todo
func IndexPermissions(ctx *gin.Context) {
	params, ok := ctx.Request.Context().Value(global.GvaParams).(*dto.IndexPermissionsParams)
	if !ok {
		//todo 记录日志
		ctx.JSONP(http.StatusInternalServerError, gin.H{"msg": "system err!"})
		return
	}
	ctx.JSONP(http.StatusOK, admin.IndexPermissions(params))
}

func CreatePermissions(ctx *gin.Context) {
	params, ok := ctx.Request.Context().Value(global.GvaParams).(*dto.CreatePermissionsParams)
	if !ok {
		//todo 记录日志
		ctx.JSONP(http.StatusInternalServerError, gin.H{"msg": "system err!"})
		return
	}
	//拼接权限数据
	permission := entity.Permissions{}
	permission.PermissionsName = params.PermissionsName
	permission.Status = 1
	routerData, err := utils.GetRouterData("params.Path")
	if err != nil {
		//todo 记录日志
		ctx.JSONP(http.StatusInternalServerError, gin.H{"msg": "Please enter correct parameters path!"})
		return
	}
	permission.Handle = routerData.Handle
	permission.Path = params.Path
	permission.Action = routerData.Action
	permission.Remark = params.Remark
	ctx.JSONP(http.StatusOK, admin.CreatePermissions(&permission))
}

func DeletePermissions(ctx *gin.Context) {
	params, ok := ctx.Request.Context().Value(global.GvaParams).(*dto.DelPermissionsParams)
	permission := make([]entity.Permissions, len(params.Ids))
	for v, _ := range params.Ids {
		perm := entity.Permissions{PermissionsId: uint64(v)}
		permission = append(permission, perm)
	}
	if !ok {
		//todo 记录日志
		ctx.JSONP(http.StatusInternalServerError, gin.H{"msg": "system err!"})
		return
	}
	ctx.JSONP(http.StatusOK, admin.DeletePermissions(permission))
}

func UpdatePermissions(ctx *gin.Context) {
	params, ok := ctx.Request.Context().Value(global.GvaParams).(*dto.UpdatePermissionsParams)
	if !ok {
		//todo 记录日志
		ctx.JSONP(http.StatusInternalServerError, gin.H{"msg": "system err!"})
		return
	}
	//拼接权限数据
	permission := entity.Permissions{}
	permission.PermissionsId = params.PermissionsId
	permission.PermissionsName = params.PermissionsName
	permission.Status = params.Status
	routerData, err := utils.GetRouterData("params.Path")
	if err != nil {
		//todo 记录日志
		ctx.JSONP(http.StatusInternalServerError, gin.H{"msg": "Please enter correct parameters path!"})
		return
	}
	permission.Handle = routerData.Handle
	permission.Path = params.Path
	permission.Action = routerData.Action
	permission.Remark = params.Remark
	ctx.JSONP(http.StatusOK, admin.CreatePermissions(&permission))
}

func ListPermissions(ctx *gin.Context) {
	params, ok := ctx.Request.Context().Value(global.GvaParams).(*dto.ListPermissionsParams)
	if !ok {
		//todo 记录日志
		ctx.JSONP(http.StatusInternalServerError, gin.H{"msg": "system err!"})
		return
	}
	ctx.JSONP(http.StatusOK, admin.ListPermissions(params))
}

//todo
/*func LinkMenu(ctx *gin.Context){
	params,ok := ctx.Request.Context().Value(global.GvaParams).(*dto.LinkMenuPermissionsParams)
	if !ok{
		//todo 记录日志
		ctx.JSONP(http.StatusInternalServerError,gin.H{"msg":"system err!"})
		return
	}
	ctx.JSONP(http.StatusOK,admin.LinkMenu(params))
}*/
