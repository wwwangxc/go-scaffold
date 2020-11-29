package http

import (
	"go-scaffold/internal/constant"
	"go-scaffold/internal/http/handler"
	"go-scaffold/internal/http/middleware"
	"go-scaffold/pkg/cache/xredis"
	"go-scaffold/pkg/config"
	"go-scaffold/pkg/database/xgorm"
	"go-scaffold/pkg/log"
	"go-scaffold/pkg/net/xgin"
)

// Serve ..
func Serve() {
	// config
	config.Init()

	// log
	log.RawConfig("app.log", config.GetHandler()).Init()
	defer log.Sync()

	// mysql
	xgorm.RawConfig("app.mysql.master", config.GetHandler()).Append(constant.MySQLStoreNameMaster)
	xgorm.RawConfig("app.mysql.slave", config.GetHandler()).Append(constant.MySQLStoreNameSlave)
	defer xgorm.CloseAll()

	// redis
	xredis.RawConfig("app.redis.0", config.GetHandler()).Append(constant.RedisStoreNameDB0)
	xredis.RawConfig("app.redis.1", config.GetHandler()).Append(constant.RedisStoreNameDB1)
	defer xredis.CloseAll()

	// http server
	xgin.RawConfig("app.http", config.GetHandler()).
		WithMiddlewares(
			middleware.Logger,         // 日志
			middleware.Recovery,       // 错误恢复
			middleware.Authentication, // 鉴权
			middleware.Swagger,        // swagger
		).
		WithRoutes(
			handler.RoutePing,           // ping
			handler.RouteAuthentication, // 鉴权接口
		).Build().ListenAndServe()
}
