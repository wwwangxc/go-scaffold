package etcd

import (
	"strings"
	"time"
)

// ConfigHandler ..
type ConfigHandler interface {
	GetStringSlice(key string) []string
	GetString(key string) string
	GetDuration(key string) time.Duration
}

// RegisterConfig ..
type RegisterConfig struct {
	Endpoints      []string
	DialTimeout    time.Duration
	HeartbeatCycle time.Duration
	LeaseTTL       time.Duration
}

// RawRegisterConfig ..
func RawRegisterConfig(confPrefix string, confHandler ConfigHandler) *RegisterConfig {
	if strings.HasSuffix(confPrefix, ".") {
		confPrefix = confPrefix[:len(confPrefix)-1]
	}
	return &RegisterConfig{
		Endpoints:      confHandler.GetStringSlice(confPrefix + ".endpoints"),
		DialTimeout:    confHandler.GetDuration(confPrefix + ".dial_timeout"),
		HeartbeatCycle: confHandler.GetDuration(confPrefix + ".heartbeat_cycle"),
		LeaseTTL:       confHandler.GetDuration(confPrefix + ".lease_ttl"),
	}
}

// Build ..
func (t *RegisterConfig) Build() *Register {
	return newRegister(t)
}
