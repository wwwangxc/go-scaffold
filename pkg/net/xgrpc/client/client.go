package client

import (
	"context"
	"fmt"
	"time"

	"google.golang.org/grpc"
	"google.golang.org/grpc/resolver"
)

func newClient(conf *Config) (*grpc.ClientConn, error) {
	if conf.resolver != nil {
		resolver.Register(conf.resolver)
	}
	ctx := context.Background()
	if conf.DialTimeout > 0 {
		var cancel context.CancelFunc
		ctx, cancel = context.WithTimeout(ctx, time.Duration(conf.DialTimeout)*time.Second)
		defer cancel()
	}
	return grpc.DialContext(
		ctx,
		conf.Addr,
		grpc.WithDefaultServiceConfig(fmt.Sprintf(`{"LoadBalancingPolicy": "%s"}`, conf.BalancerName)),
		grpc.WithInsecure(),
	)
}
