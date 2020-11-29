package resolver

import (
	"strings"
	"time"

	"google.golang.org/grpc/resolver"
)

// ConfigHandler ..
type ConfigHandler interface {
	GetStringSlice(key string) []string
	GetString(key string) string
	GetDuration(key string) time.Duration
}

// Config ..
type Config struct {
	Endpoints   []string
	DialTimeout time.Duration
	Scheme      string
}

// RawConfig ..
func RawConfig(confPrefix string, confHandler ConfigHandler) *Config {
	if strings.HasSuffix(confPrefix, ".") {
		confPrefix = confPrefix[:len(confPrefix)-1]
	}
	return &Config{
		Endpoints:   confHandler.GetStringSlice(confPrefix + ".endpoints"),
		DialTimeout: confHandler.GetDuration(confPrefix + ".dial_timeout"),
		Scheme:      confHandler.GetString(confPrefix + ".scheme"),
	}
}

// Build ..
func (t *Config) Build() (resolver.Builder, error) {
	return newResolver(t)
}
