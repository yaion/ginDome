package initialize

import (
	"fmt"
	"ginDome/global"
	"testing"
)

func TestInitConfig(t *testing.T){
	// 目录问题导致test不通过 todo
	InitConfig()
	fmt.Println(global.GvaConfig)
}
