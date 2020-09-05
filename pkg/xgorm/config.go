package xgorm

import (
	"strings"
	"sync"

	"github.com/jinzhu/gorm"
)

var _initOnce sync.Once

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
	if strings.HasSuffix(confPrefix, ".") {
		confPrefix = confPrefix[:len(confPrefix)-1]
	}
	return &Config{
		DSN: confHandler.GetString(confPrefix + ".dsn"),
	}
}

// Init ..
func (t *Config) Init() {
	_initOnce.Do(func() {
		initialize(t)
	})
}

// Build ..
func (t *Config) Build() *gorm.DB {
	return newClient(t)
}
