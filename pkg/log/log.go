package log

import (
	"go.uber.org/zap"
)

var (
	logger = DefaultConfig().Build()
)

// Sync ..
func Sync() {
	if logger == nil {
		return
	}
	_ = logger.Sync()
}

// Debug ..
func Debug(msg string, fields ...zap.Field) {
	if logger == nil {
		return
	}
	logger.Debug(msg, fields...)
}

// Info ..
func Info(msg string, fields ...zap.Field) {
	if logger == nil {
		return
	}
	logger.Info(msg, fields...)
}

// Warn ..
func Warn(msg string, fields ...zap.Field) {
	if logger == nil {
		return
	}
	logger.Warn(msg, fields...)
}

// Error ..
func Error(msg string, fields ...zap.Field) {
	if logger == nil {
		return
	}
	logger.Error(msg, append(fields, zap.Stack("stack"))...)
}

// Panic ..
func Panic(msg string, fields ...zap.Field) {
	if logger == nil {
		return
	}
	logger.Panic(msg, fields...)
}

// Fatal ..
func Fatal(msg string, fields ...zap.Field) {
	if logger == nil {
		return
	}
	logger.Fatal(msg, fields...)
}
