package main

import (
    "os"
    "log"
    "os/signal"

    "gopkg.in/Shopify/sarama.v1"
    "time"
    "flag"
    "strings"
    "fmt"
)
// 121.40.158.64:9092
func main()  {
    broker := flag.String("broker", "localhost:9092", "kafka broker")
    flag.Parse()
    brokers := strings.Split(*broker, ",")
    fmt.Println(brokers)
    config := sarama.NewConfig()
    config.Producer.Return.Successes = true
    producer, err := sarama.NewAsyncProducer(brokers, config)
    if err != nil {
        panic(err)
    }

    defer func() {
        if err := producer.Close(); err != nil {
            log.Fatalln(err)
        }
    }()

    // Trap SIGINT to trigger a shutdown.
    signals := make(chan os.Signal, 1)
    signal.Notify(signals, os.Interrupt)

    var enqueued, errors int
    ProducerLoop:
    for {
        select {
        case producer.Input() <- &sarama.ProducerMessage{Topic: "my_topic", Key: nil, Value: sarama.StringEncoder("testing 123")}:
            enqueued++
            time.Sleep(3 * time.Second)
        case ok := <-producer.Successes():
            log.Println(ok.Key, ok.Value, ok.Offset)
        case err := <-producer.Errors():
            log.Println("Failed to produce message", err)
            errors++
        case <-signals:
            break ProducerLoop
        }
    }

    log.Printf("Enqueued: %d; errors: %d\n", enqueued, errors)
}


