package jwt

import (
	"strings"
	"time"
)

// ConfigHandler ..
type ConfigHandler interface {
	GetDuration(key string) time.Duration
	GetString(key string) string
}

// Config ..
type Config struct {
	Issuer string
	Secret string
	TTL    time.Duration
}

// RawConfig ..
func RawConfig(confPrefix string, confHandler ConfigHandler) *Config {
	confPrefix = strings.TrimSuffix(confPrefix, ".")

	return &Config{
		Issuer: confHandler.GetString(confPrefix + ".name"),
		Secret: confHandler.GetString(confPrefix + ".secret"),
		TTL:    confHandler.GetDuration(confPrefix + ".token_ttl"),
	}
}

var config *Config

// Init ..
func (t *Config) Init() {
	config = t
}
