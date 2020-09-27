package kafka

import (
	"fmt"
	"testing"
	"time"

	"github.com/Shopify/sarama"
)

func TestConsumerGroup(t *testing.T) {
	conf := &ConsumerGroupConfig{

		GroupID:      "groupid",
		BalancerName: "roundrobin",
		Topics:       []string{"app_test", "app_test2"},
		Brokers:      []string{"127.0.0.1:9094", "127.0.0.1:9093"},
		ConsumerNum:  2,
		Version:      "0.10.2.0",
	}

	conf.WithCallbackSetup(func(session sarama.ConsumerGroupSession) {
		fmt.Println("setup...")
	})
	conf.WithCallbackConsumeClaim(func(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) {
		fmt.Println("consume claim...")
	})
	conf.WithCallbackCleanup(func(session sarama.ConsumerGroupSession) {
		fmt.Println("cleanup...")
	})

	c, err := conf.Build()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = c.Watch()
	if err != nil {
		fmt.Println(err)
		return
	}
	defer c.Close()
	<-time.Tick(10 * time.Second)
}

func TestConsumer(t *testing.T) {
	conf := &ConsumerConfig{
		ClientID:  "clientid",
		Version:   "0.10.2.0",
		Topic:     "app_test",
		Partition: 0,
		Brokers:   []string{"127.0.0.1:9094", "127.0.0.1:9093"},
		SASL: ConsumerSASLConfig{
			Enable: false,
		},
		TLS: ConsumerTLSConfig{
			Enable: false,
		},
	}

	conf.WithCallbackError(func(s *sarama.ConsumerError) {
		fmt.Println("callback error.")
	})

	c, err := conf.Build()
	defer func() {
		err := c.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()
	if err != nil {
		fmt.Println(err)
		return
	}

	ch, err := c.Watch()
	if err != nil {
		fmt.Println(err)
		return
	}

	go func() {
		<-time.Tick(5 * time.Second)
		c.Close()
	}()

	for v := range ch {
		fmt.Println(v.Topic, ": ", string(v.Value))
	}
}

func TestAsyncConsumer(t *testing.T) {
	conf := &ConsumerConfig{
		ClientID:  "clientid",
		Version:   "0.10.2.0",
		Topic:     "app_test",
		Partition: 0,
		Brokers:   []string{"127.0.0.1:9094", "127.0.0.1:9093"},
		SASL: ConsumerSASLConfig{
			Enable: false,
		},
		TLS: ConsumerTLSConfig{
			Enable: false,
		},
	}

	conf.WithCallbackError(func(s *sarama.ConsumerError) {
		fmt.Println("callback error.")
	}).WithCallbackSuccess(func(s *sarama.ConsumerMessage) {
		fmt.Println(s.Topic, ": ", string(s.Value))
	})

	c, err := conf.Build()
	defer func() {
		err := c.Close()
		if err != nil {
			fmt.Println(err)
		}
	}()
	if err != nil {
		fmt.Println(err)
		return
	}

	err = c.AsyncWatch()
	if err != nil {
		fmt.Println(err)
		return
	}
	<-time.Tick(10 * time.Second)
}
