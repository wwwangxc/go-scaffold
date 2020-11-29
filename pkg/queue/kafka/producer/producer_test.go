package producer

import (
	"sync"
	"testing"
	"time"

	"github.com/Shopify/sarama"
)

func TestAsyncProducer(t *testing.T) {
	conf := &ProducerConfig{
		Topic:   "app_test",
		Brokers: []string{"127.0.0.1:9093"},
		Version: "0.10.2.0",
	}

	var wg sync.WaitGroup
	wg.Add(1)

	p, err := conf.WithCallbackError(func(s *sarama.ProducerError) {
		t.Log(s.Err)
		wg.Done()
	}).Build2Async()
	if err != nil {
		t.Log(err)
		return
	}
	defer p.Close()

	err = p.Send("Hello Kafka!!!")
	if err != nil {
		t.Log(err)
		return
	}

	go func() {
		<-time.Tick(time.Second)
		t.Log("Success")
		wg.Done()
	}()

	wg.Wait()
}

func TestSyncProducer(t *testing.T) {
	conf := &ProducerConfig{
		Topic:   "app_test",
		Brokers: []string{"127.0.0.1:9093"},
		Version: "0.10.2.0",
	}
	p, err := conf.Build()
	if err != nil {
		t.Log(err)
		return
	}
	defer p.Close()
	err = p.Send("Hello Kafka!!!")
	if err != nil {
		t.Log(err)
	} else {
		t.Log("Success")
	}
}
