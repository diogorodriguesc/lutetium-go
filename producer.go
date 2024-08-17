package main

import (
    "encoding/json"
    "log"

    "github.com/streadway/amqp"
)

// Here we set the way error messages are displayed in the terminal.
func failOnError(err error, msg string) {
    if err != nil {
        log.Fatalf("%s: %s", msg, err)
    }
}

type Message struct {
    Name string
    Body string
    Time int64
    Event string
}

func main() {
    conn, err := amqp.Dial("amqp://guest:guest@localhost:5672/")
    failOnError(err, "Failed to connect to RabbitMQ")
    defer conn.Close()

    ch, err := conn.Channel()
    failOnError(err, "Failed to open a channel")
    defer ch.Close()

    // We create a Queue to send the message to.
    q, err := ch.QueueDeclare(
        "golang-queue", // name
        false,          // durable
        false,          // delete when unused
        false,          // exclusive
        false,          // no-wait
        nil,            // arguments
    )
    failOnError(err, "Failed to declare a queue")

    for i := range [10000]int{} {
        m := Message{"Alice", "Hello", 1294706395881547000, "product-created"}
        b, err := json.Marshal(m)

        err = ch.Publish(
            "",     // exchange
            q.Name, // routing key
            false,  // mandatory
            false,  // immediate
            amqp.Publishing{
                ContentType: "application/json",
                Body:        b,
            })
        // If there is an error publishing the message, a log will be displayed in the terminal.
        failOnError(err, "Failed to publish a message")
        log.Printf(" [x] Congrats, sending message %d: %s", i, b)
    }
}