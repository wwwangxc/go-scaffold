package group

import (
	"context"
	"sync"

	"github.com/pkg/errors"

	"github.com/Shopify/sarama"
)

var (
	ErrConsumerGroupClosed = errors.New("[kafka] the consumer group is closed")
)

type (
	// ConsumerGroupHandlerSetup is run at the beginning of a new session, before ConsumeClaim.
	ConsumerGroupHandlerSetup func(sarama.ConsumerGroupSession)

	// ConsumerGroupHandlerConsumerClaim triggered every time when a message is received.
	ConsumerGroupHandlerConsumeClaim func(sarama.ConsumerGroupSession, sarama.ConsumerGroupClaim)

	// ConsumerGroupHandlerCleanup is run at the end of a session, once all ConsumeClaim goroutines have exited.
	ConsumerGroupHandlerCleanup func(sarama.ConsumerGroupSession)

	// ConsumerGroupHandlerError is run at the occurs of an error, when consume joins a cluster of consumers for a given list of topics
	ConsumerGroupHandlerError func(error)
)

// ConsumerGroup
type ConsumerGroup struct {
	config *Config
	cg     sarama.ConsumerGroup

	ctx       context.Context
	ctxCancel context.CancelFunc

	watching bool
	closed   bool

	m sync.Mutex
}

// create consumer group instance.
func newConsumerGroup(conf *Config) (*ConsumerGroup, error) {
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

// Watch the topics.
// return ConsumerGroupErrClosed error when the consumer group is closed.
// run CallbackError handle at the occurs of an error, when consume joins a cluster of consumers for a given list of topics.
func (t *ConsumerGroup) Watch() error {
	t.m.Lock()
	defer t.m.Unlock()

	if t.closed {
		return ErrConsumerGroupClosed
	}
	if t.watching {
		return nil
	}

	for i := 0; i < t.config.ConsumerNum; i++ {
		go func() {
			handler := &consumerHandler{
				callBackSetup:        t.config.CallbackSetup,
				callBackConsumeClaim: t.config.CallbackConsumeClaim,
				callBackCleanup:      t.config.CallbackCleanup,
			}
			for {
				if err := t.cg.Consume(t.ctx, t.config.Topics, handler); err != nil {
					if t.config.CallbackError != nil {
						t.config.CallbackError(err)
					}
				}
				if t.ctx.Err() != nil {
					break
				}
			}
		}()
	}

	t.watching = true
	return nil
}

// Close consumer group.
func (t *ConsumerGroup) Close() error {
	t.m.Lock()
	defer t.m.Unlock()
	if t.closed {
		return nil
	}
	if err := t.cg.Close(); err != nil {
		return err
	}
	t.ctxCancel()
	t.closed = true
	return nil
}

type consumerHandler struct {
	callBackSetup        ConsumerGroupHandlerSetup
	callBackConsumeClaim ConsumerGroupHandlerConsumeClaim
	callBackCleanup      ConsumerGroupHandlerCleanup
}

// Setup is run at the beginning of a new session, before ConsumeClaim
func (t *consumerHandler) Setup(sess sarama.ConsumerGroupSession) error {
	if t.callBackSetup != nil {
		t.callBackSetup(sess)
	}
	return nil
}

// ConsumeClaim triggered every time when a message is received.
func (t *consumerHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for message := range claim.Messages() {
		if t.callBackConsumeClaim != nil {
			t.callBackConsumeClaim(sess, claim)
		}
		sess.MarkMessage(message, "")
	}
	return nil
}

// Cleanup is run at the end of a session, once all ConsumeClaim goroutines have exited
func (t *consumerHandler) Cleanup(sess sarama.ConsumerGroupSession) error {
	if t.callBackCleanup != nil {
		t.callBackCleanup(sess)
	}
	return nil
}
