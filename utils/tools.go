package utils

import (
	"crypto/md5"
	"fmt"
	"os"
)

func DirExist(path string) bool {
	if _, err := os.Stat(path); err != nil {
		return false
	}
	return true
}

func CreateDir(path string) bool {
	err := os.MkdirAll(path, os.ModePerm)
	if err != nil {
		return false
	}
	return true
}

func Md5str(str string) string {
	result := []byte(str)
	has := md5.Sum(result)
	md5str := fmt.Sprintf("%x", has)
	return md5str
}

type RouterData struct {
	Handle string
	Path   string
	Action string
}

func GetRouterData(path string) map[string]RouterData {
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
	return routerData
}
