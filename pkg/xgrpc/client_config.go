package xgrpc

import (
	"strings"

	"google.golang.org/grpc"

	"google.golang.org/grpc/resolver"
)

// RawClientConfig
func RawClientConfig(confPrefix string, confHandler ConfigHandler) *ClientConfig {
	if strings.HasSuffix(confPrefix, ".") {
		confPrefix = confPrefix[:len(confPrefix)-1]
	}
	return &ClientConfig{
		Addr:         confHandler.GetString(confPrefix + ".address"),
		BalancerName: confHandler.GetString(confPrefix + ".balancer_name"),
		Timeout:      confHandler.GetInt(confPrefix + ".timeout"),
	}
}

// ClientConfig ..
type ClientConfig struct {
	Addr         string
	BalancerName string
	Timeout      int

	resolver resolver.Builder
}

func (t *ClientConfig) WithResolver(resolver resolver.Builder) *ClientConfig {
	t.resolver = resolver
	return t
}

func (t *ClientConfig) Build() (*grpc.ClientConn, error) {
	return newClient(t)
}
