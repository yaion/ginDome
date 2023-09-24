package middleware

import (
	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(context *gin.Context) {
		// 获取用户信息 todo
		//通过用户信息获取权限 todo
		//判断当前请求是否有权限 todo

		// 继续往下执行
		context.Next()
		// context.Abort() 停止往下执行
		// 请求后业务逻辑...
	}
}

