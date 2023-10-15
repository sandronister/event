package main

import (
	"fmt"

	"github.com/rabbitmq/amqp091-go"
	"github.com/sandronister/event/pkg/rabbimq"
)

func main() {
	ch := rabbimq.OpenChannel()
	defer ch.Close()

	msgs := make(chan amqp091.Delivery)
	go rabbimq.Consume(ch, msgs, "orders")

	for msg := range msgs {
		fmt.Println(string(msg.Body))
		msg.Ack(false)
	}
}
