package kafka

import "strings"

type (
	SetupHandler        func() error
	ConsumeClaimHandler func() error
	Cleanuphandler      func() error
)

// ProducerConfig ..
type ConsumerGroupConfig struct {
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
func RawConsumerGroupConfig(confPrefix string, confHandler ConfigHandler) *ConsumerGroupConfig {
	if strings.HasSuffix(confPrefix, ".") {
		confPrefix = confPrefix[:len(confPrefix)-1]
	}
	return &ConsumerGroupConfig{
		GroupID:      confHandler.GetString(confPrefix + ".group_id"),
		BalancerName: confHandler.GetString(confPrefix + ".balancer_name"),
		Version:      confHandler.GetString(confPrefix + ".version"),
		Topics:       confHandler.GetStringSlice(confPrefix + ".topics"),
		Brokers:      confHandler.GetStringSlice(confPrefix + ".brokers"),
		ConsumerNum:  confHandler.GetInt(confPrefix + ".consumer_num"),
	}
}

// WithSetup set setup callback.
func (t *ConsumerGroupConfig) WithSetup(handler SetupHandler) *ConsumerGroupConfig {
	t.SetupCallBack = handler
	return t
}

// WithSetup set consume claim callback.
func (t *ConsumerGroupConfig) WithConsumeClaim(handler ConsumeClaimHandler) *ConsumerGroupConfig {
	t.ConsumeClaimCallBack = handler
	return t
}

// WithSetup set cleanup callback.
func (t *ConsumerGroupConfig) WithCleanup(handler Cleanuphandler) *ConsumerGroupConfig {
	t.CleanupCallBack = handler
	return t
}

// WithLogger set log handler.
func (t *ConsumerGroupConfig) WithLogger(handler LoggerHandler) *ConsumerGroupConfig {
	t.L = handler
	return t
}

// BuildGroup build consumer group.
func (t *ConsumerGroupConfig) Build() (*ConsumerGroup, error) {
	return newConsumerGroup(t)
}
