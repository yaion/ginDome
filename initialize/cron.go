package initialize

import (
	"fmt"
	"ginDome/global"
	"github.com/robfig/cron/v3"
)

// 定时任务初始化 todo
func initCron(){
	//判断配置文件定时任务是否开启
	if !global.GvaConfig.App.Cron{
		return
	}
	//启动定时任务
	secondParser := cron.NewParser(cron.Second | cron.Minute |
		cron.Hour | cron.Dom | cron.Month | cron.DowOptional | cron.Descriptor)
	c:= cron.New(cron.WithParser(secondParser), cron.WithChain())
	funcId, err := c.AddFunc("0 */1 * * * ?", func() {
		fmt.Println("cron string")
	})
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(funcId)


	// 开始
	c.Start()
	select {}

	//获取数据库表中定时任务


}
