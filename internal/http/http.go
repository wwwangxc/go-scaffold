package http

import (
	"go-scaffold/internal/constant"
	"go-scaffold/internal/http/handler"
	"go-scaffold/internal/http/middleware"
	"go-scaffold/pkg/conf"
	"go-scaffold/pkg/log"
	"go-scaffold/pkg/xgin"
	"go-scaffold/pkg/xgorm"
	"go-scaffold/pkg/xredis"
)

func Run() {
	conf.Init()                                                                            // 加载配置文件
	log.RawConfig("app.log", conf.GetHandler()).Init()                                     // 加载日志
	defer log.Sync()                                                                       // 日志落盘
	xgorm.RawConfig("app.mysql.db1", conf.GetHandler()).Append(constant.MySQLStoreNameDB1) // 加载gorm
	xgorm.RawConfig("app.mysql.db2", conf.GetHandler()).Append(constant.MySQLStoreNameDB2) // 加载gorm
	defer xgorm.CloseAll()
	xredis.RawConfig("app.redis.0", conf.GetHandler()).Append(constant.RedisStoreNameDB0) // 加载Redis 0
	xredis.RawConfig("app.redis.1", conf.GetHandler()).Append(constant.RedisStoreNameDB1) // 加载Redis 1
	defer xredis.CloseAll()
	xgin.RawConfig("app.http", conf.GetHandler()).
		WithMiddlewares(
			middleware.Logger,
			middleware.Recovery,
			middleware.Authentication,
			middleware.Swagger,
		).
		WithRoutes(
			handler.RoutePing,
			handler.RouteAuthentication,
		).Build().ListenAndServe()
}
