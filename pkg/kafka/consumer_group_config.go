package kafka

import (
	"strings"

	"github.com/Shopify/sarama"
)

type (
	SetupHandler        func(sarama.ConsumerGroupSession) error
	ConsumeClaimHandler func(sarama.ConsumerGroupSession, sarama.ConsumerGroupClaim) error
	CleanupHandler      func(sarama.ConsumerGroupSession) error
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

	CallBackSetup        SetupHandler
	CallBackConsumeClaim ConsumeClaimHandler
	CallBackCleanup      CleanupHandler
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
	t.CallBackSetup = handler
	return t
}

// WithSetup set consume claim callback.
func (t *ConsumerGroupConfig) WithConsumeClaim(handler ConsumeClaimHandler) *ConsumerGroupConfig {
	t.CallBackConsumeClaim = handler
	return t
}

// WithSetup set cleanup callback.
func (t *ConsumerGroupConfig) WithCleanup(handler CleanupHandler) *ConsumerGroupConfig {
	t.CallBackCleanup = handler
	return t
}

// WithLogger set log handler.
func (t *ConsumerGroupConfig) WithLogger(handler LoggerHandler) *ConsumerGroupConfig {
	t.L = handler
	return t
}

// Build build consumer group.
func (t *ConsumerGroupConfig) Build() (*ConsumerGroup, error) {
	return newConsumerGroup(t)
}
