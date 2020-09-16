// grpc 服务入口

package main

import (
	"go-scaffold/internal/constant"
	"go-scaffold/internal/grpc"
	"go-scaffold/pkg/conf"
	"go-scaffold/pkg/log"
	"go-scaffold/pkg/xgorm"
	"go-scaffold/pkg/xredis"
)

func main() {
	conf.Init()                                                                            // 加载配置文件
	log.RawConfig("app.log", conf.GetHandler()).Init()                                     // 加载日志
	defer log.Sync()                                                                       // 日志落盘
	xgorm.RawConfig("app.mysql.db1", conf.GetHandler()).Append(constant.MySQLStoreNameDB1) // 加载Mysql 1
	defer xgorm.Close(constant.MySQLStoreNameDB1)
	xgorm.RawConfig("app.mysql.db2", conf.GetHandler()).Append(constant.MySQLStoreNameDB2) // 加载Mysql 2
	defer xgorm.Close(constant.MySQLStoreNameDB2)
	xredis.RawConfig("app.redis.0", conf.GetHandler()).Append(constant.RedisStoreNameDB0) // 加载Redis 0
	defer xredis.Close(constant.RedisStoreNameDB0)
	xredis.RawConfig("app.redis.1", conf.GetHandler()).Append(constant.RedisStoreNameDB1) // 加载Redis 1
	defer xredis.Close(constant.RedisStoreNameDB1)
	grpc.Serve()
}
