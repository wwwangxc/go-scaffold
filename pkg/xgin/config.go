package xgin

import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"
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

// DefaultConfig ..
func DefaultConfig() *Config {
	return &Config{
		Mode:            gin.DebugMode,
		Port:            8080,
		ShutdownTimeout: 5,
	}
}

// RawConfig ..
func RawConfig(confPrefix string, confHandler ConfigHandler) *Config {
	if strings.HasSuffix(confPrefix, ".") {
		confPrefix = confPrefix[:len(confPrefix)-1]
	}
	return &Config{
		Mode:            confHandler.GetString(confPrefix + ".mode"),
		Port:            confHandler.GetInt(confPrefix + ".port"),
		ShutdownTimeout: confHandler.GetInt(confPrefix + ".shutdown_timeout"),
	}
}

// Config ..
type Config struct {
	Mode            string
	Port            int
	ShutdownTimeout int

	fns []func(*gin.Engine)
}

// Setup ..
func (t *Config) Setup(fns ...func(*gin.Engine)) *Config {
	t.fns = fns
	return t
}

// Build ..
func (t *Config) Build() *HTTPServer {
	return newServer(t)
}
