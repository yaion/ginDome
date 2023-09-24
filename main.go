package main

import (
	"ginDome/global"
	"ginDome/initialize"
	"ginDome/router"
	"github.com/gin-gonic/gin"
)




func main()  {
	initialize.New()
	//创建一个默认路由引擎
	gin.SetMode(global.GvaConfig.App.Env)
	engine := gin.Default()
	//注册路由，并设置一个匿名的handlers
	router.InitRouter(engine)
	//启动服务，并监听8080端口
	_=engine.Run(global.GvaConfig.App.Addr)

}



