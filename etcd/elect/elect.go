package main

import (
    "os"
    "time"
    "fmt"
    "errors"
    "context"
    "os/signal"
    "syscall"

    "gopkg.in/coreos/etcd.v3/clientv3/concurrency"
    "github.com/coreos/etcd/clientv3"
)

func campaign(c *clientv3.Client, election string, prop string) error {
    s, err := concurrency.NewSession(c)
    if err != nil {
        return err
    }
    e := concurrency.NewElection(s, election)
    ctx, cancel := context.WithCancel(context.TODO())

    donec := make(chan struct{})
    sigc := make(chan os.Signal, 1)
    signal.Notify(sigc, syscall.SIGINT, syscall.SIGTERM)
    go func() {
        <-sigc
        cancel()
        close(donec)
    }()

    if err = e.Campaign(ctx, prop); err != nil {
        return err
    }

    // print key since elected
    resp, err := c.Get(ctx, e.Key())
    if err != nil {
        return err
    }
    fmt.Printf("%+v %s", *resp, e.Key())

    select {
    case <-donec:
    case <-s.Done():
        return errors.New("elect: session expired")
    }

    return e.Resign(context.TODO())
}

func main()  {
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
    campaign(client, "e","evalue")
}



