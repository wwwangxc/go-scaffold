package group

import (
	"fmt"
	"testing"
	"time"

	"github.com/Shopify/sarama"
)

func TestConsumerGroup(t *testing.T) {
	conf := &Config{

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
