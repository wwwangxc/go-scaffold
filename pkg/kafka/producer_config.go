package kafka

import "strings"

// ProducerConfig ..
type ProducerConfig struct {
	Version string
	Topic   string
	Brokers []string
}

// RawProducerConfig ..
func RawProducerConfig(confPrefix string, confHandler ConfigHandler) *ProducerConfig {
	if strings.HasSuffix(confPrefix, ".") {
		confPrefix = confPrefix[:len(confPrefix)-1]
	}
	return &ProducerConfig{
		Version: confHandler.GetString(confPrefix + ".version"),
		Topic:   confHandler.GetString(confPrefix + ".topic"),
		Brokers: confHandler.GetStringSlice(confPrefix + ".brokers"),
	}
}

// Build2Async ..
func (t *ProducerConfig) Build2Async() (*AsyncProducer, error) {
	return newAsyncProducer(t)
}

func (t *ProducerConfig) Build() (*SyncProducer, error) {
	return newSyncProducer(t)
}
