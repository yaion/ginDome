package initialize

import (
	"fmt"
)

func init(){
	//初始化config
	InitConfig()
	//初始化logger
	InitLogger()
	//初始化gorm
	InitGorm()
	//初始化redis
	InitRedis()

	//初始化定时任务 有问题
	initCron()

}

func New(){
	fmt.Println("system init ....")
}

