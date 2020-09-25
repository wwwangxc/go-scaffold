package kafka

import "strings"

type (
	SetupHandler        func() error
	ConsumeClaimHandler func() error
	Cleanuphandler      func() error
)

// ProducerConfig ..
type ConsumerConfig struct {
	GroupID      string
	BalancerName string
	Version      string
	Topics       []string
	Brokers      []string
	ConsumerNum  int

	L LoggerHandler

	SetupCallBack        SetupHandler
	ConsumeClaimCallBack ConsumeClaimHandler
	CleanupCallBack      Cleanuphandler
}

// RawConsumerConfig ..
func RawConsumerGroupConfig(confPrefix string, confHandler ConfigHandler) *ConsumerConfig {
	if strings.HasSuffix(confPrefix, ".") {
		confPrefix = confPrefix[:len(confPrefix)-1]
	}
	return &ConsumerConfig{
		GroupID:      confHandler.GetString(confPrefix + ".group_id"),
		BalancerName: confHandler.GetString(confPrefix + ".balancer_name"),
		Version:      confHandler.GetString(confPrefix + ".version"),
		Topics:       confHandler.GetStringSlice(confPrefix + ".topics"),
		Brokers:      confHandler.GetStringSlice(confPrefix + ".brokers"),
		ConsumerNum:  confHandler.GetInt(confPrefix + ".consumer_num"),
	}
}

// WithSetup set setup callback.
func (t *ConsumerConfig) WithSetup(handler SetupHandler) *ConsumerConfig {
	t.SetupCallBack = handler
	return t
}

// WithSetup set consume claim callback.
func (t *ConsumerConfig) WithConsumeClaim(handler ConsumeClaimHandler) *ConsumerConfig {
	t.ConsumeClaimCallBack = handler
	return t
}

// WithSetup set cleanup callback.
func (t *ConsumerConfig) WithCleanup(handler Cleanuphandler) *ConsumerConfig {
	t.CleanupCallBack = handler
	return t
}

// WithLogger set log handler.
func (t *ConsumerConfig) WithLogger(handler LoggerHandler) *ConsumerConfig {
	t.L = handler
	return t
}

// BuildGroup build consumer group.
func (t *ConsumerConfig) Build() (*ConsumerGroup, error) {
	return newConsumerGroup(t)
}
