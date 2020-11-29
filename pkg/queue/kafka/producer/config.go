package producer

import "strings"

// ConfigHandler ..
type ConfigHandler interface {
	GetString(string) string
	GetStringSlice(string) []string
	GetInt(string) int
	GetInt32(string) int32
	GetBool(string) bool
}

// Config ..
type Config struct {
	Version string
	Topic   string
	Brokers []string

	CallbackError ProducerHandlerError
}

// RawConfig ..
func RawConfig(confPrefix string, confHandler ConfigHandler) *Config {
	if strings.HasSuffix(confPrefix, ".") {
		confPrefix = confPrefix[:len(confPrefix)-1]
	}
	return &Config{
		Version: confHandler.GetString(confPrefix + ".version"),
		Topic:   confHandler.GetString(confPrefix + ".topic"),
		Brokers: confHandler.GetStringSlice(confPrefix + ".brokers"),
	}
}

// WithCallbackError set error callback.
//
// triggered every time when a error message is reveived.
func (t *Config) WithCallbackError(handler ProducerHandlerError) *Config {
	t.CallbackError = handler
	return t
}

// Build2Async build async producer.
func (t *Config) Build2Async() (*AsyncProducer, error) {
	return newAsyncProducer(t)
}

// Build build sync producer.
func (t *Config) Build() (*Producer, error) {
	return newProducer(t)
}
