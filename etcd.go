package main

import (
	"context"
	etcd "go.etcd.io/etcd/clientv3"
	"log"
	"oea-go/common"
	"reflect"
	"strconv"
)

func loadEtcdConfig(etcdAddrs []string) common.Config {
	log.Println("Reading config from", etcdAddrs)
	timeoutCtx, _ := context.WithTimeout(context.Background(), configTimeout)

	kv, err := etcd.New(etcd.Config{
		Endpoints:   etcdAddrs,
		DialTimeout: configTimeout,
	})
	if err != nil {
		panic(err)
	}
	defer kv.Close()

	membersList, membersErr := kv.MemberList(timeoutCtx)
	if membersErr != nil {
		panic(membersErr)
	}
	log.Printf("connected to etcd: %+v\n", membersList.Members)

	c := common.NewConfig()

	c.ForeachKey(func(name string, val interface{}, isMandatory bool, dstKind reflect.Kind) (interface{}, bool) {
		keyResp, keyErr := kv.Get(timeoutCtx, name)
		mandatoryStr := ""
		if isMandatory {
			mandatoryStr = "MANDATORY "
		}
		if keyErr != nil {
			log.Printf("config: error fetching %v%v: %v\n", mandatoryStr, name, keyErr)
			return nil, !isMandatory
		}
		if keyResp.Count == 0 {
			log.Printf("config: %v%v: empty value\n", mandatoryStr, name)
			return nil, !isMandatory
		}

		newVal := string(keyResp.Kvs[0].Value)

		switch dstKind {
		case reflect.Uint:
			newValUint, convErr := strconv.ParseUint(newVal, 10, 64)
			if convErr != nil {
				return nil, true
			}
			return uint(newValUint), true
		case reflect.Bool:
			boolVal := newVal != "0" && newVal != "false" && newVal != "no"
			return boolVal, true
		default:
			return newVal, true
		}
	})

	c.MustValidate()

	return c
}
