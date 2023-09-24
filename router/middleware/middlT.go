package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"time"
)

func MiddleName() gin.HandlerFunc {
	return func(context *gin.Context) {
		t:=time.Now().Unix()
		// 请求前业务逻辑...
		time.Sleep(2*time.Second)
		fmt.Println("这是中间件前")
		// 继续往下执行
		context.Next()
		// context.Abort() 停止往下执行
		// 请求后业务逻辑...
		fmt.Println("这是中间件后")
		end := time.Now().Unix()
		fmt.Printf("接口耗时:%v 秒 \n", end - t)
	}
}

