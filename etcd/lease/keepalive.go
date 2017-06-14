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

	leaseGrantRsp, err := client.Grant(context.TODO(), 30)
	fmt.Printf("%+v\n", leaseGrantRsp)

	putResp, err := client.Put(context.TODO(), "foo", "bar", clientv3.WithLease(leaseGrantRsp.ID))
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("1 %+v\n", *putResp)
	for {
		keepAliveRespCh, err := client.KeepAlive(context.TODO(), leaseGrantRsp.ID)
		if err != nil {
			log.Fatal("2", err)
		}
		keepAliveResp := <-keepAliveRespCh
		log.Printf("keepAlive %d", keepAliveResp.TTL)
		time.Sleep(3 * time.Second)
	}
}
