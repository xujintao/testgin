/**
 * Copyright 2015-2017, Wothing Co., Ltd.
 * All rights reserved.
 *
 * Created by elvizlai on 2017/11/28 11:40.
 */

package etcd3

import (
	"context"
	"strings"
	"time"

	"github.com/coreos/etcd/clientv3"
)

type reg struct {
	stopSignal chan bool
	leaseId    clientv3.LeaseID
	cli        *clientv3.Client
}

var defaultReg *reg = &reg{stopSignal: make(chan bool, 1)}

// Register register service with name as prefix to etcd, multi etcd addr should use ; to split
func Register(etcdAddr, name string, addr string, ttl int64) error {
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   strings.Split(etcdAddr, ";"),
		DialTimeout: 15 * time.Second,
	})
	if err != nil {
		return err
	}

	leaseResp, err := cli.Grant(context.Background(), ttl)
	if err != nil {
		return err
	}

	_, err = cli.Put(context.Background(), "/"+schema+"/"+name+"/"+addr, addr, clientv3.WithLease(leaseResp.ID))
	if err != nil {
		return err
	}

	_, err = cli.KeepAlive(context.Background(), leaseResp.ID)
	if err != nil {
		return err
	}

	defaultReg.cli = cli
	defaultReg.leaseId = leaseResp.ID

	// go func() {
	// 	select {
	// 	case <-defaultReg.stopSignal:

	// 	}
	// }()

	return nil
}

// UnRegister remove service from etcd
func UnRegister(name string, addr string) error {
	if _, err := defaultReg.cli.Revoke(context.Background(), defaultReg.leaseId); err != nil {
		return err
	}

	if _, err := defaultReg.cli.Delete(context.Background(), "/"+schema+"/"+name+"/"+addr); err != nil {
		return err
	}
	return nil
}
