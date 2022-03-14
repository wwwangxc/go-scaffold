package xgorm

import (
	"strings"
	"time"

	"gorm.io/gorm"
)

// ConfigHandler ..
type ConfigHandler interface {
	GetString(key string) string
	GetInt(key string) int
	GetDuration(key string) time.Duration
}

// Config ..
type Config struct {
	DSN             string
	MaxIdleConns    int           // 最大空闲连接数
	MaxOpenConns    int           // 最大活动连接数
	ConnMaxLifetime time.Duration // 连接的最大存活时间
}

// RawConfig ..
func RawConfig(confPrefix string, confHandler ConfigHandler) *Config {
	confPrefix = strings.TrimSuffix(confPrefix, ".")

	return &Config{
		DSN:             confHandler.GetString(confPrefix + ".dsn"),
		MaxIdleConns:    confHandler.GetInt(confPrefix + ".max_idle_conns"),
		MaxOpenConns:    confHandler.GetInt(confPrefix + ".max_open_conns"),
		ConnMaxLifetime: confHandler.GetDuration(confPrefix + ".conn_max_lifetime"),
	}
}

// Build ..
func (t *Config) Build() *gorm.DB {
	return newDB(t)
}

// Append append db to dbStore and return.
func (t *Config) Append(storeName string) *gorm.DB {
	db := newDB(t)
	store.append(storeName, db)
	return db
}
