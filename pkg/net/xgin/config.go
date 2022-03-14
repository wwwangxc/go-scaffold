package xgin

import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

// ConfigHandler ..
type ConfigHandler interface {
	GetInt(key string) int
	GetString(key string) string
	GetDuration(key string) time.Duration
}

// Config ..
type Config struct {
	Mode        string
	Port        int
	ShutdownTTL time.Duration

	middlewares []func(*gin.Engine)
	routes      []func(*gin.Engine)
}

// DefaultConfig ..
func DefaultConfig() *Config {
	return &Config{
		Mode:        gin.DebugMode,
		Port:        8080,
		ShutdownTTL: 5 * time.Second,
	}
}

// RawConfig ..
func RawConfig(confPrefix string, confHandler ConfigHandler) *Config {
	confPrefix = strings.TrimSuffix(confPrefix, ".")

	return &Config{
		Mode:        confHandler.GetString(confPrefix + ".mode"),
		Port:        8080,
		ShutdownTTL: confHandler.GetDuration(confPrefix + ".shutdown_ttl"),
	}
}

func (t *Config) WithPort(port int) *Config {
	t.Port = port
	return t
}

// WithMiddlewares ..
func (t *Config) WithMiddlewares(middlewares ...func(*gin.Engine)) *Config {
	t.middlewares = middlewares
	return t
}

// WithRoutes ..
func (t *Config) WithRoutes(routes ...func(*gin.Engine)) *Config {
	t.routes = routes
	return t
}

// Build ..
func (t *Config) Build() *HTTPServer {
	return newServer(t)
}
