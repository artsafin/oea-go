package common

import (
	"context"
	"errors"
	"fmt"
	etcd "go.etcd.io/etcd/clientv3"
	"time"
)

const (
	connectTimeout = 30 * time.Second
	readTimeout    = 2 * time.Second
)

type Etcd struct {
	Client         *etcd.Client
	dialContext    context.Context
	requestContext context.Context
	endpoints      []string
}

func NewEtcdConnection(addrs []string) *Etcd {
	return &Etcd{
		endpoints: addrs,
	}
}

func (e *Etcd) Connect() error {
	if e.Client != nil {
		e.Close()
	}

	dialCtx, _ := context.WithTimeout(context.Background(), connectTimeout)
	e.dialContext = dialCtx

	requestCtx, _ := context.WithTimeout(context.Background(), readTimeout)
	e.requestContext = requestCtx

	kv, dialErr := etcd.New(etcd.Config{
		Endpoints:   e.endpoints,
		DialTimeout: connectTimeout,
		Context:     dialCtx,
	})
	if dialErr != nil {
		return dialErr
	}

	e.Client = kv
	return nil
}

func (e *Etcd) Close() {
	if e.Client == nil {
		return
	}
	_ = e.Client.Close()
	e.Client = nil
}

func (e *Etcd) MemberList() (*etcd.MemberListResponse, error) {
	return e.Client.MemberList(e.dialContext)
}

func (e *Etcd) GetNotEmpty(key string, opts ...etcd.OpOption) (*etcd.GetResponse, error) {
	get, err := e.Client.Get(e.requestContext, key, opts...)

	if err != nil {
		return nil, err
	}
	if get.Count == 0 || len(get.Kvs) == 0 {
		return nil, errors.New(fmt.Sprintf("empty key %v", key))
	}

	return get, nil
}

func (e *Etcd) MustGetBytes(key string, opts ...etcd.OpOption) []byte {
	get, err := e.GetNotEmpty(key, opts...)

	if err != nil {
		panic(err)
	}

	val := get.Kvs[0].Value

	if len(val) == 0 {
		panic("etcd: MustGetBytes: zero bytes")
	}

	return val
}
