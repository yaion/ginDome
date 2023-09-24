package initialize

import (
	"flag"
	"fmt"
	"ginDome/global"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// InitConfig 初始化viper配置
func InitConfig(){
	var configFile string
	//读取配置文件优先级: 命令行 > 默认值
	flag.StringVar(&configFile,"c",global.ConfigFile,"config")
	if len(configFile) == 0 {
		panic(any("config undefined"))
	}
	//读取config
	v:=viper.New()
	v.SetConfigFile(configFile)
	if err:=v.ReadInConfig();err!=nil{
		panic(any(fmt.Errorf("conf analyze err:%s",err)))
	}

	//动态检测配置文件
	v.WatchConfig()
	v.OnConfigChange(func(in fsnotify.Event) {
		//todo 打日志 config change
		if err:=v.Unmarshal(&global.GvaConfig);err!=nil{
			panic(any(fmt.Errorf("config overload err:%s",err)))
		}
	})
	if err:=v.Unmarshal(&global.GvaConfig);err!=nil{
		panic(any(fmt.Errorf("config overload err:%s",err)))
	}
	global.GvaConfig.App.ConfigFile = configFile
}
