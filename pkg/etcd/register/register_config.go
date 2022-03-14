package register

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

// Config ..
type Config struct {
	Endpoints      []string
	DialTimeout    time.Duration
	HeartbeatCycle time.Duration
	LeaseTTL       time.Duration
}

// RawConfig ..
func RawConfig(confPrefix string, confHandler ConfigHandler) *Config {
	confPrefix = strings.TrimSuffix(confPrefix, ".")

	return &Config{
		Endpoints:      confHandler.GetStringSlice(confPrefix + ".endpoints"),
		DialTimeout:    confHandler.GetDuration(confPrefix + ".dial_timeout"),
		HeartbeatCycle: confHandler.GetDuration(confPrefix + ".heartbeat_cycle"),
		LeaseTTL:       confHandler.GetDuration(confPrefix + ".lease_ttl"),
	}
}

// Build ..
func (t *Config) Build() *Register {
	return newRegister(t)
}
