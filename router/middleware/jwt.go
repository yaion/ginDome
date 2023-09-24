package middleware

import (
	"context"
	"ginDome/global"
	"ginDome/model/entity"
	"ginDome/utils"
	"github.com/gin-gonic/gin"
)


func JwtMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// 请求前业务逻辑...
		//获取token,写在header里面
		token:=ctx.Request.Header.Get("token")
		//验证token
		uid,err:=utils.VerifyJwt(token)
		if err!=nil{
			//放回登录失败
			ctx.JSON(500,gin.H{"msg":"please login!"})
			ctx.Abort() //停止往下执行
			return
		}
		user := entity.User{
			UserId: uid,
		}
		err=setUserData(ctx,&user)
		if err!=nil{
			//todo 记录日志
			ctx.JSON(500,gin.H{"msg":"system error please try again!"})
			ctx.Abort() //停止往下执行
			return
		}
		//通过token获取用户信息，设置到上下文
		// 继续往下执行
		ctx.Next()
		// 请求后业务逻辑...
	}
}

//ParamsContext


//获取token信息，设置到上下文
func setUserData(ctx *gin.Context,user *entity.User)error{
	if err:=user.GetUserOne();err!=nil{
		return err
	}
	c:=context.WithValue(ctx.Request.Context(),global.GvaUserData,user)
	ctx.Request=ctx.Request.WithContext(c)
	return nil
}

