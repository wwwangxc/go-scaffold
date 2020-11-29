package config

import (
	"time"

	"github.com/spf13/viper"
)

// GetHandler ..
func GetHandler() *viper.Viper {
	return viper.GetViper()
}

// GetInt ..
func GetInt(key string) int {
	return viper.GetInt(key)
}

// GetInt32 ..
func GetInt32(key string) int32 {
	return viper.GetInt32(key)
}

// GetInt64 ..
func GetInt64(key string) int64 {
	return viper.GetInt64(key)
}

// GetString ..
func GetString(key string) string {
	return viper.GetString(key)
}

// GetBool ..
func GetBool(key string) bool {
	return viper.GetBool(key)
}

// GetTime ..
func GetTime(key string) time.Time {
	return viper.GetTime(key)
}

// GetDuration ..
func GetDuration(key string) time.Duration {
	return viper.GetDuration(key)
}

// InConfig ..
func InConfig(key string) bool {
	return viper.InConfig(key)
}

// AllKeys ..
func AllKeys() []string {
	viper.GetViper()
	return viper.AllKeys()
}
