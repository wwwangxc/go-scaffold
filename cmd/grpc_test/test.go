package main

import (
	"context"
	"fmt"
	"go-scaffold/internal/grpc/pb"
	"go-scaffold/pkg/conf"
	"go-scaffold/pkg/etcd"
	"go-scaffold/pkg/xgrpc"
)

func main() {
	conf.Init() // 加载配置文件
	resolver, err := etcd.RawResolverConfig("grpc.resolver.etcd", conf.GetHandler()).Build()
	if err != nil {
		panic(err)
	}
	conn, err := xgrpc.RawClientConfig("grpc.ping", conf.GetHandler()).WithResolver(resolver).Build()
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
