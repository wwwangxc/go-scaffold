package etcd

import (
	"strings"

	"google.golang.org/grpc/resolver"
)

// RawResolverConfig ..
func RawResolverConfig(confPrefix string, confHandler ConfigHandler) *ResolverConfig {
	if strings.HasSuffix(confPrefix, ".") {
		confPrefix = confPrefix[:len(confPrefix)-1]
	}
	return &ResolverConfig{
		Endpoints: confHandler.GetStringSlice(confPrefix + ".endpoints"),
		TTL:       confHandler.GetInt64(confPrefix + ".ttl"),
		Scheme:    confHandler.GetString(confPrefix + ".scheme"),
	}
}

// ResolverConfig ..
type ResolverConfig struct {
	Endpoints []string
	TTL       int64
	Scheme    string
}

func (t *ResolverConfig) Build() (resolver.Builder, error) {
	return newResolver(t)
}
