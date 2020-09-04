package log

import (
	"fmt"
	"os"

	"github.com/natefinch/lumberjack"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func initializa(config *Config) {
	encoder := newJSONEncoder()
	lv := zap.NewAtomicLevelAt(zapcore.InfoLevel)
	if err := lv.UnmarshalText([]byte(config.Level)); err != nil {
		panic(err.Error())
	}
	zapcores := []zapcore.Core{
		zapcore.NewCore(
			encoder,
			newWriter(config),
			lv,
		),
		zapcore.NewCore(
			encoder,
			newErrorWriter(config),
			zapcore.ErrorLevel,
		),
	}
	if config.Debug {
		zapcores = append(zapcores, zapcore.NewCore(
			zapcore.NewConsoleEncoder(zap.NewDevelopmentEncoderConfig()),
			zapcore.Lock(os.Stdout),
			zapcore.DebugLevel,
		))
	}
	zap.ReplaceGlobals(
		zap.New(
			zapcore.NewTee(zapcores...), zap.AddCaller()))
}

// 日志生成格式
func newJSONEncoder() zapcore.Encoder {
	return zapcore.NewJSONEncoder(zapcore.EncoderConfig{
		TimeKey:       "timestamp",
		LevelKey:      "level",
		NameKey:       "logger",
		CallerKey:     "caller",
		MessageKey:    "message",
		StacktraceKey: "stacktrace",
		LineEnding:    zapcore.DefaultLineEnding,
		EncodeLevel:   zapcore.LowercaseLevelEncoder,
		// EncodeTime:     zapcore.ISO8601TimeEncoder, // 人类可读的日期格式
		EncodeTime:     zapcore.EpochTimeEncoder, // 时间戳日期格式，
		EncodeDuration: zapcore.SecondsDurationEncoder,
		EncodeCaller:   zapcore.ShortCallerEncoder,
	})
}

// 日志生成位置
func newWriter(config *Config) zapcore.WriteSyncer {
	return zapcore.AddSync(&lumberjack.Logger{
		Filename:   fmt.Sprintf("%s/%s", config.Dir, config.Name),
		MaxSize:    config.MaxSize,    // 日志文件大小，单位：MB
		MaxBackups: config.MaxBackups, // 备份数量
		MaxAge:     config.MaxAge,     // 备份时间，单位：天
		Compress:   config.Compress,   // 是否压缩
	})
}

// 日志生成位置
func newErrorWriter(config *Config) zapcore.WriteSyncer {
	return zapcore.AddSync(&lumberjack.Logger{
		Filename:   fmt.Sprintf("%s/%s.err", config.Dir, config.Name),
		MaxSize:    config.MaxSize,    // 日志文件大小，单位：MB
		MaxBackups: config.MaxBackups, // 备份数量
		MaxAge:     config.MaxAge,     // 备份时间，单位：天
		Compress:   config.Compress,   // 是否压缩
	})
}
