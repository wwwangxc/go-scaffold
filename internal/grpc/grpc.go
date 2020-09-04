package grpc

import (
	"go-scaffold/internal/grpc/handler"
	"go-scaffold/internal/grpc/pb"
	"go-scaffold/pkg/conf"
	"go-scaffold/pkg/etcd"
	"go-scaffold/pkg/xgrpc"

	"google.golang.org/grpc"
)

// Serve ..
func Serve() {
	grpcConf := xgrpc.RawServerConfig("app.grpc", conf.GetHandler())
	grpcConf.Setup(func(server *grpc.Server) {
		pb.RegisterPingServer(server, &handler.Ping{})
	})
	grpcConf.WithRegister(
		etcd.RawRegisterConfig("app.grpc.register.etcd", conf.GetHandler()).Build())
	srv := grpcConf.Build()
	defer srv.Close()
	srv.Serve()
}
