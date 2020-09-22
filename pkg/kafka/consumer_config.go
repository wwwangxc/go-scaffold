package kafka

import "strings"

// ProducerConfig ..
type ConsumerConfig struct {
	GroupID      string
	BalancerName string
	Topics       []string
	Brokers      []string
}

// RawConsumerConfig ..
func RawConsumerConfig(confPrefix string, confHandler ConfigHandler) *ConsumerConfig {
	if strings.HasSuffix(confPrefix, ".") {
		confPrefix = confPrefix[:len(confPrefix)-1]
	}
	return &ConsumerConfig{
		GroupID:      confHandler.GetString(confPrefix + ".group_id"),
		BalancerName: confHandler.GetString(confPrefix + ".balancer_name"),
		Topics:       confHandler.GetStringSlice(confPrefix + ".topics"),
		Brokers:      confHandler.GetStringSlice(confPrefix + "brokers"),
	}
}
