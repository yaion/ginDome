package utils

import (
	"errors"
	"ginDome/global"
	"github.com/golang-jwt/jwt/v5"
	"time"
)

// VerifyJwt 验证token
func VerifyJwt(tokenString string)(uint64,error){
	var uid uint64
	secret := []byte(global.GvaConfig.Jwt.Secret)
	parsedToken,err:=jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secret,nil
	})
	if err !=nil{
		//todo 记录日志
		return 0,err
	}
	if claims,ok:=parsedToken.Claims.(jwt.MapClaims);ok&&parsedToken.Valid{
		userId:=claims["user_id"].(float64)
		uid=uint64(userId)
	}else {
		err := errors.New("token is not valid")
		return 0,err
	}
	return uid,nil
}

// CreateJwt 创建jwt
func CreateJwt(uid uint64)(tokens string,err error){
	// 设置声明（claims）
	claims := jwt.MapClaims{
		"user_id":uid,
		"exp":time.Now().Add(time.Second * time.Duration(global.GvaConfig.Jwt.Expire)).Unix(),
		"issuer":global.GvaConfig.Jwt.Issuer,
	}
	// 创建一个新的 JWT
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	secret := []byte(global.GvaConfig.Jwt.Secret)
	return token.SignedString(secret)
}

