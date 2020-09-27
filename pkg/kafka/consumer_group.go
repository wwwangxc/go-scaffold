package kafka

import (
	"context"
	"errors"
	"sync"

	"github.com/Shopify/sarama"
)

// ConsumerGroup
type ConsumerGroup struct {
	cg        sarama.ConsumerGroup
	config    *ConsumerGroupConfig
	ctx       context.Context
	ctxCancel context.CancelFunc

	watching bool
	closed   bool

	m sync.Mutex
}

func newConsumerGroup(conf *ConsumerGroupConfig) (*ConsumerGroup, error) {
	c := sarama.NewConfig()
	version, err := sarama.ParseKafkaVersion(conf.Version)
	if err != nil {
		return nil, err
	}
	c.Version = version
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
		watching:  false,
		closed:    false,
	}, nil
}

// Watch ..
func (t *ConsumerGroup) Watch() error {
	if t.closed {
		return errors.New("the consumer group is closed")
	}

	t.m.Lock()
	defer t.m.Unlock()
	if t.watching {
		return nil
	}

	for i := 0; i < t.config.ConsumerNum; i++ {
		go func(group *ConsumerGroup) {
			handler := &consumer{
				l:                    group.config.L,
				callBackSetup:        group.config.CallBackSetup,
				callBackConsumeClaim: group.config.CallBackConsumeClaim,
				callBackCleanup:      group.config.CallBackCleanup,
			}
			for {
				if err := group.cg.Consume(group.ctx, group.config.Topics, handler); err != nil {
					if group.config.L != nil {
						group.config.L.Error(err.Error())
					}
				}
				if group.ctx.Err() != nil {
					break
				}
			}
		}(t)
	}

	t.watching = true
	return nil
}

// Close ..
func (t *ConsumerGroup) Close() {
	t.m.Lock()
	defer t.m.Unlock()
	if t.closed {
		return
	}
	t.cg.Close()
	t.ctxCancel()
	t.closed = true
}

type consumer struct {
	l LoggerHandler

	callBackSetup        SetupHandler
	callBackConsumeClaim ConsumeClaimHandler
	callBackCleanup      CleanupHandler
}

// Setup ..
func (t *consumer) Setup(sess sarama.ConsumerGroupSession) error {
	var err error
	if t.callBackSetup != nil {
		err = t.callBackSetup(sess)
	}
	return err
}

// ConsumeClaim ..
func (t *consumer) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		if t.callBackConsumeClaim != nil {
			err := t.callBackConsumeClaim(sess, claim)
			if err != nil && t.l != nil {
				t.l.Error(err.Error())
			}
		}
		sess.MarkMessage(message, "")
	}
	return nil
}

// Cleanup ..
func (t *consumer) Cleanup(sess sarama.ConsumerGroupSession) error {
	var err error
	if t.callBackCleanup != nil {
		err = t.callBackCleanup(sess)
	}
	return err
}
