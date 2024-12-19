package admin

import (
	"ginDome/global"
	"ginDome/model/entity"
	"ginDome/service/admin"
	"ginDome/service/admin/dto"
	"ginDome/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

// Login v1 登录
func Login(ctx *gin.Context) {
	//获取请求参数
	params, ok := ctx.Request.Context().Value(global.GvaParams).(*dto.LoginParams)
	if !ok {
		ctx.JSONP(http.StatusInternalServerError, gin.H{"error": "系统错误"})
		return
	}
	ctx.JSON(http.StatusOK, admin.Login(params))

}

func UserIndex(ctx *gin.Context) {
	user, ok := ctx.Request.Context().Value(global.GvaUserData).(*entity.User)
	if !ok {
		ctx.JSONP(http.StatusInternalServerError, gin.H{"error": "系统错误"})
		return
	}
	ctx.JSON(http.StatusOK, admin.UserIndex(user))
}

func Register(ctx *gin.Context) {
	params, ok := ctx.Request.Context().Value(global.GvaParams).(*dto.RegisterParams)
	if !ok {
		ctx.JSONP(http.StatusInternalServerError, gin.H{"error": "系统错误"})
	}
	user := entity.User{}
	user.UserName = params.UserName
	user.Password = utils.Md5str(params.Password)
	user.NickName = params.NickName
	user.Phone = params.Phone
	user.Email = params.Email
	user.Sex = params.Sex
	user.Remark = params.Remark
	user.CreatedTime = time.Now()
	user.ModifiedTime = time.Now()
	ctx.JSON(http.StatusOK, admin.Register(&user))
}

// UserList 用户列表
func UserList(ctx *gin.Context) {
	//获取参数
	params, ok := ctx.Request.Context().Value(global.GvaParams).(*dto.UserListParams)
	if !ok {
		ctx.JSONP(http.StatusInternalServerError, gin.H{"error": "系统错误"})
		return
	}
	ctx.JSONP(http.StatusOK, admin.UserList(params))
}

// Destroy 注销
func Destroy(ctx *gin.Context) {
	//获取参数
	user, ok := ctx.Request.Context().Value(global.GvaUserData).(*entity.User)
	if !ok {
		ctx.JSONP(http.StatusInternalServerError, gin.H{"error": "系统错误"})
		return
	}
	ctx.JSONP(http.StatusOK, admin.Destroy(user))
}

// UpdateUser  更新用户信息
func UpdateUser(ctx *gin.Context) {
	//获取参数
	params, ok := ctx.Request.Context().Value(global.GvaParams).(*dto.UpdateUserParams)
	if !ok {
		//todo 记录日志
		ctx.JSONP(http.StatusInternalServerError, gin.H{"error": "系统错误"})
		return
	}
	user, ok := ctx.Request.Context().Value(global.GvaUserData).(*entity.User)
	if !ok {
		//todo 记录日志
		ctx.JSONP(http.StatusInternalServerError, gin.H{"error": "系统错误"})
		return
	}
	param := make(map[string]interface{})
	// 先设置为必填，以后有需要在做判断
	if params.UserName != "" {
		//user.UserName = params.UserName
		param["user_name"] = params.UserName
	}
	if params.Password != "" {
		//user.Password = utils.Md5str(params.Password)
		param["password"] = params.Password
	}
	if params.NickName != "" {
		user.NickName = params.NickName
	}
	if params.Phone != "" {
		user.Phone = params.Phone
	}
	if params.Avatar != "" {
		user.Avatar = params.Avatar
	}
	if params.Sex != 0 {
		user.Sex = params.Sex
	}
	if params.Email != "" {
		user.Email = params.Email
	}
	if params.Status != 0 {
		user.Status = params.Status
	}
	if params.Remark != "" {
		user.Remark = params.Remark
	}
	ctx.JSONP(http.StatusOK, admin.UpdateUser(user))
}
