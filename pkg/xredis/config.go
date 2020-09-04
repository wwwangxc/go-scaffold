package xredis

import (
	"strings"
	"sync"
	"time"
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
		Addr:       confHandler.GetString(confPrefix + ".addr"),
		Password:   confHandler.GetString(confPrefix + ".password"),
		DB:         confHandler.GetInt(confPrefix + ".db"),
		MaxRetries: 3,
	}
}

var initOnce sync.Once

// Config ..
type Config struct {
	Addr       string
	Password   string
	DB         int
	MaxRetries int // 网络错误，最大重试次数
}

// Init ..
func (t *Config) Init() {
	initOnce.Do(func() {
		initialize(t)
	})
}

// Build ..
func (t *Config) Build() *Client {
	return newClient(t)
}
