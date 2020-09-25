package kafka

// LoggerHandler ..
type LoggerHandler interface {
	Info(message string)
	Error(message string)
}

// ConfigHandler ..
type ConfigHandler interface {
	GetString(key string) string
	GetStringSlice(key string) []string
	GetInt(key string) int
}
