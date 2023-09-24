package router

import (
	admin "ginDome/api/admin"
	"ginDome/router/middleware"
	"ginDome/service"
	"github.com/gin-gonic/gin"
)

// todo 集成Swagger 文档
func InitRouter(e *gin.Engine){
	v1:=e.Group("v1")
	{
		user := v1.Group("user")
		{
			user.POST("login", middleware.ValidatorMiddleware(), admin.Login)
			user.POST("register", middleware.ValidatorMiddleware(), admin.Register)
			//需要登录鉴权
			auth := user.Group("").Use(middleware.JwtMiddleware(),middleware.Auth())
			{
				auth.GET("index", admin.UserIndex)//当前用户信息
				auth.GET("destroy", admin.Destroy)//注销用户
			}
			validator:=user.Group("").Use(middleware.JwtMiddleware(),middleware.ValidatorMiddleware(),middleware.Auth())
			{
				validator.POST("list", admin.UserList)//当前用户列表
				validator.POST("update", admin.UpdateUser)//更新用户数据
			}
		}
		//菜单
		/*menu:=v1.Group("menu").Use(middleware.JwtMiddleware(),middleware.Auth(),middleware.ValidatorMiddleware())
		{
			menu.POST("Index", menu.Index)//当前用户可以看到的菜单列表
			menu.POST("create", menu.Create)//创建菜单
			menu.POST("del", menu.Delete)//删除菜单
			menu.POST("update", menu.Update)//更新菜单
	    }
		//权限 permissions
		permissions:=v1.Group("permissions").Use(middleware.JwtMiddleware(),middleware.Auth(),middleware.ValidatorMiddleware())
		{
			permissions.POST("Index", permissions.Index)//当前用户可以看到的权限列表
			permissions.POST("create", permissions.Create)//创建权限
			permissions.POST("del", permissions.Delete)//删除权限
			permissions.POST("update", permissions.Update)//更新权限
			permissions.POST("menu", permissions.LinkMenu)//权限关联菜单
		}
		//角色 role
		role:=v1.Group("role").Use(middleware.JwtMiddleware(),middleware.Auth(),middleware.ValidatorMiddleware())
		{
			role.POST("Index", role.Index)//当前用户可以看到的角色列表
			role.POST("create", role.Create)//创建角色
			role.POST("del", role.Delete)//删除角色
			role.POST("update", role.Update)//更新角色
			role.POST("user", role.LinkUser)//角色关联用户
			role.POST("permissions", role.LinkPermissions)//角色关联权限
		}
		// 系统相关设置
		system:=v1.Group("system").Use(middleware.JwtMiddleware(),middleware.Auth(),middleware.ValidatorMiddleware())
		{
			system.POST("resource",admin.resource)//获取当前服务器资源及占用 github.com/shirou/gopsutil

		}*/
	}

	//验证码
	e.GET("verify",service.Verify)
	//上传文件
	up:=e.Group("upload")
	{
		up.POST("file",service.UploadFile)
		up.POST("files",service.UploadFiles)
	}

}
