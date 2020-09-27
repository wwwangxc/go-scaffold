package kafka

import (
	"fmt"
	"testing"
	"time"

	"github.com/Shopify/sarama"
)

func TestConsumer(t *testing.T) {
	conf := &ConsumerGroupConfig{

		GroupID:      "groupid",
		BalancerName: "roundrobin",
		Topics:       []string{"app_test", "app_test2"},
		Brokers:      []string{"127.0.0.1:9094", "127.0.0.1:9093"},
		ConsumerNum:  2,
		Version:      "0.10.2.0",
	}

	conf.WithSetup(func(session sarama.ConsumerGroupSession) error {
		fmt.Println("setup...")
		return nil
	})
	conf.WithConsumeClaim(func(session sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
		fmt.Println("consume claim...")
		return nil
	})
	conf.WithCleanup(func(session sarama.ConsumerGroupSession) error {
		fmt.Println("cleanup...")
		return nil
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
