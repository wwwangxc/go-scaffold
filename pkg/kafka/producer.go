package kafka

import (
	"sync"

	"github.com/Shopify/sarama"
)

// Producer ..
type AsyncProducer struct {
	p sarama.AsyncProducer

	config *ProducerConfig
	rwmu   sync.RWMutex
}

func newAsyncProducer(conf *ProducerConfig) (*AsyncProducer, error) {
	version, err := sarama.ParseKafkaVersion(conf.Version)
	if err != nil {
		return nil, err
	}
	c := sarama.NewConfig()
	c.Version = version
	c.Producer.RequiredAcks = sarama.WaitForAll
	c.Producer.Partitioner = sarama.NewRandomPartitioner
	producer, err := sarama.NewAsyncProducer(conf.Brokers, c)
	if err != nil {
		return nil, err
	}
	return &AsyncProducer{
		config: conf,
		p:      producer,
	}, nil
}

// Send ..
func (t *AsyncProducer) Send(message string, callBack func(err error)) {
	t.rwmu.RLock()
	defer t.rwmu.RUnlock()
	t.p.Input() <- &sarama.ProducerMessage{
		Topic: t.config.Topic,
		Value: sarama.StringEncoder(message),
	}
	go func() {
		err := <-t.p.Errors()
		if callBack != nil {
			callBack(err)
		}
	}()
}

// Close ..
func (t *AsyncProducer) Close() {
	t.rwmu.Lock()
	defer t.rwmu.Unlock()
	t.p.AsyncClose()
}

// SyncProducer ..
type SyncProducer struct {
	p      sarama.SyncProducer
	config *ProducerConfig
	rwmu   sync.RWMutex
}

func newSyncProducer(conf *ProducerConfig) (*SyncProducer, error) {
	version, err := sarama.ParseKafkaVersion(conf.Version)
	if err != nil {
		return nil, err
	}
	c := sarama.NewConfig()
	c.Version = version
	c.Producer.RequiredAcks = sarama.WaitForAll
	c.Producer.Partitioner = sarama.NewRandomPartitioner
	c.Producer.Return.Successes = true
	producer, err := sarama.NewSyncProducer(conf.Brokers, c)
	if err != nil {
		return nil, err
	}
	return &SyncProducer{
		p:      producer,
		config: conf,
	}, nil
}

// Send ..
func (t *SyncProducer) Send(message string) error {
	t.rwmu.RLock()
	defer t.rwmu.RUnlock()
	_, _, err := t.p.SendMessage(&sarama.ProducerMessage{
		Topic: t.config.Topic,
		Value: sarama.StringEncoder(message),
	})
	return err
}

// Close ..
func (t *SyncProducer) Close() {
	t.rwmu.Lock()
	defer t.rwmu.Unlock()
	t.p.Close()
}
