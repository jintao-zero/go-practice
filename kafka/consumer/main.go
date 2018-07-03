package main

import (
    "gopkg.in/bsm/sarama-cluster.v2"
    "github.com/Shopify/sarama"
    "log"
)

func main() {
    srcConfig := cluster.NewConfig()
    srcConfig.Consumer.Fetch.Default = 3 * 1024 * 1024
    srcConfig.Consumer.Offsets.Initial = sarama.OffsetNewest

    addrs := []string{"10.46.226.168:9092"}
    client, err := cluster.NewClient(addrs, srcConfig)
    if err != nil {
        log.Println(err)
        return
    }

    //topics := []string{"ss-liantong-" + "20171106"}
    topics := []string{"kline_stock_0"}
    consumer, err := cluster.NewConsumerFromClient(client, "kl", topics)
    if err != nil {
        log.Println("")
    }
    defer consumer.Close()

    for {
       select {
        case m, ok := <- consumer.Messages():
            if !ok {
                log.Println("message chan close")
                return
            }
            log.Println(string(m.Value))
        }
    }
}
