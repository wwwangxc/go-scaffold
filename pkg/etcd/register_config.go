package etcd

import (
	"go-scaffold/pkg/conf"
	"strings"
)

// ConfigHandler ..
type ConfigHandler interface {
	GetInt64(key string) int64
	GetStringSlice(key string) []string
	GetString(key string) string
}

// RegisterConfig ..
type RegisterConfig struct {
	Endpoints []string
	TTL       int64
}

// RawRegisterConfig ..
func RawRegisterConfig(confPrefix string, confHandler ConfigHandler) *RegisterConfig {
	if strings.HasSuffix(confPrefix, ".") {
		confPrefix = confPrefix[:len(confPrefix)-1]
	}
	return &RegisterConfig{
		Endpoints: confHandler.GetStringSlice(confPrefix + ".endpoints"),
		TTL:       conf.GetInt64(confPrefix + ".ttl"),
	}
}

// Build ..
func (t *RegisterConfig) Build() *Register {
	return newRegister(t)
}
