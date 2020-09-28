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

// Serve ..
func Serve() {
	// 加载配置文件
	conf.Init()

	// 加载日志
	log.RawConfig("app.log", conf.GetHandler()).Init()
	defer log.Sync()

	// 加载mysql
	xgorm.RawConfig("app.mysql.db1", conf.GetHandler()).Append(constant.MySQLStoreNameDB1)
	xgorm.RawConfig("app.mysql.db2", conf.GetHandler()).Append(constant.MySQLStoreNameDB2)
	defer xgorm.CloseAll()

	// 加载redis
	xredis.RawConfig("app.redis.0", conf.GetHandler()).Append(constant.RedisStoreNameDB0)
	xredis.RawConfig("app.redis.1", conf.GetHandler()).Append(constant.RedisStoreNameDB1)
	defer xredis.CloseAll()

	// http服务
	xgin.RawConfig("app.http", conf.GetHandler()).
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
