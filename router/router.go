package router

import (
	"errors"
	admin "ginDome/api/admin"
	"ginDome/router/middleware"
	"ginDome/service"
	"github.com/gin-gonic/gin"
)

// todo 集成Swagger 文档
func InitRouter(e *gin.Engine) {
	v1 := e.Group("v1")
	{
		user := v1.Group("user")
		{
			user.POST("login", middleware.ValidatorMiddleware(), admin.Login)
			user.POST("register", middleware.ValidatorMiddleware(), admin.Register)
			//需要登录鉴权
			auth := user.Group("").Use(middleware.JwtMiddleware(), middleware.Auth())
			{
				auth.GET("index", admin.UserIndex) //当前用户信息
				auth.GET("destroy", admin.Destroy) //注销用户
			}
			validator := user.Group("").Use(middleware.JwtMiddleware(), middleware.ValidatorMiddleware(), middleware.Auth())
			{
				validator.POST("list", admin.UserList)     //当前用户列表
				validator.POST("update", admin.UpdateUser) //更新用户数据
			}
		}

		//权限 permissions
		permissions := v1.Group("permissions").Use(middleware.JwtMiddleware(), middleware.ValidatorMiddleware())
		{
			permissions.POST("Index", admin.IndexPermissions)   //当前用户可以看到的权限列表 todo
			permissions.POST("create", admin.CreatePermissions) //创建权限
			permissions.POST("del", admin.DeletePermissions)    //删除权限
			permissions.POST("update", admin.UpdatePermissions) //更新权限
			permissions.POST("List", admin.ListPermissions)     //更新权限
			//permissions.POST("linkMenu", admin.LinkMenu)//权限关联菜单
		}
		//角色 role
		role := v1.Group("role").Use(middleware.JwtMiddleware(), middleware.Auth(), middleware.ValidatorMiddleware())
		{
			role.POST("Index", admin.IndexRole)                 //当前用户可以看到的角色列表 todo
			role.POST("create", admin.CreateRole)               //创建角色
			role.POST("del", admin.DeleteRole)                  //删除角色
			role.POST("update", admin.UpdateRole)               //更新角色
			role.POST("list", admin.ListRole)                   //创建角色
			role.POST("linkUser", admin.LinkUser)               //角色关联用户
			role.POST("linkPermissions", admin.LinkPermissions) //角色关联权限
		}
		//菜单
		menu := v1.Group("menu").Use(middleware.JwtMiddleware(), middleware.Auth(), middleware.ValidatorMiddleware())
		{
			menu.POST("Index", admin.IndexMenu)   //当前用户可以看到的菜单列表 todo
			menu.POST("create", admin.CreateMenu) //创建菜单
			menu.POST("del", admin.DeleteMenu)    //删除菜单
			menu.POST("update", admin.UpdateMenu) //更新菜单
			menu.POST("list", admin.ListMenu)     //更新菜单
		}
		// 系统相关设置
		system := v1.Group("system").Use(middleware.JwtMiddleware(), middleware.Auth(), middleware.ValidatorMiddleware())
		{
			system.POST("resource", admin.Resource) //获取当前服务器资源及占用 github.com/shirou/gopsutil

		}
	}

	//验证码
	e.GET("verify", service.Verify)
	//上传文件
	up := e.Group("upload")
	{
		up.POST("file", service.UploadFile)
		up.POST("files", service.UploadFiles)
	}

}

type RouterData struct {
	Handle string
	Path   string
	Action string
}

func GetRouterData(path string) (RouterData, error) {
	routerData := make(map[string]RouterData)
	routerData["/v1/user/login"] = RouterData{"admin.user.login", "/v1/user/login", "POST"}
	routerData["/v1/user/login"] = RouterData{"admin.user.login", "/v1/user/login", "POST"}
	routerData["/v1/user/login"] = RouterData{"admin.user.login", "/v1/user/login", "POST"}
	routerData["/v1/user/login"] = RouterData{"admin.user.login", "/v1/user/login", "POST"}
	routerData["/v1/user/login"] = RouterData{"admin.user.login", "/v1/user/login", "POST"}
	routerData["/v1/user/login"] = RouterData{"admin.user.login", "/v1/user/login", "POST"}
	routerData["/v1/user/login"] = RouterData{"admin.user.login", "/v1/user/login", "POST"}
	routerData["/v1/user/login"] = RouterData{"admin.user.login", "/v1/user/login", "POST"}
	routerData["/v1/user/login"] = RouterData{"admin.user.login", "/v1/user/login", "POST"}
	routerData["/v1/user/login"] = RouterData{"admin.user.login", "/v1/user/login", "POST"}
	routerData["/v1/user/login"] = RouterData{"admin.user.login", "/v1/user/login", "POST"}
	routerData["/v1/user/login"] = RouterData{"admin.user.login", "/v1/user/login", "POST"}
	routerData["/v1/user/login"] = RouterData{"admin.user.login", "/v1/user/login", "POST"}
	routerData["/v1/user/login"] = RouterData{"admin.user.login", "/v1/user/login", "POST"}
	routerData["/v1/user/login"] = RouterData{"admin.user.login", "/v1/user/login", "POST"}
	routerData["/v1/user/login"] = RouterData{"admin.user.login", "/v1/user/login", "POST"}
	routerData["/v1/user/login"] = RouterData{"admin.user.login", "/v1/user/login", "POST"}
	routerData["/v1/user/login"] = RouterData{"admin.user.login", "/v1/user/login", "POST"}
	routerData["/v1/user/login"] = RouterData{"admin.user.login", "/v1/user/login", "POST"}
	routerData["/v1/user/login"] = RouterData{"admin.user.login", "/v1/user/login", "POST"}
	v, ok := routerData[path]
	if !ok {
		return v, errors.New("not have this router")
	}
	return v, nil
}
