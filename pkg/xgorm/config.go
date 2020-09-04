package xgorm

import (
	"strings"
	"sync"
	"time"

	"github.com/jinzhu/gorm"
)

// ConfigHandler ..
type ConfigHandler interface {
	GetInt(key string) int
	GetInt32(key string) int32
	GetInt64(key string) int64
	GetString(key string) string
	GetBool(key string) bool
	GetTime(key string) time.Time
}

// RawConfig ..
func RawConfig(confPrefix string, confHandler ConfigHandler) *Config {
	if strings.HasSuffix(confPrefix, ".") {
		confPrefix = confPrefix[:len(confPrefix)-1]
	}
	return &Config{
		DSN: confHandler.GetString(confPrefix + ".dsn"),
	}
}

var initOnce sync.Once

// Config ..
type Config struct {
	DSN string
}

// Init ..
func (t *Config) Init() {
	initOnce.Do(func() {
		initialize(t)
	})
}

// Build ..
func (t *Config) Build() *gorm.DB {
	return newClient(t)
}
