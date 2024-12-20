package admin

import (
	"ginDome/global"
	"ginDome/model/entity"
	"ginDome/service/admin"
	"ginDome/service/admin/dto"
	"github.com/gin-gonic/gin"
	"net/http"
)

func IndexMenu(ctx *gin.Context) {

}

func CreateMenu(ctx *gin.Context) {
	params, ok := ctx.Request.Context().Value(global.GvaParams).(*dto.CreateMenuParams)
	if !ok {
		//todo 记录日志
		ctx.JSONP(http.StatusInternalServerError, gin.H{"msg": "system err!"})
		return
	}
	menu := entity.Menu{
		MenuName: params.MenuName,
		Pid:      params.Pid,
		Level:    params.Level,
		Remark:   params.Remark,
	}
	ctx.JSONP(http.StatusOK, admin.CreateMenu(&menu))
}

func DeleteMenu(ctx *gin.Context) {
	params, ok := ctx.Request.Context().Value(global.GvaParams).(*dto.DeleteMenuParams)
	menus := make([]entity.Menu, len(params.MenuIds))
	i := 0
	for _, v := range params.MenuIds {
		menu := entity.Menu{MenuId: uint64(v)}
		menus[i] = menu
		i++
	}
	if !ok {
		//todo 记录日志
		ctx.JSONP(http.StatusInternalServerError, gin.H{"msg": "system err!"})
		return
	}
	ctx.JSONP(http.StatusOK, admin.DeleteMenu(&menus))
}

func UpdateMenu(ctx *gin.Context) {
	params, ok := ctx.Request.Context().Value(global.GvaParams).(*dto.UpdateMenuParams)
	if !ok {
		//todo 记录日志
		ctx.JSONP(http.StatusInternalServerError, gin.H{"msg": "system err!"})
		return
	}
	role := entity.Menu{
		MenuId:   params.MenuId,
		MenuName: params.MenuName,
		Pid:      params.Pid,
		Level:    params.Level,
		Status:   params.Status,
		Remark:   params.Remark,
	}
	ctx.JSONP(http.StatusOK, admin.UpdateMenu(&role))
}

func ListMenu(ctx *gin.Context) {
	params, ok := ctx.Request.Context().Value(global.GvaParams).(*dto.MenuListParams)
	if !ok {
		//todo 记录日志
		ctx.JSONP(http.StatusInternalServerError, gin.H{"msg": "system err!"})
		return
	}

	ctx.JSONP(http.StatusOK, admin.ListMenu(params))
}
