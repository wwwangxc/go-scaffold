package etcd

import (
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
	Endpoints      []string
	DialTimeout    int64
	HeartbeatCycle int64
	LeaseTTL       int64
}

// RawRegisterConfig ..
func RawRegisterConfig(confPrefix string, confHandler ConfigHandler) *RegisterConfig {
	if strings.HasSuffix(confPrefix, ".") {
		confPrefix = confPrefix[:len(confPrefix)-1]
	}
	return &RegisterConfig{
		Endpoints:      confHandler.GetStringSlice(confPrefix + ".endpoints"),
		DialTimeout:    confHandler.GetInt64(confPrefix + ".dial_timeout"),
		HeartbeatCycle: confHandler.GetInt64(confPrefix + ".heartbeat_cycle"),
		LeaseTTL:       confHandler.GetInt64(confPrefix + ".lease_ttl"),
	}
}

// Build ..
func (t *RegisterConfig) Build() *Register {
	return newRegister(t)
}
