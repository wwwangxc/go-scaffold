package xgrpc

import (
	"strings"
	"time"

	"google.golang.org/grpc"

	"google.golang.org/grpc/resolver"
)

// ClientConfig ..
type ClientConfig struct {
	Addr         string
	BalancerName string
	DialTimeout  time.Duration

	resolver resolver.Builder
}

// RawClientConfig ..
func RawClientConfig(confPrefix string, confHandler ConfigHandler) *ClientConfig {
	if strings.HasSuffix(confPrefix, ".") {
		confPrefix = confPrefix[:len(confPrefix)-1]
	}
	return &ClientConfig{
		Addr:         confHandler.GetString(confPrefix + ".address"),
		BalancerName: confHandler.GetString(confPrefix + ".balancer_name"),
		DialTimeout:  confHandler.GetDuration(confPrefix + ".dial_timeout"),
	}
}

// WithResolver ..
func (t *ClientConfig) WithResolver(resolver resolver.Builder) *ClientConfig {
	t.resolver = resolver
	return t
}

// Build ..
func (t *ClientConfig) Build() (*grpc.ClientConn, error) {
	return newClient(t)
}
