package log

import (
	"strings"
	"sync"

	"go.uber.org/zap"
)

var initOnce sync.Once

// ConfigHandler ..
type ConfigHandler interface {
	GetInt(key string) int
	GetString(key string) string
	GetBool(key string) bool
}

// Config ..
type Config struct {
	Dir        string // 日志输出目录
	Name       string // 日志文件名称
	Level      string // 日志输出等级
	MaxSize    int    // 日志文件大小，单位：MB
	MaxAge     int    // 备份数量
	MaxBackups int    // 备份时间，单位：天
	Debug      bool   // 是否控制台输出dibug日志
	Compress   bool   // 是否压缩日志
}

// DefaultConfig ..
func DefaultConfig() *Config {
	return &Config{
		Dir:        ".",
		Name:       "default.log",
		Level:      "info",
		Debug:      true,
		MaxSize:    100,
		MaxAge:     5,
		MaxBackups: 7,
		Compress:   false,
	}
}

// RawConfig ..
func RawConfig(confPrefix string, confHandler ConfigHandler) *Config {
	confPrefix = strings.TrimSuffix(confPrefix, ".")

	return &Config{
		Dir:        confHandler.GetString(confPrefix + ".dir"),
		Name:       confHandler.GetString(confPrefix + ".name"),
		Level:      confHandler.GetString(confPrefix + ".level"),
		Debug:      confHandler.GetBool(confPrefix + ".debug"),
		MaxSize:    confHandler.GetInt(confPrefix + ".max_size"),
		MaxAge:     confHandler.GetInt(confPrefix + ".max_age"),
		MaxBackups: confHandler.GetInt(confPrefix + ".max_backups"),
		Compress:   confHandler.GetBool(confPrefix + ".compress"),
	}
}

// Init ..
func (t *Config) Init() {
	initOnce.Do(func() {
		initializa(t)
		logger = zap.L()
	})
}

// Build ..
func (t *Config) Build() *zap.Logger {
	initializa(t)
	return zap.L()
}
