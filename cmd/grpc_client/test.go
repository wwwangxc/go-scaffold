package main

import (
	"context"
	"fmt"
	"go-scaffold/internal/grpc/pb"
	"go-scaffold/pkg/config"
	"go-scaffold/pkg/etcd/resolver"
	grpcClient "go-scaffold/pkg/net/xgrpc/client"
)

func main() {
	config.Init() // 加载配置文件
	resolver, err := resolver.RawConfig("grpc.resolver.etcd", config.GetHandler()).Build()
	if err != nil {
		panic(err)
	}
	conn, err := grpcClient.RawConfig("grpc.ping", config.GetHandler()).WithResolver(resolver).Build()
	if err != nil {
		panic(err)
	}
	defer conn.Close()
	cli := pb.NewPingClient(conn)
	ret, err := cli.Ping(context.Background(), &pb.PingRequest{})
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	fmt.Println(ret.GetMessage())
}
