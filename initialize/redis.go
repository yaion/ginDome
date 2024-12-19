package initialize

import (
	"context"
	"fmt"
	"ginDome/global"
	"github.com/redis/go-redis/v9"
)

// todo redis 启动项
func InitRedis() {
	addr := fmt.Sprintf("%s:%s", global.GvaConfig.Redis.Host, global.GvaConfig.Redis.Port)
	redisClient := redis.NewClient(&redis.Options{
		Addr:     addr,
		Password: global.GvaConfig.Redis.Password,  // no password set
		DB:       global.GvaConfig.Redis.DefaultDB, // use default DB
	})
	//使用超时上下文，验证redis
	timeoutCtx, canceFunc := context.WithTimeout(context.Background(), global.GvaConfig.Redis.DialTimeOut)
	defer canceFunc()
	_, err := redisClient.Ping(timeoutCtx).Result()
	if err != nil {
		panic(any(fmt.Sprintf("redis 初始化失败：%s", err.Error())))
	}
	global.GvaRedisClient = redisClient
}
