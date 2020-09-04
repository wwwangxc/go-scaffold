package log

import "go.uber.org/zap"

var (
	logger = DefaultConfig().Build()
)

// Sync ..
func Sync() {
	logger.Sync()
}

// Debug ..
func Debug(msg string, fields ...zap.Field) {
	logger.Debug(msg, fields...)
}

// Info ..
func Info(msg string, fields ...zap.Field) {
	logger.Info(msg, fields...)
}

// Warn ..
func Warn(msg string, fields ...zap.Field) {
	logger.Warn(msg, fields...)
}

// Error ..
func Error(msg string, fields ...zap.Field) {
	logger.Error(msg, fields...)
}

// Panic ..
func Panic(msg string, fields ...zap.Field) {
	logger.Panic(msg, fields...)
}

// Fatal ..
func Fatal(msg string, fields ...zap.Field) {
	logger.Fatal(msg, fields...)
}
