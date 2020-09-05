package xredis

import (
	"strings"
	"sync"
)

var _initOnce sync.Once

// ConfigHandler ..
type ConfigHandler interface {
	GetInt(key string) int
	GetString(key string) string
}

// Config ..
type Config struct {
	Addr       string
	Password   string
	DB         int
	MaxRetries int // 网络错误，最大重试次数
}

// RawConfig ..
func RawConfig(confPrefix string, confHandler ConfigHandler) *Config {
	if strings.HasSuffix(confPrefix, ".") {
		confPrefix = confPrefix[:len(confPrefix)-1]
	}
	return &Config{
		Addr:       confHandler.GetString(confPrefix + ".addr"),
		Password:   confHandler.GetString(confPrefix + ".password"),
		DB:         confHandler.GetInt(confPrefix + ".db"),
		MaxRetries: 3,
	}
}

// Init ..
func (t *Config) Init() {
	_initOnce.Do(func() {
		initialize(t)
	})
}

// Build ..
func (t *Config) Build() *Client {
	return newClient(t)
}
