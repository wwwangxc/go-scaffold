package kafka

import "strings"

// ProducerConfig ..
type ProducerConfig struct {
	Version string
	Topic   string
	Brokers []string

	CallbackError ProducerHandlerError
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

// WithCallbackError set error callback.
//
// triggered every time when a error message is reveived.
func (t *ProducerConfig) WithCallbackError(handler ProducerHandlerError) *ProducerConfig {
	t.CallbackError = handler
	return t
}

// Build2Async build async producer.
func (t *ProducerConfig) Build2Async() (*AsyncProducer, error) {
	return newAsyncProducer(t)
}

// Build build sync producer.
func (t *ProducerConfig) Build() (*Producer, error) {
	return newProducer(t)
}
