package main

import (
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"golang.org/x/net/context"
	"log"
	"time"
)

func main() {
	config := clientv3.Config{
		Endpoints:   []string{"localhost:2379"},
		DialTimeout: 5 * time.Second,
	}
	client, err := clientv3.New(config)
	if err != nil {
		panic(err)
		return
	}
	defer client.Close()

	lease := clientv3.NewLease(client)
	defer lease.Close()
	leaseGrantRsp, err := lease.Grant(context.TODO(), 30)
	fmt.Printf("%+v", leaseGrantRsp)

	putResp, err := client.Put(context.TODO(), "foo", "bar", clientv3.WithLease(leaseGrantRsp.ID))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("%+v", putResp)
}
