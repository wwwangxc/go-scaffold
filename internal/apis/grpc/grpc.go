package grpc

import (
	"go-scaffold/internal/apis/grpc/handler"
	pingPB "go-scaffold/internal/apis/grpc/proto/ping"
	"go-scaffold/internal/constant"
	"go-scaffold/pkg/cache/xredis"
	"go-scaffold/pkg/config"
	"go-scaffold/pkg/database/xgorm"
	"go-scaffold/pkg/etcd/register"
	"go-scaffold/pkg/log"
	grpcServer "go-scaffold/pkg/net/xgrpc/server"

	"google.golang.org/grpc"
)

// Serve ..
func Serve(port int) {
	// config
	config.Init()

	// log
	log.RawConfig("app.log", config.GetHandler()).Init()
	defer log.Sync()

	// mysql master
	xgorm.RawConfig("app.mysql.master", config.GetHandler()).Append(constant.MySQLStoreNameMaster)

	// mysql slave
	xgorm.RawConfig("app.mysql.slave", config.GetHandler()).Append(constant.MySQLStoreNameSlave)

	// redis0
	xredis.RawConfig("app.redis.0", config.GetHandler()).Append(constant.RedisStoreNameDB0)
	defer xredis.Close(constant.RedisStoreNameDB0)

	// redis1
	xredis.RawConfig("app.redis.1", config.GetHandler()).Append(constant.RedisStoreNameDB1)
	defer xredis.Close(constant.RedisStoreNameDB1)

	// grpc server
	register := register.RawConfig("app.grpc.register.etcd", config.GetHandler()).Build()
	srv := grpcServer.RawConfig("app.grpc", config.GetHandler()).
		WithPort(port).
		WithRegister(register).
		WithService(func(server *grpc.Server) {
			pingPB.RegisterPingServer(server, &handler.Ping{})
		}).
		Build()
	srv.Serve()
}
