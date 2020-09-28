package grpc

import (
	"go-scaffold/internal/constant"
	"go-scaffold/internal/grpc/handler"
	"go-scaffold/internal/grpc/pb"
	"go-scaffold/pkg/conf"
	"go-scaffold/pkg/etcd"
	"go-scaffold/pkg/log"
	"go-scaffold/pkg/xgorm"
	"go-scaffold/pkg/xgrpc"
	"go-scaffold/pkg/xredis"

	"google.golang.org/grpc"
)

// Serve ..
func Serve() {
	// 加载配置文件
	conf.Init()

	// 加载日志
	log.RawConfig("app.log", conf.GetHandler()).Init()
	defer log.Sync()

	// 加载mysql 1
	xgorm.RawConfig("app.mysql.db1", conf.GetHandler()).Append(constant.MySQLStoreNameDB1)
	defer xgorm.Close(constant.MySQLStoreNameDB1)

	// 加载mysql2
	xgorm.RawConfig("app.mysql.db2", conf.GetHandler()).Append(constant.MySQLStoreNameDB2)
	defer xgorm.Close(constant.MySQLStoreNameDB2)

	// 加载redis0
	xredis.RawConfig("app.redis.0", conf.GetHandler()).Append(constant.RedisStoreNameDB0)
	defer xredis.Close(constant.RedisStoreNameDB0)

	// 加载redis1
	xredis.RawConfig("app.redis.1", conf.GetHandler()).Append(constant.RedisStoreNameDB1)
	defer xredis.Close(constant.RedisStoreNameDB1)

	// grpc服务
	srv := xgrpc.RawServerConfig("app.grpc", conf.GetHandler()).
		WithRegister(
			etcd.RawRegisterConfig("app.grpc.register.etcd", conf.GetHandler()).Build()).
		WithService(func(server *grpc.Server) {
			pb.RegisterPingServer(server, &handler.Ping{})
		}).Build()
	defer srv.Close()
	srv.Serve()
}
