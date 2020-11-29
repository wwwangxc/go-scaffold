package group

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
	Version      string
	GroupID      string
	BalancerName string
	Topics       []string
	Brokers      []string
	ConsumerNum  int

	CallbackSetup        ConsumerGroupHandlerSetup
	CallbackConsumeClaim ConsumerGroupHandlerConsumeClaim
	CallbackCleanup      ConsumerGroupHandlerCleanup
	CallbackError        ConsumerGroupHandlerError
}

// RawConsumerConfig ..
func RawConfig(confPrefix string, confHandler ConfigHandler) *Config {
	if strings.HasSuffix(confPrefix, ".") {
		confPrefix = confPrefix[:len(confPrefix)-1]
	}
	return &Config{
		GroupID:      confHandler.GetString(confPrefix + ".group_id"),
		BalancerName: confHandler.GetString(confPrefix + ".balancer_name"),
		Version:      confHandler.GetString(confPrefix + ".version"),
		Topics:       confHandler.GetStringSlice(confPrefix + ".topics"),
		Brokers:      confHandler.GetStringSlice(confPrefix + ".brokers"),
		ConsumerNum:  confHandler.GetInt(confPrefix + ".consumer_num"),
	}
}

// WithCallbackSetup set setup callback.
//
// setup is run at the beginning of a new session, before ConsumeClaim.
func (t *Config) WithCallbackSetup(handler ConsumerGroupHandlerSetup) *Config {
	t.CallbackSetup = handler
	return t
}

// WithCallbackConsumeClaim set consume claim callback.
//
// triggered every time when a message is received.
func (t *Config) WithCallbackConsumeClaim(handler ConsumerGroupHandlerConsumeClaim) *Config {
	t.CallbackConsumeClaim = handler
	return t
}

// WithCallbackCleanup set cleanup callback.
//
// cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited.
func (t *Config) WithCallbackCleanup(handler ConsumerGroupHandlerCleanup) *Config {
	t.CallbackCleanup = handler
	return t
}

// Build build consumer group.
func (t *Config) Build() (*ConsumerGroup, error) {
	return newConsumerGroup(t)
}
