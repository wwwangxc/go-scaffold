package grpc

import (
	"go-scaffold/internal/constant"
	"go-scaffold/internal/grpc/handler"
	"go-scaffold/internal/grpc/pb"
	"go-scaffold/pkg/cache/xredis"
	"go-scaffold/pkg/config"
	"go-scaffold/pkg/database/xgorm"
	"go-scaffold/pkg/etcd/register"
	"go-scaffold/pkg/log"
	grpcServer "go-scaffold/pkg/net/xgrpc/server"

	"google.golang.org/grpc"
)

// Serve ..
func Serve() {
	// config
	config.Init()

	// log
	log.RawConfig("app.log", config.GetHandler()).Init()
	defer log.Sync()

	// mysql1
	xgorm.RawConfig("app.mysql.db1", config.GetHandler()).Append(constant.MySQLStoreNameDB1)
	defer xgorm.Close(constant.MySQLStoreNameDB1)

	// mysql2
	xgorm.RawConfig("app.mysql.db2", config.GetHandler()).Append(constant.MySQLStoreNameDB2)
	defer xgorm.Close(constant.MySQLStoreNameDB2)

	// redis0
	xredis.RawConfig("app.redis.0", config.GetHandler()).Append(constant.RedisStoreNameDB0)
	defer xredis.Close(constant.RedisStoreNameDB0)

	// redis1
	xredis.RawConfig("app.redis.1", config.GetHandler()).Append(constant.RedisStoreNameDB1)
	defer xredis.Close(constant.RedisStoreNameDB1)

	// grpc server
	register := register.RawConfig("app.grpc.register.etcd", config.GetHandler()).Build()
	srv := grpcServer.RawConfig("app.grpc", config.GetHandler()).
		WithRegister(register).
		WithService(func(server *grpc.Server) {
			pb.RegisterPingServer(server, &handler.Ping{})
		}).
		Build()
	defer srv.Close()
	srv.Serve()
}
