package etcd

import (
	"strings"
	"time"

	"google.golang.org/grpc/resolver"
)

// ResolverConfig ..
type ResolverConfig struct {
	Endpoints   []string
	DialTimeout time.Duration
	Scheme      string
}

// RawResolverConfig ..
func RawResolverConfig(confPrefix string, confHandler ConfigHandler) *ResolverConfig {
	if strings.HasSuffix(confPrefix, ".") {
		confPrefix = confPrefix[:len(confPrefix)-1]
	}
	return &ResolverConfig{
		Endpoints:   confHandler.GetStringSlice(confPrefix + ".endpoints"),
		DialTimeout: confHandler.GetDuration(confPrefix + ".dial_timeout"),
		Scheme:      confHandler.GetString(confPrefix + ".scheme"),
	}
}

// Build ..
func (t *ResolverConfig) Build() (resolver.Builder, error) {
	return newResolver(t)
}
