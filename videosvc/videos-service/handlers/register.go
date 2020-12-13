package handlers

import (
	"context"
	"fmt"
	"time"

	"go.etcd.io/etcd/clientv3"
)

func Register() {
	var (
		etcd1Addr = "localhost:2379"
		etcd2Addr = "localhost:22379"
		etcd3Addr = "localhost:32379"
		etcd4Addr = "localhost:42379"
		etcd5Addr = "localhost:52379"
	)
	cli, err := clientv3.New(clientv3.Config{
		Endpoints:   []string{etcd1Addr, etcd2Addr, etcd3Addr, etcd4Addr, etcd5Addr},
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		panic(err)
	}
	defer cli.Close()

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	resp, err := cli.Put(ctx, "/video", "http://localhost:5040")
	cancel()
	if err != nil {
		panic(err)
	}
	fmt.Println(resp)
}
