package utils

import (
	"crypto/md5"
	"fmt"
	"ginDome/global"
	"testing"
	"time"
)

func TestCreateJwt(t *testing.T) {
	global.GvaConfig.Jwt.Secret="goDome"
	global.GvaConfig.Jwt.Expire=3600
	var uid uint64
	uid=2
	token,err:=CreateJwt(uid)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(token)
	time.Sleep(2*time.Second)
	uid,err=VerifyJwt(token)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(uid)

}
func TestDirExist(t *testing.T) {
	path :="../logs"
	fmt.Println(DirExist(path))
}

func TestCreateDir(t *testing.T) {
	path := "../testLogs"
	fmt.Println(CreateDir(path))
}

func TestPassworld(t *testing.T) {
	passworld:=[]byte("123456")
	has:=md5.Sum(passworld)
	md5str := fmt.Sprintf("%x", has)
	fmt.Println(md5str)
}