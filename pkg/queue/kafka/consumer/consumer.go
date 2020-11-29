package consumer

import (
	"sync"

	"github.com/Shopify/sarama"
	"github.com/pkg/errors"
)

type (
	// ConsumerHandlerError triggered every time when a error message is reveived.
	ConsumerHandlerError func(*sarama.ConsumerError)

	// ConsumerHandlerSuccess triggered every time when a success message is received.
	ConsumerHandlerSuccess func(*sarama.ConsumerMessage)
)

var (
	ErrConsumerClosed = errors.New("[kafka] the consumer is closed")
)

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
// return ConsumerErrClosed error when the consumer is closed.
// triggered CallbackError handle every time when a error message is received.
func (t *Consumer) Watch() (<-chan *sarama.ConsumerMessage, error) {
	t.m.Lock()
	defer t.m.Unlock()

	if t.closed {
		return nil, ErrConsumerClosed
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
// return ConsumerErrClosed error when the consumer is closed.
// triggered CallbackError handle every time when a error message is received.
// triggered CallbackSuccess handle every time when a success message is received.
func (t *Consumer) AsyncWatch() error {
	t.m.Lock()
	defer t.m.Unlock()

	if t.closed {
		return ErrConsumerClosed
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
