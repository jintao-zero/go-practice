package main

import (
    "log"
    "github.com/streadway/amqp"
    "go-practice/rabbitmq/common"
    "fmt"
)

func main()  {
    conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
    common.FailOnError(err, "Failed to connect to RabbitMQ")
    defer conn.Close()

    ch, err := conn.Channel()
    common.FailOnError(err, "Failed to open a channel")
    defer ch.Close()

    // declare a queue
    purged, err := ch.QueueDelete("hello",false, false, false)
    common.FailOnError(err, "Failed to delete the hello queue")
    fmt.Println("Success to delete the hello queue, purge msg ", purged)

    ch.QueueDeclare(
        "hello",  // name
        true, // durable
        false, // auto-deleted
        false,  // exclusive
        false, // no-wait
        nil, // arguments
    )
    body := "hello"
    err = ch.Publish(
        "",     // exchange
        "hello", // routing key
        false,  // mandatory
        false,  // immediate
        amqp.Publishing {
            //DeliveryMode: amqp.Persistent,
            ContentType: "text/plain",
            Body:        []byte(body),
        })
    log.Printf(" [x] Sent %s", body)
    common.FailOnError(err, "Failed to publish a message")
}



