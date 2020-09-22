package kafka

import (
	"sync"
	"testing"
	"time"
)

func TestAsyncProducer(t *testing.T) {
	conf := &ProducerConfig{
		Topic:   "app_test",
		Brokers: []string{"127.0.0.1:9093"},
	}
	p, err := conf.Build2Async()
	if err != nil {
		t.Log(err)
		return
	}
	defer p.Close()
	var wg sync.WaitGroup
	wg.Add(1)

	p.Send("Hello Kafka!!!", func(err error) {
		t.Log(err)
		wg.Done()
	})
	ticker := time.NewTicker(1 * time.Second)
	go func() {
		<-ticker.C
		t.Log("Success")
		wg.Done()
	}()
	wg.Wait()
}

func TestSyncProducer(t *testing.T) {
	conf := &ProducerConfig{
		Topic:   "app_test",
		Brokers: []string{"127.0.0.1:9093"},
	}
	p, err := conf.Build2Sync()
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
