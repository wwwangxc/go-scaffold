package kafka

import (
	"context"
	"sync"

	"github.com/Shopify/sarama"
)

// ConsumerGroup
type ConsumerGroup struct {
	config *ConsumerGroupConfig
	cg     sarama.ConsumerGroup

	ctx       context.Context
	ctxCancel context.CancelFunc

	watching bool
	closed   bool

	m sync.Mutex
}

// create consumer group instance.
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

// Watch the topics.
// return ConsumerGroupErrClosed error when consumer group closed.
// run CallbackError handle at the occurs of an error, when consume joins a cluster of consumers for a given list of topics.
func (t *ConsumerGroup) Watch() error {
	t.m.Lock()
	defer t.m.Unlock()

	if t.closed {
		return ConsumerGroupErrClosed
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

// Consumer ..
type Consumer struct {
	config *ConsumerConfig

	c sarama.Consumer
	p sarama.PartitionConsumer

	watching bool
	closed   bool

	m sync.Mutex
}

// create consumer instance.
func newConsumer(conf *ConsumerConfig) (*Consumer, error) {
	c := sarama.NewConfig()
	version, err := sarama.ParseKafkaVersion(conf.Version)
	if err != nil {
		return nil, err
	}
	c.Version = version
	c.ClientID = conf.ClientID
	c.Producer.Return.Errors = true
	if conf.SASL.Enable {
		c.Net.SASL.Enable = true
		c.Net.SASL.User = conf.SASL.Username
		c.Net.SASL.Password = conf.SASL.Password
	}
	if conf.TLS.Enable {
		c.Net.TLS.Enable = true
		tlsConfig, err := conf.TLS.tlsConfig()
		if err != nil {
			return nil, err
		}
		c.Net.TLS.Config = tlsConfig
	}

	consumer, err := sarama.NewConsumer(conf.Brokers, c)
	if err != nil {
		return nil, err
	}
	return &Consumer{
		config:   conf,
		c:        consumer,
		watching: false,
		closed:   false,
	}, nil
}

// Watch sync watch the topic.
// return ConsumerErrClosed err when consumer closed.
// triggered CallbackError handle every time when a error message is received.
func (t *Consumer) Watch() (<-chan *sarama.ConsumerMessage, error) {
	t.m.Lock()
	defer t.m.Unlock()

	if t.closed {
		return nil, ConsumerErrClosed
	}

	if t.watching {
		return t.p.Messages(), nil
	}

	partition, err := t.c.ConsumePartition(t.config.Topic, t.config.Partition, sarama.OffsetNewest)
	if err != nil {
		return nil, err
	}

	go func() {
		for {
			consumerError, ok := <-partition.Errors()
			if ok && t.config.CallbackError != nil {
				t.config.CallbackError(consumerError)
			}
		}
	}()

	t.p = partition
	t.watching = true
	return partition.Messages(), nil
}

// AsyncWatch async watch the topic.
// return ConsumerErrClosed err when consumer closed.
// triggered CallbackError handle every time when a error message is received.
// triggered CallbackSuccess handle every time when a success message is received.
func (t *Consumer) AsyncWatch() error {
	t.m.Lock()
	defer t.m.Unlock()

	if t.closed {
		return ConsumerErrClosed
	}

	if t.watching {
		return nil
	}

	partition, err := t.c.ConsumePartition(t.config.Topic, t.config.Partition, sarama.OffsetNewest)
	if err != nil {
		return nil
	}

	go func() {
		for {
			select {
			case consumerError, ok := <-partition.Errors():
				if ok && t.config.CallbackError != nil {
					go t.config.CallbackError(consumerError)
				}
			case consumerMessage, ok := <-partition.Messages():
				if ok && t.config.CallbackSuccess != nil {
					go t.config.CallbackSuccess(consumerMessage)
				}
			}
		}
	}()

	t.p = partition
	t.watching = true
	return nil
}

// Close consumer.
func (t *Consumer) Close() error {
	t.m.Lock()
	defer t.m.Unlock()

	if t.closed {
		return nil
	}

	var err error

	if t.p != nil {
		err = t.p.Close()
	}
	if err != nil {
		return err
	}

	if t.c != nil {
		err = t.c.Close()
	}

	t.closed = true
	return err
}
