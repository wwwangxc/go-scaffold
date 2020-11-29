package client

import (
	"strings"
	"time"

	"google.golang.org/grpc"

	"google.golang.org/grpc/resolver"
)

// ConfigHandler ..
type ConfigHandler interface {
	GetString(key string) string
	GetInt(key string) int
	GetDuration(key string) time.Duration
}

// Config ..
type Config struct {
	Addr         string
	BalancerName string
	DialTimeout  time.Duration

	resolver resolver.Builder
}

// RawConfig ..
func RawConfig(confPrefix string, confHandler ConfigHandler) *Config {
	if strings.HasSuffix(confPrefix, ".") {
		confPrefix = confPrefix[:len(confPrefix)-1]
	}
	return &Config{
		Addr:         confHandler.GetString(confPrefix + ".address"),
		BalancerName: confHandler.GetString(confPrefix + ".balancer_name"),
		DialTimeout:  confHandler.GetDuration(confPrefix + ".dial_timeout"),
	}
}

// WithResolver ..
func (t *Config) WithResolver(resolver resolver.Builder) *Config {
	t.resolver = resolver
	return t
}

// Build ..
func (t *Config) Build() (*grpc.ClientConn, error) {
	return newClient(t)
}
