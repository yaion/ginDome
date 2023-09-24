package middleware

import (
	"github.com/gin-gonic/gin"
)

// LoggerMiddleware 记录请求日志
func LoggerMiddleware() gin.HandlerFunc {
	return func(context *gin.Context) {

		// 请求前业务逻辑...

		// 继续往下执行
		context.Next()
		// context.Abort() 停止往下执行
		// 请求后业务逻辑...
	}
}
