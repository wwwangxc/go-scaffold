package server

import (
	"strings"
	"time"

	"google.golang.org/grpc"
)

// ConfigHandler ..
type ConfigHandler interface {
	GetString(key string) string
	GetInt(key string) int
	GetDuration(key string) time.Duration
}

// Register ..
type Register interface {
	RegistryService(key, value string)
	UnRegistryService(key string)
}

// DefaultConfig ..
func DefaultConfig() *ServerConfig {
	return &ServerConfig{
		Network: "tcp",
		Addr:    "127.0.0.1:3000",
		Scheme:  "schema",
		Name:    "application",
	}
}

// RawConfig ..
func RawConfig(confPrefix string, confHandler ConfigHandler) *ServerConfig {
	if strings.HasSuffix(confPrefix, ".") {
		confPrefix = confPrefix[:len(confPrefix)-1]
	}
	return &ServerConfig{
		Network: confHandler.GetString(confPrefix + ".network"),
		Addr:    confHandler.GetString(confPrefix + ".addr"),
		Scheme:  confHandler.GetString(confPrefix + ".scheme"),
		Name:    confHandler.GetString(confPrefix + ".name"),
	}
}

// ServerConfig ..
type ServerConfig struct {
	Addr    string
	Network string
	Scheme  string
	Name    string

	services []func(*grpc.Server)
	register Register
}

// WithService ..
func (t *ServerConfig) WithService(services ...func(*grpc.Server)) *ServerConfig {
	t.services = services
	return t
}

// SetRegistry ..
func (t *ServerConfig) WithRegister(register Register) *ServerConfig {
	t.register = register
	return t
}

// Build ..
func (t *ServerConfig) Build() *GrpcServer {
	return newServer(t)
}
