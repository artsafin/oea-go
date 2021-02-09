package common

import (
	"context"
	"errors"
	"fmt"
	etcd "go.etcd.io/etcd/clientv3"
	"go.uber.org/zap"
	"google.golang.org/grpc/connectivity"
	"time"
)

const (
	connectTimeout = 30 * time.Second
	readTimeout    = 2 * time.Second
)

type EtcdService struct {
	dialContext context.Context
	endpoints   []string
	logger      *zap.SugaredLogger
}

func NewEtcdService(addrs []string, logger *zap.SugaredLogger) *EtcdService {
	dialContext, _ := context.WithTimeout(context.Background(), connectTimeout)

	return &EtcdService{
		endpoints:   addrs,
		dialContext: dialContext,
		logger:      logger,
	}
}

type EtcdConnection struct {
	client         *etcd.Client
	service        *EtcdService
	requestContext context.Context
}

func (e *EtcdService) ConnectAndPing() (conn *EtcdConnection, err error) {
	conn, connErr := e.connect()
	if connErr != nil {
		return nil, connErr
	}

	e.logger.Infow("Waiting for etcd...", "endpoints", conn.client.Endpoints(), "timeout", connectTimeout)

	membersList, membersErr := conn.MemberList()
	if membersErr != nil {
		conn.Close()
		return nil, membersErr
	}
	e.logger.Infow("Connected to etcd", "etcd members", membersList.Members)

	return conn, nil
}

func (e *EtcdService) connect() (conn *EtcdConnection, err error) {
	kv, dialErr := etcd.New(etcd.Config{
		Endpoints:   e.endpoints,
		DialTimeout: connectTimeout,
		Context:     e.dialContext,
	})
	if dialErr != nil {
		return nil, dialErr
	}

	requestContext, _ := context.WithTimeout(context.Background(), readTimeout)
	return &EtcdConnection{
		client:         kv,
		service:        e,
		requestContext: requestContext,
	}, nil
}

func (e *EtcdConnection) AliveOrReconnect() (*EtcdConnection, error) {
	if e.client == nil ||
		e.client.ActiveConnection() == nil ||
		e.client.ActiveConnection().GetState() != connectivity.Ready {
		e.Close()
		return e.service.connect()
	}

	return e, nil
}

func (e *EtcdConnection) Close() error {
	if e.client == nil {
		return nil
	}
	err := e.client.Close()
	e.client = nil

	return err
}

func (e *EtcdConnection) MemberList() (*etcd.MemberListResponse, error) {
	return e.client.MemberList(e.service.dialContext)
}

func (e *EtcdConnection) GetNotEmpty(key string, opts ...etcd.OpOption) (*etcd.GetResponse, error) {
	get, err := e.client.Get(e.requestContext, key, opts...)

	if err != nil {
		return nil, err
	}
	if get.Count == 0 || len(get.Kvs) == 0 {
		return nil, errors.New(fmt.Sprintf("empty key %v", key))
	}

	return get, nil
}

func (e *EtcdConnection) MustGetBytes(key string, opts ...etcd.OpOption) []byte {
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
