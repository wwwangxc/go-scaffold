package register

import (
	"context"
	"fmt"
	"time"

	"github.com/coreos/etcd/clientv3"
)

// Register ..
type Register struct {
	cli  *clientv3.Client
	conf *Config
}

func newRegister(conf *Config) *Register {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   conf.Endpoints,
		DialTimeout: conf.DialTimeout,
	})
	if err != nil {
		panic(err.Error())
	}
	return &Register{
		conf: conf,
		cli:  cli,
	}
}

// RegistryService ..
func (t *Register) RegistryService(key, value string) {
	ticker := time.NewTicker(t.conf.HeartbeatCycle)
	go func() {
		for {
			resp, err := t.cli.Get(context.Background(), key)
			if err != nil {
				fmt.Printf("etcd get error:%s\n", err)
			}
			if resp.Count == 0 {
				err = t.keepAlive(key, value)
				if err != nil {
					fmt.Printf("etcd registe error:%s\n", err)
				}
			}
			<-ticker.C
		}
	}()
}

// UnRegistryService ..
func (t *Register) UnRegistryService(key string) error {
	_, err := t.cli.Delete(context.Background(), key)
	return err
}

func (t *Register) keepAlive(key, value string) error {
	lease, err := t.cli.Grant(context.Background(), int64(t.conf.LeaseTTL.Seconds()))
	if err != nil {
		fmt.Printf("etcd grant error:%s", err)
		return err
	}

	_, err = t.cli.Put(context.Background(), key, value, clientv3.WithLease(lease.ID))
	if err != nil {
		fmt.Printf("etcd put error:%s", err)
		return err
	}

	ch, err := t.cli.KeepAlive(context.Background(), lease.ID)
	if err != nil {
		fmt.Printf("etcd keep alive error:%s", err)
		return err
	}
	go func() {
		<-ch
	}()
	return nil
}
