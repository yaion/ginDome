package admin

import (
	"ginDome/global"
	"ginDome/model/entity"
	"ginDome/service/admin"
	"ginDome/service/admin/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

// todo
func IndexRole(ctx *gin.Context) {

}

func CreateRole(ctx *gin.Context) {
	params, ok := ctx.Request.Context().Value(global.GvaParams).(*dto.CreateRoleParams)
	if !ok {
		//todo 记录日志
		ctx.JSONP(http.StatusInternalServerError, gin.H{"msg": "system err!"})
		return
	}
	role := entity.Role{
		Remark:   params.Remark,
		RoleName: params.RoleName,
		Status:   params.Status,
	}
	ctx.JSONP(http.StatusOK, admin.CreateRole(&role))
}

func DeleteRole(ctx *gin.Context) {
	params, ok := ctx.Request.Context().Value(global.GvaParams).(*dto.DelRoleParams)
	roles := make([]entity.Role, len(params.Ids))
	for v, _ := range params.Ids {
		perm := entity.Role{RoleId: uint64(v)}
		roles = append(roles, perm)
	}
	if !ok {
		//todo 记录日志
		ctx.JSONP(http.StatusInternalServerError, gin.H{"msg": "system err!"})
		return
	}
	ctx.JSONP(http.StatusOK, admin.DeleteRole(&roles))
}

func UpdateRole(ctx *gin.Context) {
	params, ok := ctx.Request.Context().Value(global.GvaParams).(*dto.UpdateRoleParams)
	if !ok {
		//todo 记录日志
		ctx.JSONP(http.StatusInternalServerError, gin.H{"msg": "system err!"})
		return
	}
	role := entity.Role{
		RoleId:   params.RoleId,
		RoleName: params.RoleName,
		Status:   params.Status,
		Remark:   params.Remark,
	}
	ctx.JSONP(http.StatusOK, admin.UpdateRole(&role))
}

func ListRole(ctx *gin.Context) {
	params, ok := ctx.Request.Context().Value(global.GvaParams).(*dto.ListRoleParams)
	if !ok {
		//todo 记录日志
		ctx.JSONP(http.StatusInternalServerError, gin.H{"msg": "system err!"})
		return
	}
	ctx.JSONP(http.StatusOK, admin.ListRole(params))
}

func LinkUser(ctx *gin.Context) {
	/*params,ok := ctx.Request.Context().Value(global.GvaParams).(*dto.Link)
	if !ok{
		//todo 记录日志
		ctx.JSONP(http.StatusInternalServerError,gin.H{"msg":"system err!"})
		return
	}
	ctx.JSONP(http.StatusOK,admin.LinkUser(params))*/
}

func LinkPermissions(ctx *gin.Context) {
	/*params,ok := ctx.Request.Context().Value(global.GvaParams).(*dto.Link)
	if !ok{
		//todo 记录日志
		ctx.JSONP(http.StatusInternalServerError,gin.H{"msg":"system err!"})
		return
	}
	ctx.JSONP(http.StatusOK,admin.LinkPermissions(params))*/
}
