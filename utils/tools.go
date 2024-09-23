package utils

import (
	"bytes"
	"crypto/md5"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/dchest/captcha"
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

func CreateVerifyImg(height, width int) (map[string]string, error) {
	verify := make(map[string]string)
	verify["id"] = captcha.New()
	var buffer bytes.Buffer
	err := captcha.WriteImage(&buffer, verify["id"], width, height)
	if err != nil {
		return verify, err
	}
	verify["img"] = base64.StdEncoding.EncodeToString(buffer.Bytes())
	return verify, nil
}

func Verify(id, digits string) bool {
	return captcha.VerifyString(id, digits)
}
