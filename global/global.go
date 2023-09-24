package global

import (
	"ginDome/config"
	"github.com/go-playground/validator/v10"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)
// 常量
const (
	ConfigFile = "./config/app.yaml" //配置文件
)

// 变量
var (
	GvaConfig config.ServerConfig //全局配置
	GvaValidate = validator.New() //验证器
	GvaMysqlClient *gorm.DB
	GvaLogger *zap.Logger
	GvaRedisClient *redis.Client
	GvaParams = "params" //
	GvaUserData = "userData"
)

//错误处理码 todo



