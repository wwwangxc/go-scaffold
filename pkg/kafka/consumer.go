package kafka

import (
	"context"

	"github.com/Shopify/sarama"
)

type ConsumerGroup struct {
	cg        sarama.ConsumerGroup
	config    *ConsumerConfig
	ctx       context.Context
	ctxCancel context.CancelFunc
}

func newConsumerGroup(conf *ConsumerConfig) (*ConsumerGroup, error) {
	c := sarama.NewConfig()
	switch conf.BalancerName {
	case "sticky":
		c.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategySticky
	case "roundrobin":
		c.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin
	case "range":
		c.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRange
	default:
		c.Consumer.Group.Rebalance.Strategy = sarama.BalanceStrategyRoundRobin
	}
	ctx, cancel := context.WithCancel(context.Background())
	cg, err := sarama.NewConsumerGroup(conf.Brokers, conf.GroupID, c)
	if err != nil {
		cancel()
		return nil, err
	}
	return &ConsumerGroup{
		cg:        cg,
		config:    conf,
		ctx:       ctx,
		ctxCancel: cancel,
	}, nil
}

type consumer struct{}

// Setup ..
func (t *consumer) Setup(sess sarama.ConsumerGroupSession) error {
	return nil
}

// ConsumeClaim ..
func (t *consumer) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		sess.MarkMessage(message, "")
	}
	return nil
}

// Cleanup ..
func (t *consumer) Cleanup(sess sarama.ConsumerGroupSession) error {
	return nil
}
