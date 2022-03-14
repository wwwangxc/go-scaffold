package sqlite

import (
	"strings"
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
func RawConfig(confPrefix string, confHandler ConfigHandler) *Config {
	confPrefix = strings.TrimSuffix(confPrefix, ".")

	return &Config{
		DSN: confHandler.GetString(confPrefix + ".dsn"),
	}
}

// Build ..
func (t *Config) Build() *DB {
	return newDB(t)
}

// Append ..
func (t *Config) Append(storeName string) *DB {
	db := newDB(t)
	store.append(storeName, db)
	return db
}
