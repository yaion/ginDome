package middleware

import (
	"fmt"
	"ginDome/global"
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 获取用户信息 todo
		//通过用户信息获取权限 todo
		//判断当前请求是否有权限 todo
		//path:=ctx.Request.URL.Path
		//验证当前用户的权限
		getData, err := global.GvaRedisClient.Get(ctx.Request.Context(), "name").Result()
		if err != nil {

		}
		fmt.Println(getData)
		//permissions := []
		//json.Unmarshal([]byte(getData))
		// 继续往下执行
		ctx.Next()
		// context.Abort() 停止往下执行
		// 请求后业务逻辑...
	}
}
