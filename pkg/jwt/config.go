package jwt

import "strings"

// ConfigHandler ..
type ConfigHandler interface {
	GetInt(key string) int
	GetString(key string) string
}

// Config ..
type Config struct {
	Issuer string
	Secret string
	TTL    int
}

// RawConfig ..
func RawConfig(confPrefix string, confHandler ConfigHandler) *Config {
	if strings.HasSuffix(confPrefix, ".") {
		confPrefix = confPrefix[:len(confPrefix)-1]
	}
	return &Config{
		Issuer: confHandler.GetString(confPrefix + ".name"),
		Secret: confHandler.GetString(confPrefix + ".secret"),
		TTL:    confHandler.GetInt(confPrefix + ".token_ttl"),
	}
}

var _config *Config

func (t *Config) Init() {
	_config = t
}
