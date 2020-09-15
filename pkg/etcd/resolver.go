package etcd

import (
	"context"
	"fmt"
	"time"

	"github.com/coreos/etcd/clientv3"
	"github.com/coreos/etcd/mvcc/mvccpb"
	"google.golang.org/grpc/resolver"
)

type etcdResolver struct {
	conf *ResolverConfig
	cli  *clientv3.Client

	clientConn resolver.ClientConn
}

func newResolver(t *ResolverConfig) (resolver.Builder, error) {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   t.Endpoints,
		DialTimeout: time.Duration(t.DialTimeout) * time.Second,
	})
	if err != nil {
		return nil, err
	}
	return &etcdResolver{
		conf: t,
		cli:  cli,
	}, nil
}

func (t *etcdResolver) Build(target resolver.Target, cc resolver.ClientConn, opts resolver.BuildOptions) (resolver.Resolver, error) {
	t.clientConn = cc
	go t.watch(fmt.Sprintf("%s://%s/%s", target.Scheme, target.Authority, target.Endpoint))
	return t, nil
}

func (t *etcdResolver) Scheme() string {
	return t.conf.Scheme
}

func (t *etcdResolver) ResolveNow(rn resolver.ResolveNowOptions) {
}

func (t *etcdResolver) Close() {
}

func (t *etcdResolver) watch(keyPrefix string) {
	var addrList []resolver.Address

	resp, err := t.cli.Get(context.Background(), keyPrefix, clientv3.WithPrefix())
	if err != nil {
		fmt.Println(err)
	} else {
		for _, v := range resp.Kvs {
			addrList = append(addrList, resolver.Address{
				Addr: string(v.Value),
			})
		}
	}
	t.clientConn.UpdateState(resolver.State{Addresses: addrList})

	watchChan := t.cli.Watch(context.Background(), keyPrefix, clientv3.WithPrefix())
	for v := range watchChan {
		for _, event := range v.Events {
			addr := string(event.Kv.Value)
			switch event.Type {
			case mvccpb.PUT:
				if !exist(addrList, addr) {
					addrList = append(addrList, resolver.Address{
						Addr: addr,
					})
					t.clientConn.UpdateState(resolver.State{Addresses: addrList})
				}
			case mvccpb.DELETE:
				if newAddrList, ok := remove(addrList, addr); ok {
					addrList = newAddrList
					t.clientConn.UpdateState(resolver.State{Addresses: addrList})
				}
			}
		}
	}
}

func exist(addrList []resolver.Address, addr string) bool {
	for _, v := range addrList {
		if v.Addr == addr {
			return true
		}
	}
	return false
}

func remove(addrList []resolver.Address, addr string) ([]resolver.Address, bool) {
	for i := range addrList {
		if addrList[i].Addr == addr {
			addrList[i] = addrList[len(addrList)-1]
			return addrList[:len(addrList)-1], true
		}
	}
	return nil, false
}
