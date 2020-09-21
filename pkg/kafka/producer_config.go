package kafka

import "strings"

// ConfigHandler ..
type ConfigHandler interface {
	GetString(key string) string
	GetStringSlice(key string) []string
}

// ProducerConfig ..
type ProducerConfig struct {
	Topic         string
	BrokerAddress []string
}

// RawProducerConfig ..
func RawProducerConfig(confPrefix string, confHandler ConfigHandler) *ProducerConfig {
	if strings.HasSuffix(confPrefix, ".") {
		confPrefix = confPrefix[:len(confPrefix)-1]
	}
	return &ProducerConfig{
		Topic:         confHandler.GetString(confPrefix + ".topic"),
		BrokerAddress: confHandler.GetStringSlice(confPrefix + ".broker_address"),
	}
}

// Build2Async ..
func (t *ProducerConfig) Build2Async() (*AsyncProducer, error) {
	return newAsyncProducer(t)
}

func (t *ProducerConfig) Build2Sync() (*SyncProducer, error) {
	return newSyncProducer(t)
}
