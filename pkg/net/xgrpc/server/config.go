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

// Config ..
type Config struct {
	Port    int
	Network string
	Scheme  string
	Name    string

	services []func(*grpc.Server)
	register Register
}

// DefaultConfig ..
func DefaultConfig() *Config {
	return &Config{
		Port:    3000,
		Network: "tcp",
		Scheme:  "schema",
		Name:    "application",
	}
}

// RawConfig ..
func RawConfig(confPrefix string, confHandler ConfigHandler) *Config {
	if strings.HasSuffix(confPrefix, ".") {
		confPrefix = confPrefix[:len(confPrefix)-1]
	}
	return &Config{
		Port:    3000,
		Network: confHandler.GetString(confPrefix + ".network"),
		Scheme:  confHandler.GetString(confPrefix + ".scheme"),
		Name:    confHandler.GetString(confPrefix + ".name"),
	}
}

// WithPort ..
func (t *Config) WithPort(port int) *Config {
	t.Port = port
	return t
}

// WithService ..
func (t *Config) WithService(services ...func(*grpc.Server)) *Config {
	t.services = services
	return t
}

// SetRegistry ..
func (t *Config) WithRegister(register Register) *Config {
	t.register = register
	return t
}

// Build ..
func (t *Config) Build() *GrpcServer {
	return newServer(t)
}
