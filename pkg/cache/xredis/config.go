package xredis

import (
	"strings"
)

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
	confPrefix = strings.TrimSuffix(confPrefix, ".")

	return &Config{
		Addr:       confHandler.GetString(confPrefix + ".addr"),
		Password:   confHandler.GetString(confPrefix + ".password"),
		DB:         confHandler.GetInt(confPrefix + ".db"),
		MaxRetries: 3,
	}
}

// Build ..
func (t *Config) Build() *Client {
	return newClient(t)
}

// Append ..
func (t *Config) Append(storeName string) *Client {
	cli := newClient(t)
	store.append(storeName, cli)
	return cli
}
