package xmysql

import (
	"strings"

	"github.com/jmoiron/sqlx"
)

// ConfigHandler ..
type ConfigHandler interface {
	GetString(key string) string
}

// Config ..
type Config struct {
	DSN string
}

// RawConfig ..
func RawConfig(confPath string, confHandler ConfigHandler) *Config {
	if strings.HasSuffix(confPath, ",") {
		confPath = confPath[:len(confPath)-1]
	}
	return &Config{
		DSN: confHandler.GetString(confPath + ".dsn"),
	}
}

// Init ..
func (t *Config) Init() {
	initialize(t)
}

// Build ..
func (t *Config) Build() *sqlx.DB {
	return newClient(t)
}
